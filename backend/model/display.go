package model

type DisplayMatch struct {
	Poz              int    `json:"poz"`
	QueueDescription string `json:"queueDescription"`
	GameDuration     string `json:"gameDuration"`
	GameDurationInt  int    `json:"gameDurationInt"`
	GameTimeAgo      string `json:"gameTimeAgo"`
	// kp是团战参与率
	Kp string `json:"kp"`
	// NMValue 牛马值
	NMValue               int                   `json:"nmValue"`
	SpellD                Spell                 `json:"spellD"`
	SpellF                Spell                 `json:"spellF"`
	RunePrimary           Perk                  `json:"runePrimary"`
	RuneSecondary         PerkStyle             `json:"runeSecondaryStyle"`
	ToggleAdvancedDetails bool                  `json:"toggleAdvancedDetails"`
	RealChampsNames       []string              `json:"realChampsNames"`
	Participants          []Participant         `json:"participants"`
	ParticipantIdentities []ParticipantIdentity `json:"participantIdentities"`
	Overview              Overview              `json:"overview"`
	Breakdown             Breakdown             `json:"breakdown"`
	Teams                 []Team                `json:"teams"`
}

type Overview struct {
	TeamRed  DisplayTeam `json:"teamRed"`
	TeamBlue DisplayTeam `json:"teamBlue"`
	SpellDs  []Spell     `json:"spellDs"`
	SpellFs  []Spell     `json:"spellFs"`
}

type DisplayTeam struct {
	Kills   int `json:"kills"`
	Deaths  int `json:"deaths"`
	Assists int `json:"assists"`
	Gold    int `json:"gold"`
}

// RecentMatchStatistic 根据召唤师的比赛数据计算出的统计数据
type RecentMatchStatistic struct {
	MainPlayerPoz []int `json:"mainPlayerPoz"`
	Kills         int   `json:"kills"`
	Assists       int   `json:"assists"`
	Deaths        int   `json:"deaths"`
	Wins          int   `json:"wins"`
	Defeats       int   `json:"defeats"`

	// 牛马值
	NMValue int `json:"nmValue"`

	Gold            int `json:"gold"`
	VisionScore     int `json:"vision_score"`
	FirstBloodTimes int `json:"firstBloodTimes"`
	CS              int `json:"cs"`
	ControlWards    int `json:"control_wards"`

	Lane                []string `json:"lane"`
	PreferablyLane      string   `json:"preferablyLane"`
	PreferablyLaneGames int      `json:"preferablyLaneGames"`
	Champs              []int    `json:"champs"`
	PreferablyChamp1    int      `json:"preferablyChamp1"`
	PreferablyChamp2    int      `json:"preferablyChamp2"`

	PreferablyChamp1Games int `json:"preferablyChamp1_games"`
	PreferablyChamp2Games int `json:"preferablyChamp2_games"`

	KillsPreferablyLane   int `json:"kills_preferablyLane"`
	DeathsPreferablyLane  int `json:"deaths_preferablyLane"`
	AssistsPreferablyLane int `json:"assists_preferablyLane"`
	WinsPreferablyLane    int `json:"wins_preferablyLane"`
	DefeatsPreferablyLane int `json:"defeats_preferablyLane"`

	KillsPreferablyChamps   [2]int `json:"kills_preferablyChamps"`
	DeathsPreferablyChamps  [2]int `json:"deaths_preferablyChamps"`
	AssistsPreferablyChamps [2]int `json:"assists_preferablyChamps"`
	WinsPreferablyChamps    [2]int `json:"wins_preferablyChamps"`
	DefeatsPreferablyChamps [2]int `json:"defeats_preferablyChamps"`
}

type Summoner struct {
	AccountData          SummonerBaseInfo     `json:"dataAccount"`
	RankSolo             RankDetails          `json:"dataRank_solo"`
	RankFlex             RankDetails          `json:"dataRank_flex"`
	RecentMatchStatistic RecentMatchStatistic `json:"matchData"`
	DisplayMatchRecords  []DisplayMatch       `json:"displayMatchHistory"`
}
