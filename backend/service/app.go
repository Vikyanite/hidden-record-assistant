package service

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// WailsApp struct
type WailsApp struct {
	ctx  context.Context
	conn *HTTPConnector
}

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

func (a *WailsApp) InitConnector() error {
	return a.conn.Init(func() {
		runtime.EventsEmit(a.ctx, "disconnected")
	})
}

func (a *WailsApp) CurrentSummoner() (string, error) {
	return a.conn.Get("/lol-summoner/v1/current-summoner")
}
