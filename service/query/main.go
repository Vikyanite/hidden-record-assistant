package query

import (
	"encoding/json"
	"hidden-record-assistant/model"
	"io"
	"net/http"
	"net/url"
)

func query(formData QueryForm) (res model.Result, err error) {
	resp, err := http.PostForm("http://www.sslol.top/api/lol.php?act=cx", url.Values(formData))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var rres RawResult
	if err = json.Unmarshal(body, &rres); err != nil {
		return
	}

	res = rres.Cook()

	return
}

func Query(dq, name string) (res model.Result, err error) {
	return QueryN(dq, name, 1)
}

func QueryN(dq, name string, pageNum int) (res model.Result, err error) {
	formData := NewQueryForm(dq, name)
	res, err = query(formData)
	res.Name = name

	if err != nil {
		return
	}
	if pageNum > 1 {
		completeChan := make(chan struct{}, pageNum-1)
		for i := 1; i < pageNum; i++ {
			formData = formData.NextNPage(1)
			go func(formData QueryForm) {
				defer func() { completeChan <- struct{}{} }()
				r, err := query(formData)
				if err != nil {
					return
				}
				res.FightRecords = append(res.FightRecords, r.FightRecords...)
			}(formData)
		}
		for i := 1; i < pageNum; i++ {
			<-completeChan
		}
	}
	res.CalWinDefeat()
	return
}
