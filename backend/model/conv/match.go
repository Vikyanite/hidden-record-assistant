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
		QueueDescription:      manager.GetQueueById(match.QueueId).Description,
		SpellD:                manager.GetSpellById(match.Participants[Poz].Spell1Id),
		SpellF:                manager.GetSpellById(match.Participants[Poz].Spell2Id),
		RunePrimary:           manager.GetPerkById(match.Participants[Poz].Stats.Perk0),
		RuneSecondary:         manager.GetPerkStyleById(match.Participants[Poz].Stats.PerkSubStyle),
		ParticipantIdentities: match.ParticipantIdentities,
		GameDurationInt:       match.GameDuration,
	}

	// Kp & Participants
	teamkills := 0
	for i := range match.Participants {
		if match.Participants[i].TeamId == match.Participants[Poz].TeamId {
			teamkills += match.Participants[i].Stats.Kills
		}
		data.Participants = append(data.Participants, GetDisplayParticipantFromParticipant(match.Participants[i], manager))
	}
	if teamkills != 0 {
		data.Kp = fmt.Sprintf("%d", (match.Participants[Poz].Stats.Kills+match.Participants[Poz].Stats.Assists)*100/teamkills)
	}

	// RealChampsName
	for i := range match.Participants {
		data.RealChampsNames = append(data.RealChampsNames, manager.GetChampionById(match.Participants[i].ChampionId).Name)
	}

	return
}

func GetDisplayParticipantFromParticipant(p model.Participant, manager IAssets) (data model.DisplayParticipant) {
	data.Stats = GetDisplayStatsFromStats(p.Stats, manager)
	data.ChampionId = p.ChampionId
	return
}

func GetDisplayStatsFromStats(s model.Stats, manager IAssets) (data model.DisplayStats) {
	data = model.DisplayStats{
		Kills:                s.Kills,
		Deaths:               s.Deaths,
		Assists:              s.Assists,
		Win:                  s.Win,
		TotalMinionsKilled:   s.TotalMinionsKilled,
		NeutralMinionsKilled: s.NeutralMinionsKilled,
		ChampLevel:           s.ChampLevel,
		VisionScore:          s.VisionScore,
		Item0:                manager.GetItemById(s.Item0),
		Item1:                manager.GetItemById(s.Item1),
		Item2:                manager.GetItemById(s.Item2),
		Item3:                manager.GetItemById(s.Item3),
		Item4:                manager.GetItemById(s.Item4),
		Item5:                manager.GetItemById(s.Item5),
		Item6:                manager.GetItemById(s.Item6),
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
