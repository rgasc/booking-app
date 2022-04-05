package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/rgasc/booking-app/internal/config"
	"github.com/rgasc/booking-app/internal/handlers"
	"github.com/rgasc/booking-app/internal/models"
	"github.com/rgasc/booking-app/internal/render"
)

const portNumber = ":1337"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	// change to true when in production
	app.InProduction = false

	gob.Register(models.Reservation{})

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("can not create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	handlers.NewRepo(&app)
	render.App = &app

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
