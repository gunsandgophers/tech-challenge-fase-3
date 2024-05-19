package main

import "tech-challenge-fase-1/internal/adapter/driver/app"

func main() {
	app := app.NewAPIApp()
	app.Run()
	defer app.Shutdown()
}
