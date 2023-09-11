package conv

import (
	"fmt"
	"hidden-record-assistant/backend/model"
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
	for i := range allMatches {
		for j := range allMatches[i].ParticipantIdentities {
			if puuid == allMatches[i].ParticipantIdentities[j].Player.Puuid {
				ret.MainPlayerPoz = append(ret.MainPlayerPoz, j)
				break
			}
		}
	}
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
		// 有点奇怪的，LOL中role才是lane
		ret.Lane = append(ret.Lane, allMatches[i].Participants[ret.MainPlayerPoz[i]].Timeline.Role)

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

	//get most 2 played champs
	ret.PreferablyChamp1, ret.PreferablyChamp1Games, ret.PreferablyChamp2, ret.PreferablyChamp2Games = findMostFrequentElements(ret.Champs)

	//get stats for most played lane and most played champs games
	for i := range allMatches {
		if allMatches[i].Participants[ret.MainPlayerPoz[i]].Timeline.Role == ret.PreferablyLane {
			ret.KillsPreferablyLane += allMatches[i].Participants[ret.MainPlayerPoz[i]].Stats.Kills
			ret.DeathsPreferablyLane += allMatches[i].Participants[ret.MainPlayerPoz[i]].Stats.Deaths
			ret.AssistsPreferablyLane += allMatches[i].Participants[ret.MainPlayerPoz[i]].Stats.Assists
			if allMatches[i].Participants[ret.MainPlayerPoz[i]].Stats.Win {
				ret.WinsPreferablyLane++
			} else {
				ret.DefeatsPreferablyLane++
			}
		}
		if allMatches[i].Participants[ret.MainPlayerPoz[i]].ChampionId == ret.PreferablyChamp1 {
			ret.KillsPreferablyChamps[0] += allMatches[i].Participants[ret.MainPlayerPoz[i]].Stats.Kills
			ret.DeathsPreferablyChamps[0] += allMatches[i].Participants[ret.MainPlayerPoz[i]].Stats.Deaths
			ret.AssistsPreferablyChamps[0] += allMatches[i].Participants[ret.MainPlayerPoz[i]].Stats.Assists

			if allMatches[i].Participants[ret.MainPlayerPoz[i]].Stats.Win {
				ret.WinsPreferablyChamps[0]++
			} else {
				ret.DefeatsPreferablyChamps[0]++
			}
		}

		if allMatches[i].Participants[ret.MainPlayerPoz[i]].ChampionId == ret.PreferablyChamp2 {
			ret.KillsPreferablyChamps[1] += allMatches[i].Participants[ret.MainPlayerPoz[i]].Stats.Kills
			ret.DeathsPreferablyChamps[1] += allMatches[i].Participants[ret.MainPlayerPoz[i]].Stats.Deaths
			ret.AssistsPreferablyChamps[1] += allMatches[i].Participants[ret.MainPlayerPoz[i]].Stats.Assists

			if allMatches[i].Participants[ret.MainPlayerPoz[i]].Stats.Win {
				ret.WinsPreferablyChamps[1]++
			} else {
				ret.DefeatsPreferablyChamps[1]++
			}
		}
	}
	return
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
		GameDuration:          fmt.Sprintf("%02dmin %02ds", match.GameDuration/60, match.GameDuration-match.GameDuration/60*60),
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

	return
}

func msToTime(duration int) (gameTimeAgo string) {
	// ms to ANYTIME
	minutes := duration / (1000 * 60) % 60
	hours := (duration / (1000 * 60 * 60)) % 24
	days := duration / (24 * 60 * 60 * 1000)

	if duration < 3600000 {
		gameTimeAgo = fmt.Sprintf("%02d minutes", minutes)
	} else if duration <= 86399999 {
		gameTimeAgo = fmt.Sprintf("%d hours", hours)
	} else {
		if duration > 86399999 && duration <= 172799999 {
			gameTimeAgo = fmt.Sprintf("%d day", days)
		} else {
			gameTimeAgo = fmt.Sprintf("%d days", days)
		}
	}
	return
}

func FixMatchData(match *model.MatchData, manager IAssets) {
	match.QueueObject = manager.GetQueueById(match.QueueId)
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
