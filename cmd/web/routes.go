package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kure37/bookings/pkg/config"
	"github.com/kure37/bookings/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Use(Nosurf)
	mux.Use(middleware.Recoverer)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	return mux

}
