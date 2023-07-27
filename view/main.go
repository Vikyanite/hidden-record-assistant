package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Show() {
	a := app.New()
	w := a.NewWindow("LOL隐藏战绩查询小助手")
	w.Resize(fyne.NewSize(400, 300))

	hello := container.NewBorder(nil, nil, nil,
		container.NewBorder(nil, nil,
			widget.NewLabel("大区: "), widget.NewSelect([]string{"艾欧尼亚", "德玛西亚"}, func(s string) {}),
		),
	)
	w.SetContent(container.NewBorder(
		hello,
		widget.NewButton("Submit", func() {
			w.SetContent(container.NewVBox(widget.NewLabel("TODO PAGE...")))
		}),
		nil, nil,
		widget.NewEntry(),
	))

	w.ShowAndRun()
}
