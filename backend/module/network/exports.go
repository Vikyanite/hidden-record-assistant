package network

import (
	"hidden-record-assistant/backend/module/zlog"
	"io"
)

func Get(url string) (data string, err error) {
	defer func() {
		if err != nil {
			zlog.Errorf("Get %s: %s", url, err.Error())
			return
		}
	}()
	resp, err := DefaultConnector.Get(url)
	if err != nil {
		return
	}
	binData, err := io.ReadAll(resp.Body)

	data = string(binData)

	zlog.Debugf("Get %s: %s", url, data)
	return
}
