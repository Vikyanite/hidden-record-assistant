package backend

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"hidden-record-assistant/backend/model"
	"hidden-record-assistant/backend/service"
	"hidden-record-assistant/backend/service/support"
)

// WailsApp struct
type WailsApp struct {
	ctx        context.Context
	conn       *support.Connector
	supportApp *service.App
	FileLoader *support.FileLoader
}

// NewApp creates a new WailsApp application struct
func NewApp(conn *support.Connector) (app *WailsApp) {
	app = &WailsApp{
		conn:       conn,
		FileLoader: support.NewFileLoader(conn),
		supportApp: service.NewApp(conn),
	}
	return
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *WailsApp) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *WailsApp) InitBackend() (data model.InitBackendData, err error) {
	// 所有的功能都依赖于conn的连接，所以先初始化conn
	data.Auth, err = a.conn.Init(func() { runtime.EventsEmit(a.ctx, "disconnected") })
	if err != nil {
		return
	}

	// 初始化集成辅助app
	err = a.supportApp.Init()
	if err != nil {
		return
	}

	data.CurrentSummoner, err = a.supportApp.GetCurrentSummoner()
	return
}