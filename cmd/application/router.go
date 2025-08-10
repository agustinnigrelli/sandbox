package application

import (
	"net/http"

	"github.com/agustinnigrelli/sandbox/api/handlers"
)

func (app *App) BuildRoutes() {
	app.router.Handle("/ping", http.HandlerFunc(handlers.Ping)).methods("GET")

}
