package backend

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"hidden-record-assistant/backend/model"
	"hidden-record-assistant/backend/module/zlog"
	"hidden-record-assistant/backend/service/lcu"
	"hidden-record-assistant/backend/service/support"
	"time"
)

// WailsApp struct
type WailsApp struct {
	ctx        context.Context
	lcuApp     *lcu.App
	FileLoader *support.FileLoader
}

// NewApp creates a new WailsApp application struct
func NewApp() (app *WailsApp) {
	app = &WailsApp{
		FileLoader: support.NewFileLoader(),
		lcuApp:     lcu.DefaultApp,
	}
	return
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *WailsApp) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *WailsApp) InitBackend() (data model.InitBackendData, err error) {
	defer func() func() {
		start := time.Now()
		return func() {
			if err != nil {
				return
			}
			zlog.Debugf("InitBackend Cost: %s", time.Since(start).String())
		}
	}()()
	var exitChan chan struct{}
	data.Auth, exitChan, err = a.lcuApp.Init()
	if err != nil {
		return
	}
	go func() {
		<-exitChan
		runtime.EventsEmit(a.ctx, "disconnected")
	}()
	return
}

func (a *WailsApp) GetCurrentSummoner() (data model.Summoner, err error) {
	data, err = a.lcuApp.GetCurrentSummoner()
	return
}

func (a *WailsApp) GetSummonerByName(name string) (data model.Summoner, err error) {
	zlog.Debugf("GetSummonerByName: %s", name)
	data, err = a.lcuApp.GetSummonerByName(name)
	return
}

func (a *WailsApp) GetMatchRecordsByPuuid(puuid string, beg int, end int) (data []model.DisplayMatch, err error) {
	zlog.Debugf("GetMatchRecordsByPuuid: %s, %d, %d", puuid, beg, end)
	data, err = a.lcuApp.GetDisplayMatchRecordsByPuuid(puuid, beg, end)
	return
}
