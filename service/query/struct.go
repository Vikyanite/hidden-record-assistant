package query

import (
	"fmt"
	"golang.org/x/net/html"
	"hidden-record-assistant/util"
	"net/url"
	"strings"
)

func NewQueryForm(dq, name string) url.Values {
	return url.Values{
		"name":  {name},
		"dq":    {dq},
		"start": {"0"},
		"end":   {"10"},
		"type":  {"lol"},
		"sc":    {util.MD5(name, dq)},
	}
}

type RawResult struct {
	// 战绩拼接html
	Zhanji string `json:"zhanji"`
	// 单双等级（段位）
	Dsdj string `json:"dsdj"`
	// 单双胜点
	Dssd string `json:"dssd"`
	// 单双胜负（x胜x负）
	Dssf string `json:"dssf"`
}

func (r *RawResult) Cook() (ret Result) {
	ret = Result{
		Division: r.Dsdj,
		WinPoint: r.Dssd,
	}

	fmt.Sscanf(r.Dssf, "%d胜%d负", &ret.WinCount, &ret.FailCount)

	// 防止除0
	if ret.WinCount+ret.FailCount != 0 {
		ret.WinRate = float64(ret.WinCount) / float64(ret.WinCount+ret.FailCount)
	}
	ret.FightRecords = append(ret.FightRecords, GetFightRecords(r.Zhanji)...)
	return
}

func GetFightRecords(str string) (ret []FightRecord) {
	fmt.Printf("%s\n", str)
	root, err := html.Parse(strings.NewReader(str))
	if err != nil {
		return
	}

	return
}

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
