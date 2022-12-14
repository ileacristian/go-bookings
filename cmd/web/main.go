package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/ileacristian/go-bookings/internal/config"
	"github.com/ileacristian/go-bookings/internal/handlers"
	"github.com/ileacristian/go-bookings/internal/models"
	"github.com/ileacristian/go-bookings/internal/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	gob.Register(models.Reservation{})

	app.UseCache = false
	app.InProduction = false
	app.Session = session

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("could not create template cache")
	}

	app.TemplateCache = templateCache

	repo := handlers.NewRepository(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	server := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	fmt.Printf("Listening on port %s\n", portNumber)
	err = server.ListenAndServe()

	log.Fatal(err)
}
