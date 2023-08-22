package main

import (
	"hidden-record-assistant/server/module/cmdx"
	"hidden-record-assistant/server/module/httpx"
	"hidden-record-assistant/server/module/zlog"
)

func main() {
	resMap, err := cmdx.ExecWmicToMap()
	if err != nil {
		zlog.Error(err.Error())
		return
	}
	httpx.InitClient(resMap["remoting-auth-token"], resMap["riotclient-app-port"])
}
