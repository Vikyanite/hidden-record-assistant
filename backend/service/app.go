package service

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"hidden-record-assistant/backend/model"
)

// WailsApp struct
type WailsApp struct {
	ctx  context.Context
	conn *HTTPConnector
}

var DefaultApp = NewApp()

// NewApp creates a new WailsApp application struct
func NewApp() *WailsApp {
	return &WailsApp{
		conn: NewHTTPConnector(),
	}
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *WailsApp) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *WailsApp) InitConnector() (model.InitBackendData, error) {
	return a.conn.Init(func() {
		runtime.EventsEmit(a.ctx, "disconnected")
	})
}

func (a *WailsApp) GetCurrentSummoner() ([]byte, error) {
	return a.conn.Get("/lol-summoner/v1/current-summoner")
}

func (a *WailsApp) SearchSummonerByName(name string) ([]byte, error) {
	return a.conn.Get("/lol-summoner/v1/summoners?name=" + name)
}
