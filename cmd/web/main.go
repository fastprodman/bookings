package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/stasxgaming/bookings/pkg/config"
	"github.com/stasxgaming/bookings/pkg/handlers"
	"github.com/stasxgaming/bookings/pkg/render"
)

const Addr string = ":8080"

var app config.AppConfig
var sessions *scs.SessionManager

func main() {
	app.SecureMode = false
	sessions = scs.New()
	sessions.Lifetime = 30 * time.Minute
	sessions.Cookie.Persist = true
	sessions.Cookie.SameSite = http.SameSiteLaxMode
	sessions.Cookie.Secure = app.SecureMode

	app.Session = sessions

	tc, err := render.CreateTemplateCashe()
	if err != nil {
		log.Fatal(err)
	}

	render.GetNewTemplates(&app)
	app.TemplateCashe = tc

	r := handlers.NewRepo(&app)
	handlers.NewHandlers(r)

	srv := &http.Server{
		Addr:    Addr,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
