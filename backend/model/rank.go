package model

type RankDetails struct {
	Division                       string      `json:"division"`
	HighestDivision                string      `json:"highestDivision"`
	HighestTier                    string      `json:"highestTier"`
	IsProvisional                  bool        `json:"isProvisional"`
	LeaguePoints                   int         `json:"leaguePoints"`
	Losses                         int         `json:"losses"`
	MiniSeriesProgress             string      `json:"miniSeriesProgress"`
	PreviousSeasonAchievedDivision string      `json:"previousSeasonAchievedDivision"`
	PreviousSeasonAchievedTier     string      `json:"previousSeasonAchievedTier"`
	PreviousSeasonEndDivision      string      `json:"previousSeasonEndDivision"`
	PreviousSeasonEndTier          string      `json:"previousSeasonEndTier"`
	ProvisionalGameThreshold       int         `json:"provisionalGameThreshold"`
	ProvisionalGamesRemaining      int         `json:"provisionalGamesRemaining"`
	QueueType                      string      `json:"queueType"`
	RatedRating                    int         `json:"ratedRating"`
	RatedTier                      string      `json:"ratedTier"`
	Tier                           string      `json:"tier"`
	Warnings                       interface{} `json:"warnings"` // 使用 interface{} 表示可能是任何类型
	Wins                           int         `json:"wins"`
}

type Rank struct {
	QueueMap QueueMap `json:"queueMap"`
}

type QueueMap struct {
	CHERRY            RankDetails `json:"CHERRY"`
	RANKEDFLEXSR      RankDetails `json:"RANKED_FLEX_SR"`
	RANKEDSOLO5X5     RankDetails `json:"RANKED_SOLO_5x5"`
	RANKEDTFT         RankDetails `json:"RANKED_TFT"`
	RANKEDTFTDOUBLEUP RankDetails `json:"RANKED_TFT_DOUBLE_UP"`
	RANKEDTFTTURBO    RankDetails `json:"RANKED_TFT_TURBO"`
}
