package service

import (
	"hidden-record-assistant/backend/model"
	"hidden-record-assistant/backend/model/conv"
	"hidden-record-assistant/backend/module/errs"
	"hidden-record-assistant/backend/module/task"
	"hidden-record-assistant/backend/service/support"
	"time"
)

type App struct {
	*support.AssetsManager
	*support.SummonerInquirer
}

func NewApp(conn *support.Connector) *App {
	return &App{
		AssetsManager:    support.NewAssetsManager(conn),
		SummonerInquirer: support.NewSummonerInquirer(conn),
	}
}

func (a *App) Init() (err error) {
	err = a.AssetsManager.Init()
	if err != nil {
		return
	}
	return
}

func (a *App) GetCurrentSummoner() (data model.Summoner, err error) {
	if !a.IsRunning() {
		err = errs.ErrNotRunning
		return
	}
	data.AccountData, err = a.SummonerInquirer.GetCurrentSummonerBaseInfo()
	if err != nil {
		return
	}
	err = a.getSummoner(&data)
	return
}

func (a *App) GetSummonerByName(name string) (data model.Summoner, err error) {
	if !a.IsRunning() {
		err = errs.ErrNotRunning
		return
	}
	data.AccountData, err = a.SummonerInquirer.GetSummonerBaseInfoByName(name)
	if err != nil {
		return
	}
	err = a.getSummoner(&data)
	return
}

func (a *App) getSummoner(data *model.Summoner) (err error) {
	// 这些任务都是需要等待AccountData的，并且是互相独立的，所以可以并发执行
	errChan, num := support.ExecuteTaskConcurrently(
		func() (err error) {
			defer task.TimeCost("GetRank")()
			var rank model.Rank
			rank, err = a.SummonerInquirer.GetRank(data.AccountData.Puuid)
			if err != nil {
				return
			}
			data.RankSolo, data.RankFlex = rank.QueueMap.RANKEDSOLO5X5, rank.QueueMap.RANKEDFLEXSR
			data.RankSolo.TierOb = a.AssetsManager.Tier[data.RankSolo.Tier]
			data.RankFlex.TierOb = a.AssetsManager.Tier[data.RankFlex.Tier]
			return
		},
		func() (err error) {
			defer task.TimeCost("GetMatchHistory")()
			matchHistory, err := a.GetMatchHistory(data.AccountData.Puuid, 0, 19)
			if err != nil {
				return
			}
			MatchRecords := make([]model.MatchData, matchHistory.Games.GameCount)
			data.DisplayMatchRecords = make([]model.DisplayMatch, matchHistory.Games.GameCount)
			var (
				NMValueInFiveHours  = 0
				CntInFiveHours      = 0
				NMValueOutFiveHours = 0
				CntOutFiveHours     = 0
			)

			var tasks []func() error
			for i := range matchHistory.Games.Games {
				index := i
				tasks = append(tasks, func() (err error) {
					MatchRecords[index], err = a.GetMatchDetails(matchHistory.Games.Games[index].GameId)
					if err != nil {
						return
					}
					data.DisplayMatchRecords[index] = conv.GetDisplayMatchFromMatchData(data.AccountData.Puuid, MatchRecords[index], a.AssetsManager)
					if time.Now().UnixMilli()-int64(MatchRecords[index].GameCreation) <= 5*60*60*1000 {
						NMValueInFiveHours += data.DisplayMatchRecords[index].NMValue
						CntInFiveHours++
					} else {
						NMValueOutFiveHours += data.DisplayMatchRecords[index].NMValue
						CntOutFiveHours++
					}
					return
				})
			}
			errChan, num := support.ExecuteTaskConcurrently(tasks...)
			for i := 0; i < num; i++ {
				err = <-errChan
				if err != nil {
					return
				}
			}
			data.RecentMatchStatistic = conv.GetStatisticFromMatches(data.AccountData.Puuid, MatchRecords)
			if CntInFiveHours != 0 {
				data.RecentMatchStatistic.NMValue = (NMValueInFiveHours*8/CntInFiveHours + NMValueOutFiveHours*2/CntOutFiveHours) / 10
			} else {
				data.RecentMatchStatistic.NMValue = NMValueOutFiveHours / CntOutFiveHours
			}
			return
		},
	)
	for i := 0; i < num; i++ {
		err = <-errChan
		if err != nil {
			return
		}
	}
	return
}
