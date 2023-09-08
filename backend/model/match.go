package model

// MatchData 查询出的JSON一场比赛细节数据
type MatchData struct {
	GameCreation          int                   `json:"gameCreation"`
	GameCreationDate      string                `json:"gameCreationDate"`
	GameDuration          int                   `json:"gameDuration"`
	GameId                int                   `json:"gameId"`
	GameMode              string                `json:"gameMode"`
	GameType              string                `json:"gameType"`
	GameVersion           string                `json:"gameVersion"`
	MapId                 int                   `json:"mapId"`
	ParticipantIdentities []ParticipantIdentity `json:"participantIdentities"`
	Participants          []Participant         `json:"participants"`
	PlatformId            string                `json:"platformId"`
	QueueId               int                   `json:"queueId"`
	SeasonId              int                   `json:"seasonId"`
	Teams                 []Team                `json:"teams"`
}

type Game struct {
	GameCreation          int                   `json:"gameCreation"`
	GameCreationDate      string                `json:"gameCreationDate"`
	GameDuration          int                   `json:"gameDuration"`
	GameId                int                   `json:"gameId"`
	GameMode              string                `json:"gameMode"`
	GameType              string                `json:"gameType"`
	GameVersion           string                `json:"gameVersion"`
	MapId                 int                   `json:"mapId"`
	ParticipantIdentities []ParticipantIdentity `json:"participantIdentities"`
	Participants          []Participant         `json:"participants"`
	PlatformId            string                `json:"platformId"`
	QueueId               int                   `json:"queueId"`
	SeasonId              int                   `json:"seasonId"`
	Teams                 []Team                `json:"teams"`
}

type Games struct {
	GameBeginDate  string `json:"gameBeginDate"`
	GameCount      int    `json:"gameCount"`
	GameEndDate    string `json:"gameEndDate"`
	GameIndexBegin int    `json:"gameIndexBegin"`
	GameIndexEnd   int    `json:"gameIndexEnd"`
	Games          []Game `json:"games"`
}

type MatchHistory struct {
	AccountId  int    `json:"accountId"`
	Games      Games  `json:"games"`
	PlatformId string `json:"platformId"`
}

type Participant struct {
	ChampionId                int      `json:"championId"`
	HighestAchievedSeasonTier string   `json:"highestAchievedSeasonTier"`
	ParticipantId             int      `json:"participantId"`
	Spell1Id                  int      `json:"spell1Id"`
	Spell2Id                  int      `json:"spell2Id"`
	Stats                     Stats    `json:"stats"`
	TeamId                    int      `json:"teamId"`
	Timeline                  Timeline `json:"timeline"`
}

type Timeline struct {
	CreepsPerMinDeltas          map[string]float64 `json:"creepsPerMinDeltas"`
	CsDiffPerMinDeltas          map[string]float64 `json:"csDiffPerMinDeltas"`
	DamageTakenDiffPerMinDeltas map[string]float64 `json:"damageTakenDiffPerMinDeltas"`
	DamageTakenPerMinDeltas     map[string]float64 `json:"damageTakenPerMinDeltas"`
	GoldPerMinDeltas            map[string]float64 `json:"goldPerMinDeltas"`
	Lane                        string             `json:"lane"`
	Role                        string             `json:"role"`
	ParticipantId               int                `json:"participantId"`
	XpDiffPerMinDeltas          map[string]float64 `json:"xpDiffPerMinDeltas"`
	XpPerMinDeltas              map[string]float64 `json:"xpPerMinDeltas"`
}

type Stats struct {
	Assists                         int  `json:"assists"`
	CausedEarlySurrender            bool `json:"causedEarlySurrender"`
	ChampLevel                      int  `json:"champLevel"`
	CombatPlayerScore               int  `json:"combatPlayerScore"`
	DamageDealtToObjectives         int  `json:"damageDealtToObjectives"`
	DamageDealtToTurrets            int  `json:"damageDealtToTurrets"`
	DamageSelfMitigated             int  `json:"damageSelfMitigated"`
	Deaths                          int  `json:"deaths"`
	DoubleKills                     int  `json:"doubleKills"`
	EarlySurrenderAccomplice        bool `json:"earlySurrenderAccomplice"`
	FirstBloodAssist                bool `json:"firstBloodAssist"`
	FirstBloodKill                  bool `json:"firstBloodKill"`
	FirstInhibitorAssist            bool `json:"firstInhibitorAssist"`
	FirstInhibitorKill              bool `json:"firstInhibitorKill"`
	FirstTowerAssist                bool `json:"firstTowerAssist"`
	FirstTowerKill                  bool `json:"firstTowerKill"`
	GameEndedInEarlySurrender       bool `json:"gameEndedInEarlySurrender"`
	GameEndedInSurrender            bool `json:"gameEndedInSurrender"`
	GoldEarned                      int  `json:"goldEarned"`
	GoldSpent                       int  `json:"goldSpent"`
	InhibitorKills                  int  `json:"inhibitorKills"`
	Item0                           int  `json:"item0"`
	Item1                           int  `json:"item1"`
	Item2                           int  `json:"item2"`
	Item3                           int  `json:"item3"`
	Item4                           int  `json:"item4"`
	Item5                           int  `json:"item5"`
	Item6                           int  `json:"item6"`
	KillingSprees                   int  `json:"killingSprees"`
	Kills                           int  `json:"kills"`
	LargestCriticalStrike           int  `json:"largestCriticalStrike"`
	LargestKillingSpree             int  `json:"largestKillingSpree"`
	LargestMultiKill                int  `json:"largestMultiKill"`
	LongestTimeSpentLiving          int  `json:"longestTimeSpentLiving"`
	MagicDamageDealt                int  `json:"magicDamageDealt"`
	MagicDamageDealtToChampions     int  `json:"magicDamageDealtToChampions"`
	MagicalDamageTaken              int  `json:"magicalDamageTaken"`
	NeutralMinionsKilled            int  `json:"neutralMinionsKilled"`
	NeutralMinionsKilledEnemyJungle int  `json:"neutralMinionsKilledEnemyJungle"`
	NeutralMinionsKilledTeamJungle  int  `json:"neutralMinionsKilledTeamJungle"`
	ObjectivePlayerScore            int  `json:"objectivePlayerScore"`
	ParticipantId                   int  `json:"participantId"`
	PentaKills                      int  `json:"pentaKills"`
	Perk0                           int  `json:"perk0"`
	Perk0Var1                       int  `json:"perk0Var1"`
	Perk0Var2                       int  `json:"perk0Var2"`
	Perk0Var3                       int  `json:"perk0Var3"`
	Perk1                           int  `json:"perk1"`
	Perk1Var1                       int  `json:"perk1Var1"`
	Perk1Var2                       int  `json:"perk1Var2"`
	Perk1Var3                       int  `json:"perk1Var3"`
	Perk2                           int  `json:"perk2"`
	Perk2Var1                       int  `json:"perk2Var1"`
	Perk2Var2                       int  `json:"perk2Var2"`
	Perk2Var3                       int  `json:"perk2Var3"`
	Perk3                           int  `json:"perk3"`
	Perk3Var1                       int  `json:"perk3Var1"`
	Perk3Var2                       int  `json:"perk3Var2"`
	Perk3Var3                       int  `json:"perk3Var3"`
	Perk4                           int  `json:"perk4"`
	Perk4Var1                       int  `json:"perk4Var1"`
	Perk4Var2                       int  `json:"perk4Var2"`
	Perk4Var3                       int  `json:"perk4Var3"`
	Perk5                           int  `json:"perk5"`
	Perk5Var1                       int  `json:"perk5Var1"`
	Perk5Var2                       int  `json:"perk5Var2"`
	Perk5Var3                       int  `json:"perk5Var3"`
	PerkPrimaryStyle                int  `json:"perkPrimaryStyle"`
	PerkSubStyle                    int  `json:"perkSubStyle"`
	PhysicalDamageDealt             int  `json:"physicalDamageDealt"`
	PhysicalDamageDealtToChampions  int  `json:"physicalDamageDealtToChampions"`
	PhysicalDamageTaken             int  `json:"physicalDamageTaken"`
	PlayerScore0                    int  `json:"playerScore0"`
	PlayerScore1                    int  `json:"playerScore1"`
	PlayerScore2                    int  `json:"playerScore2"`
	PlayerScore3                    int  `json:"playerScore3"`
	PlayerScore4                    int  `json:"playerScore4"`
	PlayerScore5                    int  `json:"playerScore5"`
	PlayerScore6                    int  `json:"playerScore6"`
	PlayerScore7                    int  `json:"playerScore7"`
	PlayerScore8                    int  `json:"playerScore8"`
	PlayerScore9                    int  `json:"playerScore9"`
	QuadraKills                     int  `json:"quadraKills"`
	SightWardsBoughtInGame          int  `json:"sightWardsBoughtInGame"`
	TeamEarlySurrendered            bool `json:"teamEarlySurrendered"`
	TimeCCingOthers                 int  `json:"timeCCingOthers"`
	TotalDamageDealt                int  `json:"totalDamageDealt"`
	TotalDamageDealtToChampions     int  `json:"totalDamageDealtToChampions"`
	TotalDamageTaken                int  `json:"totalDamageTaken"`
	TotalHeal                       int  `json:"totalHeal"`
	TotalMinionsKilled              int  `json:"totalMinionsKilled"`
	TotalPlayerScore                int  `json:"totalPlayerScore"`
	TotalScoreRank                  int  `json:"totalScoreRank"`
	TotalTimeCrowdControlDealt      int  `json:"totalTimeCrowdControlDealt"`
	TotalUnitsHealed                int  `json:"totalUnitsHealed"`
	TripleKills                     int  `json:"tripleKills"`
	TrueDamageDealt                 int  `json:"trueDamageDealt"`
	TrueDamageDealtToChampions      int  `json:"trueDamageDealtToChampions"`
	TrueDamageTaken                 int  `json:"trueDamageTaken"`
	TurretKills                     int  `json:"turretKills"`
	UnrealKills                     int  `json:"unrealKills"`
	VisionScore                     int  `json:"visionScore"`
	VisionWardsBoughtInGame         int  `json:"visionWardsBoughtInGame"`
	WardsKilled                     int  `json:"wardsKilled"`
	WardsPlaced                     int  `json:"wardsPlaced"`
	Win                             bool `json:"win"`
}

type ParticipantIdentity struct {
	ParticipantId int `json:"participantId"`
	Player        struct {
		AccountId         int    `json:"accountId"`
		CurrentAccountId  int    `json:"currentAccountId"`
		CurrentPlatformId string `json:"currentPlatformId"`
		MatchHistoryUri   string `json:"matchHistoryUri"`
		PlatformId        string `json:"platformId"`
		ProfileIcon       int    `json:"profileIcon"`
		Puuid             string `json:"puuid"`
		SummonerId        int    `json:"summonerId"`
		SummonerName      string `json:"summonerName"`
	} `json:"player"`
}

type Team struct {
	Bans                 []Ban  `json:"bans"`
	BaronKills           int    `json:"baronKills"`
	DominionVictoryScore int    `json:"dominionVictoryScore"`
	DragonKills          int    `json:"dragonKills"`
	FirstBaron           bool   `json:"firstBaron"`
	FirstBlood           bool   `json:"firstBlood"`
	FirstDragon          bool   `json:"firstDragon"`
	FirstInhibitor       bool   `json:"firstInhibitor"`
	FirstTower           bool   `json:"firstTower"`
	InhibitorKills       int    `json:"inhibitorKills"`
	RiftHeraldKills      int    `json:"riftHeraldKills"`
	TeamId               int    `json:"teamId"`
	TowerKills           int    `json:"towerKills"`
	VilemawKills         int    `json:"vilemawKills"`
	Win                  string `json:"win"`
}

type Ban struct {
	ChampionId int `json:"championId"`
	PickTurn   int `json:"pickTurn"`
}
