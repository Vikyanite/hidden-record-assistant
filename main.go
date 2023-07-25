package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"hidden-record-assistant/service/query"
)

func ShowUI() {
	a := app.New()
	w := a.NewWindow("LOL隐藏战绩查询小助手")

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
	))

	w.ShowAndRun()
}

func TestHTTP(name, dq string) (err error) {
	res, _ := query.SendQuery(dq, name)

	fmt.Printf("%+v\n", res.Division)
	fmt.Printf("%+v\n", res.WinCount)

	return
}

func main() {
	_ = TestHTTP("乞力马扎罗的雪丶", "1")
}
