package application

import (
	"fmt"
	"net/http"
)

type App struct {
	router *http.ServeMux
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}

func NewApp() *App {
	app := &App{
		router: http.NewServeMux(),
	}
	app.BuildRoutes()
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", app)

	return app
}
