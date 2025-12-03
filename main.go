package main

import (
	"context"
	"embed"

	"ankigen/internal/api"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	extractorAPI := api.NewExtractorAPI(10, 50)

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "ankigen",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			extractorAPI.Startup(ctx)
		},
		Bind: []interface{}{
			extractorAPI,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
