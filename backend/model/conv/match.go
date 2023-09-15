package conv

import (
	"fmt"
	"hidden-record-assistant/backend/model"
	"strings"
	"time"
)

type IAssets interface {
	GetChampionById(id int) model.Champion
	GetQueueById(id int) model.Queue
	GetPerkById(id int) model.Perk
	GetPerkStyleById(id int) model.PerkStyle
	GetItemById(id int) model.Item
	GetSpellById(id int) model.Spell
}

func GetStatisticFromMatches(puuid string, allMatches []model.MatchData) (ret model.RecentMatchStatistic) {

	// 寻找主玩家在每个比赛记录中的的index
	for i := range allMatches {
		for j := range allMatches[i].ParticipantIdentities {
			if puuid == allMatches[i].ParticipantIdentities[j].Player.Puuid {
				ret.MainPlayerPoz = append(ret.MainPlayerPoz, j)
				break
			}
		}
	}

	//get most 2 played champs
	ret.PreferablyChamp1, ret.PreferablyChamp1Games, ret.PreferablyChamp2, ret.PreferablyChamp2Games = findMostFrequentElements(ret.Champs)

	for i := range allMatches {
		ret.Kills += allMatches[i].Participants[ret.MainPlayerPoz[i]].Stats.Kills
		ret.Deaths += allMatches[i].Participants[ret.MainPlayerPoz[i]].Stats.Deaths
		ret.Assists += allMatches[i].Participants[ret.MainPlayerPoz[i]].Stats.Assists
		//get wins, defeats
		if allMatches[i].Participants[ret.MainPlayerPoz[i]].Stats.Win {
			ret.Wins++
		} else {
			ret.Defeats++
		}

		//get lanes
		ret.Lane = append(ret.Lane, allMatches[i].Participants[ret.MainPlayerPoz[i]].Timeline.Lane)

		//get champs
		ret.Champs = append(ret.Champs, allMatches[i].Participants[ret.MainPlayerPoz[i]].ChampionId)

		//get gold
		ret.Gold += allMatches[i].Participants[ret.MainPlayerPoz[i]].Stats.GoldEarned / (allMatches[i].GameDuration / 60)

		//get vision score
		ret.VisionScore += allMatches[i].Participants[ret.MainPlayerPoz[i]].Stats.VisionScore

		//first blood kill or assists
		if allMatches[i].Participants[ret.MainPlayerPoz[i]].Stats.FirstBloodAssist || allMatches[i].Participants[ret.MainPlayerPoz[i]].Stats.FirstBloodKill {
			ret.FirstBloodTimes++
		}

		//get cs
		ret.CS += (allMatches[i].Participants[ret.MainPlayerPoz[i]].Stats.NeutralMinionsKilled + allMatches[i].Participants[ret.MainPlayerPoz[i]].Stats.TotalMinionsKilled) / (allMatches[i].GameDuration / 60)

		//get control wards
		ret.ControlWards += allMatches[i].Participants[ret.MainPlayerPoz[i]].Stats.VisionWardsBoughtInGame

	}

	nrmax := -100
	for i := range ret.Lane {
		nr := 0
		for j := range ret.Lane {
			if ret.Lane[j] == ret.Lane[i] {
				nr++
			}
		}
		if nr > nrmax {
			nrmax = nr
			ret.PreferablyLane = ret.Lane[i]
		}
	}
	ret.PreferablyLaneGames = nrmax

	//get stats for most played lane and most played champs games
	for i, poz := range ret.MainPlayerPoz {
		if allMatches[i].Participants[poz].Timeline.Lane == ret.PreferablyLane {
			ret.KillsPreferablyLane += allMatches[i].Participants[poz].Stats.Kills
			ret.DeathsPreferablyLane += allMatches[i].Participants[poz].Stats.Deaths
			ret.AssistsPreferablyLane += allMatches[i].Participants[poz].Stats.Assists
			if allMatches[i].Participants[poz].Stats.Win {
				ret.WinsPreferablyLane++
			} else {
				ret.DefeatsPreferablyLane++
			}
		}
		if allMatches[i].Participants[poz].ChampionId == ret.PreferablyChamp1 {
			ret.KillsPreferablyChamps[0] += allMatches[i].Participants[poz].Stats.Kills
			ret.DeathsPreferablyChamps[0] += allMatches[i].Participants[poz].Stats.Deaths
			ret.AssistsPreferablyChamps[0] += allMatches[i].Participants[poz].Stats.Assists

			if allMatches[i].Participants[poz].Stats.Win {
				ret.WinsPreferablyChamps[0]++
			} else {
				ret.DefeatsPreferablyChamps[0]++
			}
		}

		if allMatches[i].Participants[poz].ChampionId == ret.PreferablyChamp2 {
			ret.KillsPreferablyChamps[1] += allMatches[i].Participants[poz].Stats.Kills
			ret.DeathsPreferablyChamps[1] += allMatches[i].Participants[poz].Stats.Deaths
			ret.AssistsPreferablyChamps[1] += allMatches[i].Participants[poz].Stats.Assists

			if allMatches[i].Participants[poz].Stats.Win {
				ret.WinsPreferablyChamps[1]++
			} else {
				ret.DefeatsPreferablyChamps[1]++
			}
		}
	}

	return
}

func GetNMValueFromMatches(match model.MatchData, MainPlayerPoz int) (NMValue int) {

	// 牛马值计算
	NMValue = 100 // 牛马值默认初始值100
	// 计算参团率 = (个人K+A / 队伍K)
	// 计算KDA
	{
		//  参团率:第一 +10 第二 +5 第四 -5 第五 -10
		teamKAs := []int{}
		teamKills := 0
		PozInTeam := 0
		for j := range match.Participants {
			if match.Participants[j].TeamId == match.Participants[MainPlayerPoz].TeamId {
				teamKills += match.Participants[j].Stats.Kills
				if j == MainPlayerPoz {
					PozInTeam = len(teamKAs)
				}
				teamKAs = append(teamKAs, match.Participants[j].Stats.Kills+match.Participants[j].Stats.Assists)
			}
		}
		max1, max2, min1, min2 := findMaxMinTwoValues(teamKAs)
		if teamKAs[PozInTeam] == max1 {
			NMValue += 10
		} else if teamKAs[PozInTeam] == max2 {
			NMValue += 5
		} else if teamKAs[PozInTeam] == min1 {
			NMValue -= 10
		} else if teamKAs[PozInTeam] == min2 {
			NMValue -= 5
		}

		//人头占比>50%并且人头>5  + 10
		//人头占比>50%并且人头>10  + 20
		//人头占比>50%并且人头>15  + 40
		//
		//人头占比>35%并且人头>5  + 5
		//人头占比>35%并且人头>10  + 10
		//人头占比>35%并且人头>15  + 20

		if match.Participants[MainPlayerPoz].Stats.Kills*2 > teamKills {
			if match.Participants[MainPlayerPoz].Stats.Kills > 5 {
				NMValue += 10
			} else if match.Participants[MainPlayerPoz].Stats.Kills > 10 {
				NMValue += 20
			} else if match.Participants[MainPlayerPoz].Stats.Kills > 15 {
				NMValue += 40
			}
		} else if match.Participants[MainPlayerPoz].Stats.Kills*100 > teamKills*35 {
			if match.Participants[MainPlayerPoz].Stats.Kills > 5 {
				NMValue += 5
			} else if match.Participants[MainPlayerPoz].Stats.Kills > 10 {
				NMValue += 10
			} else if match.Participants[MainPlayerPoz].Stats.Kills > 15 {
				NMValue += 20
			}
		}

		//助攻占比>50%并且助攻>5  + 10
		//助攻占比>50%并且助攻>10  + 20
		//助攻占比>50%并且助攻>15  + 40
		//
		//助攻占比>35%并且助攻>5  + 5
		//助攻占比>35%并且助攻>10  + 10
		//助攻占比>35%并且助攻>15  + 20
		if match.Participants[MainPlayerPoz].Stats.Assists*2 > teamKills {
			if match.Participants[MainPlayerPoz].Stats.Assists > 5 {
				NMValue += 10
			} else if match.Participants[MainPlayerPoz].Stats.Assists > 10 {
				NMValue += 20
			} else if match.Participants[MainPlayerPoz].Stats.Assists > 15 {
				NMValue += 40
			}
		} else if match.Participants[MainPlayerPoz].Stats.Assists*100 > teamKills*35 {
			if match.Participants[MainPlayerPoz].Stats.Assists > 5 {
				NMValue += 5
			} else if match.Participants[MainPlayerPoz].Stats.Kills > 10 {
				NMValue += 10
			} else if match.Participants[MainPlayerPoz].Stats.Kills > 15 {
				NMValue += 20
			}
		}

		//再加上 (k+a)/d + (k-d)/5*参团率
		{
			k := match.Participants[MainPlayerPoz].Stats.Kills
			d := match.Participants[MainPlayerPoz].Stats.Deaths
			if d == 0 {
				d = 1
			}
			a := match.Participants[MainPlayerPoz].Stats.Assists
			if teamKills == 0 {
				teamKills = 1
			}
			NMValue += (k+a)/d + (k-d)/5*(k+a)/teamKills
		}
	}
	//zlog.Debugf("NMValue after KDA: %d", NMValue)
	// 计算金钱比
	{
		//  第一 +10 第二 +5 第四（非辅助） -5 第五（非辅助） -10
		teamGolds := []int{}
		PozInTeam := 0
		for j := range match.Participants {
			if match.Participants[j].TeamId == match.Participants[MainPlayerPoz].TeamId {
				if j == MainPlayerPoz {
					PozInTeam = len(teamGolds)
				}
				teamGolds = append(teamGolds, match.Participants[j].Stats.GoldEarned)
			}
		}
		max1, max2, min1, min2 := findMaxMinTwoValues(teamGolds)
		if teamGolds[PozInTeam] == max1 {
			NMValue += 10
		} else if teamGolds[PozInTeam] == max2 {
			NMValue += 5
		} else if match.Participants[PozInTeam].Timeline.Lane != "SUPPORT" && teamGolds[PozInTeam] == min1 {
			NMValue -= 10
		} else if match.Participants[PozInTeam].Timeline.Lane != "SUPPORT" && teamGolds[PozInTeam] == min2 {
			NMValue -= 5
		}
	}
	//zlog.Debugf("NMValue after Gold: %d", NMValue)
	// 计算伤害比
	{
		// 第一 +10 第二 +5
		SumTeamDMG := 0
		teamDMGs := []int{}
		PozInTeam := 0
		for j := range match.Participants {
			if match.Participants[j].TeamId == match.Participants[MainPlayerPoz].TeamId {
				SumTeamDMG += match.Participants[j].Stats.TotalDamageDealtToChampions
				if j == MainPlayerPoz {
					PozInTeam = len(teamDMGs)
				}
				teamDMGs = append(teamDMGs, match.Participants[j].Stats.TotalDamageDealtToChampions)
			}
		}
		max1, max2, _, _ := findMaxMinTwoValues(teamDMGs)
		if teamDMGs[PozInTeam] == max1 {
			NMValue += 10
		} else if teamDMGs[PozInTeam] == max2 {
			NMValue += 5
		}

		//伤害占比>50%并且人头>5  + 10
		//伤害占比>50%并且人头>10  + 20
		//伤害占比>50%并且人头>15  + 40
		//
		//伤害占比>30%并且人头>5  + 5
		//伤害占比>30%并且人头>10  + 10
		//伤害占比>30%并且人头>15  + 20
		if match.Participants[MainPlayerPoz].Stats.TotalDamageDealtToChampions*2 > SumTeamDMG {
			if match.Participants[MainPlayerPoz].Stats.TotalDamageDealtToChampions > 5 {
				NMValue += 10
			} else if match.Participants[MainPlayerPoz].Stats.TotalDamageDealtToChampions > 10 {
				NMValue += 20
			} else if match.Participants[MainPlayerPoz].Stats.TotalDamageDealtToChampions > 15 {
				NMValue += 40
			}
		} else if match.Participants[MainPlayerPoz].Stats.TotalDamageDealtToChampions*100 > SumTeamDMG*30 {
			if match.Participants[MainPlayerPoz].Stats.TotalDamageDealtToChampions > 5 {
				NMValue += 5
			} else if match.Participants[MainPlayerPoz].Stats.TotalDamageDealtToChampions > 10 {
				NMValue += 10
			} else if match.Participants[MainPlayerPoz].Stats.TotalDamageDealtToChampions > 15 {
				NMValue += 20
			}
		}

	}
	//zlog.Debugf("NMValue after DMG: %d", NMValue)
	// 计算视野比
	{
		// 第一 +10 第二 +5
		teamVSs := []int{}
		PozInTeam := 0
		for j := range match.Participants {
			if match.Participants[j].TeamId == match.Participants[MainPlayerPoz].TeamId {
				if j == MainPlayerPoz {
					PozInTeam = len(teamVSs)
				}
				teamVSs = append(teamVSs, match.Participants[j].Stats.VisionScore)
			}
		}
		max1, max2, _, _ := findMaxMinTwoValues(teamVSs)
		if teamVSs[PozInTeam] == max1 {
			NMValue += 10
		} else if teamVSs[PozInTeam] == max2 {
			NMValue += 5
		}
	}
	//zlog.Debugf("NMValue after VS: %d", NMValue)
	// 计算金钱转化比
	// 百度说是团队伤害占比/团队金钱占比
	{
		// 第一 +10 第二 +5
		SumTeamDMG := 0
		SumTeamGold := 0
		TeamTransRates := []int{}
		PozInTeam := 0
		for j := range match.Participants {
			if match.Participants[j].TeamId == match.Participants[MainPlayerPoz].TeamId {
				SumTeamDMG += match.Participants[j].Stats.TotalDamageDealtToChampions
				SumTeamGold += match.Participants[j].Stats.GoldEarned
			}
		}
		for j := range match.Participants {
			if match.Participants[j].TeamId == match.Participants[MainPlayerPoz].TeamId {
				if j == MainPlayerPoz {
					PozInTeam = len(TeamTransRates)
				}
				rate := match.Participants[j].Stats.TotalDamageDealtToChampions * 100 * SumTeamGold / match.Participants[j].Stats.GoldEarned
				if SumTeamDMG != 0 {
					rate /= SumTeamDMG
				}
				TeamTransRates = append(TeamTransRates, rate)
			}
		}
		max1, max2, _, _ := findMaxMinTwoValues(TeamTransRates)
		if TeamTransRates[PozInTeam] == max1 {
			NMValue += 10
		} else if TeamTransRates[PozInTeam] == max2 {
			NMValue += 5
		}

	}
	//zlog.Debugf("NMValue after TransRate: %d", NMValue)
	// 计算补刀
	{
		//每分钟 8个刀以上加5分
		//每分钟 9个刀以上加10
		//每分钟 10个刀以上加20
		for j := range match.Participants {
			if match.Participants[j].TeamId == match.Participants[MainPlayerPoz].TeamId {
				cs := match.Participants[j].Stats.TotalMinionsKilled + match.Participants[j].Stats.NeutralMinionsKilled
				if cs*60/match.GameDuration > 10 {
					NMValue += 20
				} else if cs*60/match.GameDuration > 9 {
					NMValue += 10
				} else if cs*60/match.GameDuration > 8 {
					NMValue += 5
				}
			}
		}
	}
	//zlog.Debugf("NMValue after CS: %d", NMValue)
	// 三杀四杀五杀
	{
		//三杀+5
		//四杀+10
		//五杀+20
		NMValue += 5 * match.Participants[MainPlayerPoz].Stats.TripleKills
		NMValue += 10 * match.Participants[MainPlayerPoz].Stats.QuadraKills
		NMValue += 20 * match.Participants[MainPlayerPoz].Stats.PentaKills
	}
	//zlog.Debugf("NMValue after MultiKills: %d", NMValue)

	return
}

func findMaxMinTwoValues(arr []int) (max1, max2, min1, min2 int) {
	var maxValue = []int{-1, -1}
	var minValue = []int{999, 999}
	for i := range arr {
		if maxValue[0] == -1 || arr[i] > maxValue[0] {
			maxValue[1] = maxValue[0]
			maxValue[0] = arr[i]
		} else if arr[i] != maxValue[0] && (maxValue[1] == -1 || arr[i] > maxValue[1]) {
			maxValue[1] = arr[i]
		}

		if minValue[0] == 999 || arr[i] < minValue[0] {
			minValue[1] = minValue[0]
			minValue[0] = arr[i]
		} else if arr[i] != minValue[0] && (minValue[1] == 999 || arr[i] < minValue[1]) {
			minValue[1] = arr[i]
		}
	}
	return maxValue[0], maxValue[1], minValue[0], minValue[1]
}

func findMostFrequentElements(arr []int) (max1, cnt1, max2, cnt2 int) {
	counts := make(map[int]int)
	var mostFrequent = []int{-1, -1}

	for i := range arr {
		counts[arr[i]]++
		if mostFrequent[0] == -1 || counts[arr[i]] > counts[mostFrequent[0]] {
			mostFrequent[1] = mostFrequent[0]
			mostFrequent[0] = arr[i]
		} else if arr[i] != mostFrequent[0] && (mostFrequent[1] == -1 || counts[arr[i]] > counts[mostFrequent[1]]) {
			mostFrequent[1] = arr[i]
		}
	}

	return mostFrequent[0], counts[mostFrequent[0]], mostFrequent[1], counts[mostFrequent[1]]
}

func GetDisplayMatchFromMatchData(puuid string, match model.MatchData, manager IAssets) (data model.DisplayMatch) {
	FixMatchData(&match, manager)
	// Poz
	var Poz int
	for i := range match.ParticipantIdentities {
		//pozitia in array ul cu meci al player ului cautat
		if puuid == match.ParticipantIdentities[i].Player.Puuid {
			Poz = i
			break
		}
	}

	// Other
	data = model.DisplayMatch{
		Poz:                   Poz,
		ToggleAdvancedDetails: false,
		GameDuration:          fmt.Sprintf("%02d分 %02d秒", match.GameDuration/60, match.GameDuration-match.GameDuration/60*60),
		GameTimeAgo:           msToTime(int(time.Now().UnixMilli()) - match.GameCreation),
		QueueDescription:      match.QueueObject.Description,
		SpellD:                match.Participants[Poz].Spell1Object,
		SpellF:                match.Participants[Poz].Spell2Object,
		RunePrimary:           match.Participants[Poz].Stats.Perk0Object,
		RuneSecondary:         match.Participants[Poz].Stats.PerkSubStyleObject,
		ParticipantIdentities: match.ParticipantIdentities,
		GameDurationInt:       match.GameDuration,
		Teams:                 match.Teams,
	}

	// Kp & Participants
	teamkills := 0
	for i, participant := range match.Participants {
		if match.Participants[i].TeamId == match.Participants[Poz].TeamId {
			teamkills += match.Participants[i].Stats.Kills
		}
		data.Participants = append(data.Participants, participant)

		// Overview.TeamBlue/TeamRed
		if i < 5 {
			data.Overview.TeamBlue.Kills += participant.Stats.Kills
			data.Overview.TeamBlue.Assists += participant.Stats.Assists
			data.Overview.TeamBlue.Deaths += participant.Stats.Deaths
			data.Overview.TeamBlue.Gold += participant.Stats.GoldEarned
		} else {
			data.Overview.TeamRed.Kills += participant.Stats.Kills
			data.Overview.TeamRed.Assists += participant.Stats.Assists
			data.Overview.TeamRed.Deaths += participant.Stats.Deaths
			data.Overview.TeamRed.Gold += participant.Stats.GoldEarned
		}

		// Overview.SpellDs/SpellFs
		data.Overview.SpellDs = append(data.Overview.SpellDs, match.Participants[i].Spell1Object)
		data.Overview.SpellFs = append(data.Overview.SpellFs, match.Participants[i].Spell2Object)

		// Breakdown
		dmg := float64(data.Participants[i].Stats.TotalMinionsKilled + data.Participants[i].Stats.NeutralMinionsKilled)
		if dmg > data.Breakdown.HighestDMG.Value {
			data.Breakdown.HighestDMG = model.BreakdownItem{
				Value:          dmg,
				ChampionObject: data.Participants[i].ChampionObject,
				SummonerName:   data.ParticipantIdentities[i].Player.SummonerName,
			}
		}

		cs := float64(data.Participants[i].Stats.TotalMinionsKilled + data.Participants[i].Stats.NeutralMinionsKilled)
		if cs > data.Breakdown.HighestCS.Value {
			data.Breakdown.HighestCS = model.BreakdownItem{
				Value:          cs,
				ChampionObject: data.Participants[i].ChampionObject,
				SummonerName:   data.ParticipantIdentities[i].Player.SummonerName,
			}
		}

		gold := float64(data.Participants[i].Stats.GoldEarned)
		if gold > data.Breakdown.BestGold.Value {
			data.Breakdown.BestGold = model.BreakdownItem{
				Value:          gold,
				ChampionObject: data.Participants[i].ChampionObject,
				SummonerName:   data.ParticipantIdentities[i].Player.SummonerName,
			}
		}

		vs := float64(data.Participants[i].Stats.VisionScore)
		if vs > data.Breakdown.BestVS.Value {
			data.Breakdown.BestVS = model.BreakdownItem{
				Value:          vs,
				ChampionObject: data.Participants[i].ChampionObject,
				SummonerName:   data.ParticipantIdentities[i].Player.SummonerName,
			}
		}

		kda := float64(data.Participants[i].Stats.Kills + data.Participants[i].Stats.Assists)
		if data.Participants[i].Stats.Deaths != 0 {
			kda /= float64(data.Participants[i].Stats.Deaths)
		}
		if kda > data.Breakdown.HighestKDA.Value {
			data.Breakdown.HighestKDA = model.BreakdownItem{
				Value:          kda,
				ChampionObject: data.Participants[i].ChampionObject,
				SummonerName:   data.ParticipantIdentities[i].Player.SummonerName,
			}
		}

		mvp := float64(data.Participants[i].Stats.Kills) + 0.75*float64(data.Participants[i].Stats.Assists) - float64(data.Participants[i].Stats.Deaths)*0.5
		if mvp > data.Breakdown.Mvp.Value {
			data.Breakdown.Mvp = model.BreakdownItem{
				Value:          mvp,
				ChampionObject: data.Participants[i].ChampionObject,
				SummonerName:   data.ParticipantIdentities[i].Player.SummonerName,
			}
		}

		// RealChampsName
		data.RealChampsNames = append(data.RealChampsNames, match.Participants[i].ChampionObject.Name)
	}

	if teamkills != 0 {
		data.Kp = fmt.Sprintf("%d", (match.Participants[Poz].Stats.Kills+match.Participants[Poz].Stats.Assists)*100/teamkills)
	}

	data.NMValue = GetNMValueFromMatches(match, Poz)

	return
}

func msToTime(duration int) (gameTimeAgo string) {
	// ms to ANYTIME
	minutes := duration / (1000 * 60) % 60
	hours := (duration / (1000 * 60 * 60)) % 24
	days := duration / (24 * 60 * 60 * 1000)

	if duration < 3600000 {
		gameTimeAgo = fmt.Sprintf("%02d 分", minutes)
	} else if duration <= 86399999 {
		gameTimeAgo = fmt.Sprintf("%d 小时", hours)
	} else {
		gameTimeAgo = fmt.Sprintf("%d 天", days)
	}
	return
}

func FixMatchData(match *model.MatchData, manager IAssets) {
	match.QueueObject = manager.GetQueueById(match.QueueId)
	//因为 "排位赛 单双排/灵活组排"这个描述太长了，所以要处理一下,去掉"排位赛"前缀
	match.QueueObject.Description = strings.Replace(match.QueueObject.Description, "排位赛", "", 1)

	for i := range match.Participants {
		match.Participants[i].Spell1Object = manager.GetSpellById(match.Participants[i].Spell1Id)
		match.Participants[i].Spell2Object = manager.GetSpellById(match.Participants[i].Spell2Id)

		match.Participants[i].Stats.Perk0Object = manager.GetPerkById(match.Participants[i].Stats.Perk0)

		match.Participants[i].Stats.PerkSubStyleObject = manager.GetPerkStyleById(match.Participants[i].Stats.PerkSubStyle)

		match.Participants[i].ChampionObject = manager.GetChampionById(match.Participants[i].ChampionId)

		match.Participants[i].Stats.Item0Object = manager.GetItemById(match.Participants[i].Stats.Item0)
		match.Participants[i].Stats.Item1Object = manager.GetItemById(match.Participants[i].Stats.Item1)
		match.Participants[i].Stats.Item2Object = manager.GetItemById(match.Participants[i].Stats.Item2)
		match.Participants[i].Stats.Item3Object = manager.GetItemById(match.Participants[i].Stats.Item3)
		match.Participants[i].Stats.Item4Object = manager.GetItemById(match.Participants[i].Stats.Item4)
		match.Participants[i].Stats.Item5Object = manager.GetItemById(match.Participants[i].Stats.Item5)
		match.Participants[i].Stats.Item6Object = manager.GetItemById(match.Participants[i].Stats.Item6)
	}
	for i := range match.Teams {
		for j := range match.Teams[i].Bans {
			match.Teams[i].Bans[j].ChampionObject = manager.GetChampionById(match.Teams[i].Bans[j].ChampionId)
		}
	}

}
