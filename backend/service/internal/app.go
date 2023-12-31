package internal

import (
	lcuapi "github.com/Vikyanite/lcu-driver"
	lcumodel "github.com/Vikyanite/lcu-driver/model"
	"hidden-record-assistant/backend/model"
	"hidden-record-assistant/backend/module/task"
	"time"
)

type App struct {
	State int32
}

const (
	InitQueryMatchBeg = 0
	InitQueryMatchEnd = 19
)

func NewApp() *App {
	return &App{}
}

func (a *App) Start(startCbs ...func() error) (chan error, error) {
	return lcuapi.Start(startCbs...)
}

func (a *App) GetCurrentSummoner() (data model.DisplaySummoner, err error) {
	data.AccountData, err = lcuapi.GetCurrentSummoner()
	if err != nil {
		return
	}
	err = a.getSummoner(&data)
	return
}

func (a *App) GetSummonerByName(name string) (data model.DisplaySummoner, err error) {
	data.AccountData, err = lcuapi.GetSummonerByName(name)
	if err != nil {
		return
	}
	err = a.getSummoner(&data)
	return
}

func (a *App) getSummoner(data *model.DisplaySummoner) (err error) {
	defer task.TimeCost("getSummoner")()
	// 这些任务都是需要等待AccountData的，并且是互相独立的，所以可以并发执行
	errChan, num := task.ExecuteTaskConcurrently(
		func() (err error) {
			defer task.TimeCost("GetRank")()
			var rank lcumodel.RankedStats
			rank, err = lcuapi.GetRankedStatsByPuuid(data.AccountData.Puuid)
			if err != nil {
				return
			}
			data.RankSolo, data.RankFlex = rank.QueueMap.RANKEDSOLO5X5, rank.QueueMap.RANKEDFLEXSR
			return
		},
		func() (err error) {
			MatchRecords, err := a.getMatchRecordsByPuuid(data.AccountData.Puuid, 0, 19)
			if err != nil {
				return
			}
			var (
				NMValueInFiveHours  = 0
				CntInFiveHours      = 0
				NMValueOutFiveHours = 0
				CntOutFiveHours     = 0
			)

			data.DisplayMatchRecords, err = a.matchRecordsToDisplayMatchRecords(data.AccountData.Puuid, MatchRecords)
			if err != nil {
				return
			}

			for i := 0; i < len(MatchRecords); i++ {
				if time.Now().UnixMilli()-int64(MatchRecords[i].GameCreation) <= 5*60*60*1000 {
					NMValueInFiveHours += data.DisplayMatchRecords[i].NMValue
					CntInFiveHours++
				} else {
					NMValueOutFiveHours += data.DisplayMatchRecords[i].NMValue
					CntOutFiveHours++
				}
			}

			data.RecentMatchStatistic = GetStatisticFromMatches(data.AccountData.Puuid, MatchRecords)
			if CntInFiveHours != 0 {
				data.RecentMatchStatistic.NMValue = (NMValueInFiveHours*8/CntInFiveHours + NMValueOutFiveHours*2/CntOutFiveHours) / 10
			} else if CntOutFiveHours != 0 {
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

func (a *App) GetDisplayMatchRecordsByPuuid(puuid string, beg, end int) (matchRecords []model.DisplayMatch, err error) {
	MatchRecords, err := a.getMatchRecordsByPuuid(puuid, beg, end)
	if err != nil {
		return
	}
	matchRecords, err = a.matchRecordsToDisplayMatchRecords(puuid, MatchRecords)

	return
}

func (a *App) matchRecordsToDisplayMatchRecords(puuid string, MatchRecords []lcumodel.MatchData) (matchRecords []model.DisplayMatch, err error) {
	defer task.TimeCost("matchRecordsToDisplayMatchRecords")()
	matchRecords = make([]model.DisplayMatch, len(MatchRecords))
	//var tasks []func() error
	for i := 0; i < len(MatchRecords); i++ {
		index := i
		//tasks = append(tasks, func() (err error) {
		matchRecords[index] = GetDisplayMatchFromMatchData(puuid, MatchRecords[index])
		//	return
		//})
	}
	//errChan, num := task.ExecuteTaskConcurrently(tasks...)
	//for i := 0; i < num; i++ {
	//	err = <-errChan
	//	if err != nil {
	//		return
	//	}
	//}
	return
}

func (a *App) getMatchRecordsByPuuid(puuid string, beg, end int) (MatchRecords []lcumodel.MatchData, err error) {
	defer task.TimeCost("getMatchRecordsByPuuid")()
	matchHistory, err := lcuapi.GetMatchHistoryByPuuid(puuid, beg, end)
	if err != nil {
		return
	}
	MatchRecords = make([]lcumodel.MatchData, len(matchHistory.Games.Games))
	var tasks []func() error
	for i := range matchHistory.Games.Games {
		index := i
		tasks = append(tasks, func() (err error) {
			MatchRecords[index], err = lcuapi.GetMatchDetailsByGameId(matchHistory.Games.Games[index].GameId)
			if err != nil {
				return
			}
			return
		})
	}
	errChan, num := task.ExecuteTaskConcurrently(tasks...)
	for i := 0; i < num; i++ {
		err = <-errChan
		if err != nil {
			return
		}
	}
	return
}
