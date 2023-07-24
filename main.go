package main

import (
	md52 "crypto/md5"
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"io"
	"net/http"
	"net/url"
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

func MD5(name, dq string) string {
	return fmt.Sprintf("%x", md52.Sum([]byte("name="+name+"dq-"+dq)))
}

func TestHTTP(name, dq string) {
	data := url.Values{
		"name":  {name},
		"dq":    {dq},
		"start": {"0"},
		"end":   {"10"},
		"type":  {"lol"},
		"sc":    {MD5(name, dq)},
	}
	resp, err := http.PostForm("http://www.sslol.top/api/lol.php?act=cx", data)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// handle error
		panic(err)
	}
	fmt.Println(string(body))
}

func main() {
	TestHTTP("乞力马扎罗的雪丶", "1")
}
