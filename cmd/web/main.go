package main

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/stasxgaming/bookings/internal/config"
	"github.com/stasxgaming/bookings/internal/handlers"
	"github.com/stasxgaming/bookings/internal/helpers"
	"github.com/stasxgaming/bookings/internal/models"
	"github.com/stasxgaming/bookings/internal/render"
	"github.com/stasxgaming/bookings/internal/driver"
)

const Addr string = ":8080"

var app config.AppConfig
var sessions *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	srv := &http.Server{
		Addr:    Addr,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func run() (*driver.DB, error) {
	gob.Register(models.Reservation{})

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	app.SecureMode = false
	sessions = scs.New()
	sessions.Lifetime = 30 * time.Minute
	sessions.Cookie.Persist = true
	sessions.Cookie.SameSite = http.SameSiteLaxMode
	sessions.Cookie.Secure = app.SecureMode

	app.Session = sessions

	//connectind to database
	dsn := "host=localhost port=5432 dbname=bookings user=postgres password=11111"
	db, err := driver.ConnectSQL(dsn)
	if err != nil{
		log.Fatal("Cannot connect to DB")
	}

	tc, err := render.CreateTemplateCashe()
	if err != nil {
		return nil, err
	}

	render.GetNewTemplates(&app)
	app.TemplateCashe = tc

	r := handlers.NewRepo(&app, db)
	handlers.NewHandlers(r)

	helpers.NewHelpers(&app)

	return db, nil
}
