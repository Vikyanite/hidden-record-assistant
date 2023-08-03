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

func (r *Result) CalWinDefeat() {
	for i := range r.FightRecords {
		if r.FightRecords[i].IsWin {
			r.WinCount++
		} else {
			r.FailCount++
		}
	}

	// 防止除0
	if r.WinCount+r.FailCount != 0 {
		r.WinRate = float64(r.WinCount) / float64(r.WinCount+r.FailCount)
	}
}

type FightRecord struct {
	KDA   string
	IsMVP bool
	IsWin bool
	// 比赛时间
	GameTime string
}
