package query

import (
	"fmt"
	"golang.org/x/net/html"
	"hidden-record-assistant/lib/htmlx"
	"hidden-record-assistant/util"
	"net/url"
	"strings"
	"time"
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
	defer func() func() {
		pre := time.Now()
		return func() {
			fmt.Printf("time cost: %s\n", time.Since(pre))
		}
	}()()

	// 解析html
	root, err := html.Parse(strings.NewReader(str))
	if err != nil {
		return
	}

	// 找到 "mobile-game-item mobile-game-item"类下的"type"类的Context为"单排/双排"的节点
	DSRecords := htmlx.FindElements(root, func(n *html.Node) bool {
		for _, attr := range n.Attr {
			if attr.Key == "class" && strings.HasPrefix(attr.Val, "mobile-game-item mobile-game-item-") {
				found := htmlx.FindElement(n, htmlx.HasClass("type"))
				if found == nil {
					return false
				}
				return htmlx.GetTextContent(found) == "单排/双排"
			}
		}
		return false
	})
	for i := range DSRecords {
		temp := FightRecord{
			KDA: GetKDA(DSRecords[i]),
		}
		temp.IsWin, temp.GameTime = GetGameTimeAndIsWin(DSRecords[i])
		ret = append(ret, temp)
	}
	return
}

func GetKDA(node *html.Node) (KDA string) {
	found := htmlx.FindElement(node, htmlx.HasClass("k-d-a"))
	if found == nil {
		return
	}
	// 非常离谱，这个网站的html的KDA数据是这么返回的
	// 所以我们只能遍历下面三个为htmlnode的child然后获取内容拼起来
	/*
		<div class="k-d-a">
			<span>2</span>
			/
			<span class="d">3</span>
			/
			<span>8</span>
		</div>
	*/
	for c := found.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode {
			KDA += htmlx.GetTextContent(c) + "/"
		}
	}
	return
}

// GetGameTimeAndIsWin 提取游戏时间和是否胜利信息
// 按规范不应该两个信息放在一个函数里提取，但是方便捏~
func GetGameTimeAndIsWin(node *html.Node) (IsWin bool, gameTime string) {
	found := htmlx.FindElement(node, htmlx.HasClass("game-result"))
	if found == nil {
		return
	}
	var winStr string
	_, err := fmt.Sscanf(htmlx.GetTextContent(found), "%s | 用时:%s", &winStr, &gameTime)
	if err != nil {
		return
	}
	//fmt.Printf("winStr: %s gameTime: %s\n", winStr, gameTime)
	IsWin = winStr == "胜利"
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
