package query

import (
	"encoding/json"
	"hidden-record-assistant/model"
	"io"
	"net/http"
)

func SendQuery(dq, name string) (res model.Result, err error) {
	formData := NewQueryForm(dq, name)
	resp, err := http.PostForm("http://www.sslol.top/api/lol.php?act=cx", formData)
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
	res.Name = name
	return
}
