package application

import (
	"github.com/agustinnigrelli/sandbox/api/handlers"
	"github.com/agustinnigrelli/sandbox/utils/web"
)

func (app *App) BuildRoutes(router *web.Router) {
	pingHandler := handlers.NewPingHandler()
	router.Get("/ping", pingHandler.Ping)
}
