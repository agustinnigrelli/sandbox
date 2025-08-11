package application

import (
	"log"
	"net/http"

	"github.com/agustinnigrelli/sandbox/utils/web"
)

type App struct {
	router *web.Router
}

const (
	host = ":8080"
)

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}

func NewApp() *App {
	app := &App{
		router: web.NewRouter(),
	}
	routeGroup := app.router.Group("/sandbox/v1")
	app.BuildRoutes(routeGroup)
	log.Println("Server running on port", host)
	http.ListenAndServe(host, app)

	return app
}
