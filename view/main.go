package view

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	theme2 "fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"hidden-record-assistant/view/theme"
	"sort"
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
	examplePlayerNames := "玩家1加入了房间\n" +
		"玩家2加入了房间\n" +
		"..."
	input.SetPlaceHolder(examplePlayerNames)
	input.Validator = validation.NewRegexp(`^(?:.*加入了房间\n)*.*加入了房间[\n]*$`, "请按照'xxx加入了房间'的格式输入")

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
			if bound != 0 && str[bound-1] == "" {
				bound--
			}

			completeChan := make(chan struct{}, bound)
			for i := 0; i < bound; i++ {
				go func(i int) {
					defer func() {
						// 出错也算完成
						completeChan <- struct{}{}
						if r := recover(); r != nil {

						}
					}()
					playerName := ""
					_, err := fmt.Sscanf(str[i], "%s加入了房间", &playerName)
					if err != nil {
						// TODO 提示错误
						return
					}

					//TODO 接下来是需要把res插入到查询界面中用于显示
					//res, err := query.SendQuery(playerName, strconv.Itoa(DqNum))
					//if err != nil {
					//	// TODO 提示错误
					//	return
					//}

				}(i)
			}

			// 等待所有查询完成, 然后返回处理好的查询界面
			for i := 0; i < bound; i++ {
				<-completeChan
			}
			w.SetContent(queryPage(w))
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

func queryPage(w fyne.Window) fyne.CanvasObject {
	// TODO 设计查询结果界面
	return container.NewVBox(
		widget.NewLabel(fmt.Sprintf("大区: %d", DqNum)),
		widget.NewButtonWithIcon("返回", theme2.NavigateBackIcon(), func() {
			DqNum = 1 // 因为默认是艾欧尼亚
			w.SetContent(mainPage(w))
		}),
	)
}
