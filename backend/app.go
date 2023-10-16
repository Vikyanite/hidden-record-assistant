package backend

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"hidden-record-assistant/backend/model"
	"hidden-record-assistant/backend/module/zlog"
	"hidden-record-assistant/backend/service"
)

// WailsApp struct
type WailsApp struct {
	ctx        context.Context
	lcuApp     *service.App
	FileLoader *FileLoader
}

// NewApp creates a new WailsApp application struct
func NewApp() (app *WailsApp) {
	app = &WailsApp{
		FileLoader: NewFileLoader(),
		lcuApp:     service.DefaultApp,
	}
	return
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *WailsApp) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *WailsApp) InitBackend() (err error) {
	keepalive, err := a.lcuApp.Start()
	if err != nil {
		zlog.Errorf("start err: %v", err)
		return
	}
	zlog.Info("start success")
	go func() {
		err := <-keepalive
		zlog.Errorf("keepalive error: %v", err)
		runtime.EventsEmit(a.ctx, "disconnected")
	}()
	return
}

func (a *WailsApp) GetCurrentSummoner() (data model.DisplaySummoner, err error) {
	data, err = a.lcuApp.GetCurrentSummoner()
	return
}

func (a *WailsApp) GetSummonerByName(name string) (data model.DisplaySummoner, err error) {
	zlog.Debugf("GetSummonerByName: %s", name)
	data, err = a.lcuApp.GetSummonerByName(name)
	return
}

func (a *WailsApp) GetMatchRecordsByPuuid(puuid string, beg int, end int) (data []model.DisplayMatch, err error) {
	zlog.Debugf("GetMatchRecordsByPuuid: %s, %d, %d", puuid, beg, end)
	data, err = a.lcuApp.GetDisplayMatchRecordsByPuuid(puuid, beg, end)
	return
}
