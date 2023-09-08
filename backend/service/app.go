package service

import (
	"hidden-record-assistant/backend/model"
	"hidden-record-assistant/backend/model/conv"
	"hidden-record-assistant/backend/service/support"
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
	data.AccountData, err = a.SummonerInquirer.GetCurrentSummonerBaseInfo()
	if err != nil {
		return
	}
	err = a.getSummoner(&data)
	return
}

func (a *App) GetSummonerByName(name string) (data model.Summoner, err error) {
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
			var rank model.Rank
			rank, err = a.SummonerInquirer.GetRank(data.AccountData.Puuid)
			if err != nil {
				return
			}
			data.RankSolo, data.RankFlex = rank.QueueMap.RANKEDSOLO5X5, rank.QueueMap.RANKEDFLEXSR
			return
		},
		func() (err error) {
			matchHistory, err := a.GetMatchHistory(data.AccountData.Puuid, 0, 10)
			if err != nil {
				return
			}
			data.MatchRecords = make([]model.MatchData, matchHistory.Games.GameCount)
			data.DisplayMatchRecords = make([]model.DisplayMatch, matchHistory.Games.GameCount)
			var tasks []func() error
			for i := range matchHistory.Games.Games {
				index := i
				tasks = append(tasks, func() (err error) {
					data.MatchRecords[index], err = a.GetMatchDetails(matchHistory.Games.Games[index].GameId)
					if err != nil {
						return
					}
					data.DisplayMatchRecords[index] = conv.GetDisplayMatchFromMatchData(data.AccountData.Puuid, data.MatchRecords[index], a.AssetsManager)
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
			data.RecentMatchStatistic = conv.GetStatisticFromMatches(data.AccountData.Puuid, data.MatchRecords)
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
