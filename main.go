package main

import (
	"embed"
	"hello-wails/internal/repository"
	"hello-wails/internal/service"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	userRepo := repository.NewMemoryUserRepo()
	userSvc := service.NewUserService(userRepo)
	calcSvc := service.NewCalculatorService()
	app := NewApp(userSvc, calcSvc)

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "hello-wails",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
