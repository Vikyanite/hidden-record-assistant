package view

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	theme2 "fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"hidden-record-assistant/model"
	"strconv"
)

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
	return container.NewHSplit(displayItemLeftPage(data), displayItemRightPage(data.FightRecords))
}

func displayItemLeftPage(data model.Result) fyne.CanvasObject {
	return container.NewGridWithRows(6,
		container.NewHBox(
			widget.NewLabel("段位："+data.Division),
			widget.NewLabel("胜点："+strconv.Itoa(data.WinPoint)),
		),
		container.NewBorder(nil, nil, nil, widget.NewLabel("（最近20场游戏中排位表现）")),
		widget.NewLabel("胜率："+fmt.Sprintf("%.2f%%", 100.0*data.WinRate)),
		widget.NewLabel("胜场数："+strconv.Itoa(data.WinCount)),
		widget.NewLabel("败场数："+strconv.Itoa(data.FailCount)),
	)
}

type recordItem struct {
	fyne.CanvasObject
	// 比赛时间
	GameLen  *widget.Label
	GameTime *widget.Label
	// 选取英雄
	Hero *widget.Label
	// 角色
	Role *widget.Label
	// KDA
	KDA *widget.Label
}

func (r *recordItem) SetData(record model.FightRecord) {
	r.GameLen.SetText(record.GameLen)
	r.GameTime.SetText(record.GameTime)
	r.Hero.SetText(record.Hero)
	r.Role.SetText(record.Role)
	r.KDA.SetText(record.KDA)
}

func displayItemRightPage(records []model.FightRecord) fyne.CanvasObject {
	obj := widget.NewList(
		func() int {
			return len(records)
		},
		// TODO details page
		func() fyne.CanvasObject {
			return container.NewGridWithRows(2,
				container.NewGridWithColumns(2,
					container.NewBorder(nil, nil,
						widget.NewLabel("英雄"),
						widget.NewLabel("角色"),
					),
					container.NewBorder(nil, nil, nil, widget.NewLabel("时长")),
				),
				container.NewGridWithColumns(2,
					widget.NewLabel("kda"),
					container.NewBorder(nil, nil, nil, widget.NewLabel("时间")),
				),
			)
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			grid := o.(*fyne.Container)
			row0 := grid.Objects[0].(*fyne.Container)
			row0.Objects[0].(*fyne.Container).Objects[0].(*widget.Label).SetText("英雄") //records[i].Hero)
			row0.Objects[0].(*fyne.Container).Objects[1].(*widget.Label).SetText("角色") //records[i].Role)
			row0.Objects[1].(*fyne.Container).Objects[0].(*widget.Label).SetText(records[i].GameLen)

			row1 := grid.Objects[1].(*fyne.Container)
			row1.Objects[0].(*widget.Label).SetText(records[i].KDA[:len(records[i].KDA)-1])
			row1.Objects[1].(*fyne.Container).Objects[0].(*widget.Label).SetText(records[i].GameTime)
		},
	)
	obj.Refresh()
	return obj
}
