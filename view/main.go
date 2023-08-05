package view

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	theme2 "fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"hidden-record-assistant/model"
	"hidden-record-assistant/service/query"
	"hidden-record-assistant/view/theme"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type DqSlice []string

func (m DqSlice) Len() int {
	return len(m)
}

func (m DqSlice) Less(i, j int) bool {
	return DqMap[m[i]] < DqMap[m[j]]
}

func (m DqSlice) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

var (
	DqMap = map[string]int{
		"艾欧尼亚": 1, "比尔吉沃特": 2, "祖安": 3, "诺克萨斯": 4, "班德尔城": 5, "德玛西亚": 6, "皮尔特沃夫": 7,
		"战争学院": 8, "弗雷尔卓德": 9, "巨神峰": 10, "雷瑟守备": 11, "裁决之地": 13, "黑色玫瑰": 14, "暗影岛": 15,
		"恕瑞玛": 16, "钢铁烈阳": 17, "水晶之痕": 18, "均衡教派": 19, "扭曲丛林": 20, "影流": 22, "守望之海": 23,
		"征服之海": 24, "卡拉曼达": 25, "巨龙之巢": 26, "皮城警备": 27, "男爵领域": 30, "峡谷之巅": 31,
	}
	DqNames DqSlice
	DqNum   int
)

func init() {
	for k, _ := range DqMap {
		DqNames = append(DqNames, k)
	}
	sort.Sort(DqNames)
}

func Show() {
	a := app.New()
	a.Settings().SetTheme(&theme.MyTheme{})
	w := a.NewWindow("LOL隐藏战绩查询小助手")
	w.Resize(fyne.NewSize(400, 300))

	w.SetContent(mainPage(w))

	w.ShowAndRun()
}

func mainPage(w fyne.Window) fyne.CanvasObject {
	input := widget.NewMultiLineEntry()
	examplePlayerNames := "玩家1加入了队伍聊天\n" +
		"玩家2加入了队伍聊天\n" +
		"..."
	input.SetPlaceHolder(examplePlayerNames)
	input.Validator = validation.NewRegexp(`^(?:.*加入了队伍聊天\n)*.*加入了队伍聊天[\n]*$`, "请按照'xxx加入了队伍聊天'的格式输入")

	var form *widget.Form
	selector := widget.NewSelect(
		DqNames,
		func(s string) {
			DqNum = DqMap[s]
		},
	)
	selector.SetSelectedIndex(0)

	form = &widget.Form{
		Items: []*widget.FormItem{
			{Text: "输入房间文本", Widget: input},
			{
				Text:   "大区",
				Widget: selector,
			},
		},
		OnSubmit: func() {
			// 先进入到查询过度界面
			w.SetContent(loadingPage(w))

			// 开始查询
			str := strings.Split(input.Text, "\n")
			bound := len(str)
			// 取消所有的最后空行
			for {
				if bound != 0 && str[bound-1] == "" {
					bound--
				} else {
					break
				}
			}
			completeChan := make(chan struct{}, bound)
			queryResults := make([]model.Result, bound)
			for i := 0; i < bound; i++ {
				go func(i int) {
					defer func() {
						// 出错也算完成
						completeChan <- struct{}{}
						if r := recover(); r != nil {

						}
					}()
					//采用正则表达式匹配
					re := regexp.MustCompile(`^(.*?)加入了队伍聊天$`)
					match := re.FindStringSubmatch(str[i])
					if len(match) <= 1 {
						// TODO 提示错误
						return
					}
					playerName := match[1]

					res, err := query.QueryN(strconv.Itoa(DqNum), playerName, 2)
					if err != nil {
						// TODO 提示错误
						return
					}
					queryResults[i] = res
				}(i)
			}

			// 等待所有查询完成, 然后返回处理好的查询界面
			for i := 0; i < bound; i++ {
				<-completeChan
			}
			w.SetContent(queryPage(w, queryResults))
		},
		OnCancel: func() {

		},
	}
	return form
}

func loadingPage(w fyne.Window) fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabel("正在查询中..."),
	)
}

func queryPage(w fyne.Window, queryResults []model.Result) fyne.CanvasObject {
	d := displayPage(w, queryResults)
	return container.NewBorder(nil,
		widget.NewButtonWithIcon("返回", theme2.NavigateBackIcon(), func() {
			DqNum = 1 // 因为默认是艾欧尼亚
			w.SetContent(mainPage(w))
		}),
		nil, nil,
		d,
	)
}

func displayPage(w fyne.Window, queryResults []model.Result) fyne.CanvasObject {
	tabs := make([]*container.TabItem, len(queryResults))
	for i := 0; i < len(queryResults); i++ {
		tabs[i] = container.NewTabItem(queryResults[i].Name, displayItemPage(queryResults[i]))
	}
	return container.NewAppTabs(tabs...)
}

func displayItemPage(data model.Result) fyne.CanvasObject {
	obj := container.NewHSplit(
		container.NewGridWithRows(6,
			container.NewHBox(
				widget.NewLabel("段位："+data.Division),
				widget.NewLabel("胜点："+strconv.Itoa(data.WinPoint)),
			),
			container.NewBorder(nil, nil, nil, widget.NewLabel("（最近20场游戏中排位表现）")),
			widget.NewLabel("胜率："+fmt.Sprintf("%.2f%%", 100.0*data.WinRate)),
			widget.NewLabel("胜场数："+strconv.Itoa(data.WinCount)),
			widget.NewLabel("败场数："+strconv.Itoa(data.FailCount)),
		),
		// TODO 最近战绩展示界面
		widget.NewLabel("最近战绩展示界面"),
	)
	return obj
}
