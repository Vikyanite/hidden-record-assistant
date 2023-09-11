package model

type Breakdown struct {
	Mvp        BreakdownItem `json:"mvp"`
	HighestCS  BreakdownItem `json:"highestCS"`
	HighestDMG BreakdownItem `json:"highestDMG"`
	BestGold   BreakdownItem `json:"bestGold"`
	BestVS     BreakdownItem `json:"bestVS"`
	HighestKDA BreakdownItem `json:"highestKDA"`
}

type BreakdownItem struct {
	Value          float64  `json:"value"`
	ChampionObject Champion `json:"championOb"`
	SummonerName   string   `json:"summonerName"`
}
