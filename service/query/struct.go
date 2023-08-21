package query

import (
	"fmt"
	"golang.org/x/net/html"
	"hidden-record-assistant/lib/htmlx"
	"hidden-record-assistant/model"
	"hidden-record-assistant/util"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type QueryForm url.Values

func NewQueryForm(dq, name string) QueryForm {
	return QueryForm{
		"name":  {name},
		"dq":    {dq},
		"start": {"0"},
		"end":   {"10"},
		"type":  {"lol"},
		"sc":    {util.MD5(name, dq)},
	}
}

func (q QueryForm) Form() url.Values {
	return url.Values(q)
}

func (q QueryForm) NextNPage(plusNum int) QueryForm {
	start, _ := strconv.Atoi(q["start"][0])
	end, _ := strconv.Atoi(q["end"][0])

	return QueryForm{
		"name":  {q["name"][0]},
		"dq":    {q["dq"][0]},
		"start": {strconv.Itoa(start + plusNum)},
		"end":   {strconv.Itoa(end + plusNum)},
		"type":  {q["type"][0]},
		"sc":    {q["sc"][0]},
	}
}

type RawResult struct {
	// 战绩拼接html
	Zhanji string `json:"zhanji"`
	// 单双等级（段位）
	Dsdj string `json:"dsdj"`
	// 单双胜点
	Dssd string `json:"dssd"`
}

func (r *RawResult) Cook() (ret model.Result) {
	ret = model.Result{
		Division: r.Dsdj,
	}
	fmt.Sscanf(r.Dssd, "%d胜点", &ret.WinPoint)
	ret.FightRecords = append(ret.FightRecords, GetFightRecords(r.Zhanji)...)
	return
}

func GetFightRecords(str string) (ret []model.FightRecord) {
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

	// 找到 "mb-2"类下的"type"类的Context为"单排/双排"的节点
	nodes := htmlx.FindElements(root, func(n *html.Node) bool {
		for _, attr := range n.Attr {
			if attr.Key == "class" && attr.Val == "mb-2" {
				found := htmlx.FindElement(n, htmlx.HasClass("type"))
				if found == nil {
					return false
				}
				return htmlx.GetTextContent(found) == "单排/双排"
			}
		}
		return false
	})
	for i := range nodes {
		temp := model.FightRecord{
			KDA:      GetKDA(nodes[i]),
			GameTime: GetGameTime(nodes[i]),
			Role:     GetRole(nodes[i]),
		}
		temp.IsWin, temp.GameLen = GetGameLenAndIsWin(nodes[i])
		ret = append(ret, temp)
	}
	return
}

func GetGameTime(node *html.Node) (gtime string) {
	found := htmlx.FindElement(node, htmlx.HasClass("relative time-stamp"))
	if found == nil {
		return
	}
	gtime = htmlx.GetTextContent(found.FirstChild)
	fmt.Printf(gtime)
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

// GetGameLenAndIsWin 提取游戏时间和是否胜利信息
// 按规范不应该两个信息放在一个函数里提取，但是方便捏~
func GetGameLenAndIsWin(node *html.Node) (IsWin bool, gameTime string) {
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

func GetRole(node *html.Node) (role string) {
	found := htmlx.FindElement(node, htmlx.HasClass("role"))
	if found == nil {
		return
	}
	role = htmlx.GetTextContent(found)
	fmt.Printf(role)
	return
}
