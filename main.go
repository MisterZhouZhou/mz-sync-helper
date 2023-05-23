package main

import (
	"embed"
	"mz-sync-helper/server"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

var (
	port = 27149
)

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// 启动服务
	go server.Run(port, assets)

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "mz-sync-helper",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
