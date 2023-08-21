package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func loadingPage(w fyne.Window) fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabel("正在查询中..."),
	)
}
