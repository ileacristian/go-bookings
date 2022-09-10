package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/ileacristian/go-bookings/pkg/config"
	"github.com/ileacristian/go-bookings/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.HomeHandler)
	mux.Get("/about", handlers.Repo.AboutHander)

	return mux
}