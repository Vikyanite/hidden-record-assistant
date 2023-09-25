package lcu

import (
	"hidden-record-assistant/backend/service/lcu/internal"
)

type App = internal.App

var (
	DefaultApp = internal.NewApp(internal.NewConnector())
)
