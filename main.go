package main

import (
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"hidden-record-assistant/backend"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := backend.NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:         "hidden-record-assistant",
		Width:         1280,
		Height:        720,
		DisableResize: true,
		AssetServer: &assetserver.Options{
			Assets:  assets,
			Handler: app.FileLoader,
		},

		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 100},
		OnStartup:        app.Startup,
		Bind: []interface{}{
			app,
		},
		//Windows: &windows.Options{
		//	WebviewIsTransparent: true,
		//	WindowIsTranslucent:  true,
		//},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
