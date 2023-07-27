package model

type Result struct {
	// 胜负场
	WinCount  int
	FailCount int
	// 胜率
	WinRate float64
	// 胜点
	WinPoint string
	// 段位
	Division string
	// 战绩
	FightRecords []FightRecord
}

type FightRecord struct {
	KDA   string
	IsMVP bool
	IsWin bool
	// 比赛时间
	GameTime string
}
