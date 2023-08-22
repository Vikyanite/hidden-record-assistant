package main

import (
	"hidden-record-assistant/module/cmdx"
	"hidden-record-assistant/module/httpx"
	"hidden-record-assistant/module/zlog"
)

func main() {
	resMap, err := cmdx.ExecWmicToMap()
	if err != nil {
		zlog.Error(err.Error())
		return
	}
	httpx.InitClient(resMap["remoting-auth-token"], resMap["riotclient-app-port"])
}
