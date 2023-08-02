package model

type Result struct {
	// 召唤师名字
	Name string
	// 胜负场
	WinCount  int
	FailCount int
	// 胜率
	WinRate float64
	// 胜点
	WinPoint int
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
