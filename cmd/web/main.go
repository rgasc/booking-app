package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/rgasc/booking-app/internal/config"
	"github.com/rgasc/booking-app/internal/driver"
	"github.com/rgasc/booking-app/internal/handlers"
	"github.com/rgasc/booking-app/internal/helpers"
	"github.com/rgasc/booking-app/internal/models"
	"github.com/rgasc/booking-app/internal/render"
)

const portNumber = ":1337"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
		return
	}

	defer db.SQL.Close()

	fmt.Printf("Starting application on port %s", portNumber)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func run() (*driver.DB, error) {
	// change to true when in production
	app.InProduction = false
	app.UseCache = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.InfoLog = infoLog
	app.ErrorLog = errorLog

	gob.Register(models.Reservation{})

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// connect to database
	log.Println("Connecting to database...")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=bookings user=postgres password=dbadmin123")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to database")

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("can not create template cache")
		return nil, err
	}

	app.TemplateCache = tc

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	render.App = &app
	helpers.NewHelpers(&app)

	return db, nil
}
