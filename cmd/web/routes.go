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
	mux.Get("/generals-quarters", handlers.Repo.GeneralsHandler)
	mux.Get("/majors-suite", handlers.Repo.MajorsHandler)
	mux.Get("/search-availability", handlers.Repo.AvailabilityHandler)
	mux.Get("/contact", handlers.Repo.ContactHandler)
	mux.Get("/make-reservation", handlers.Repo.MakeReservationHandler)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	return mux
}
