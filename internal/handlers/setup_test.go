package handlers

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/justinas/nosurf"
	"github.com/stasxgaming/bookings/internal/config"
	"github.com/stasxgaming/bookings/internal/models"
	"github.com/stasxgaming/bookings/internal/render"
)

var app config.AppConfig
var sessions *scs.SessionManager
var pathToTemplates = "./../../templates"
var functions = template.FuncMap{}

func getRoutes() http.Handler {

	gob.Register(models.Reservation{})

	app.SecureMode = false
	sessions = scs.New()
	sessions.Lifetime = 30 * time.Minute
	sessions.Cookie.Persist = true
	sessions.Cookie.SameSite = http.SameSiteLaxMode
	sessions.Cookie.Secure = app.SecureMode

	app.Session = sessions

	tc, err := CreateTestTemplateCashe()
	if err != nil {
		log.Fatal("cannot create tmpl cashe")
	}

	render.GetNewTemplates(&app)
	app.TemplateCashe = tc

	r := NewRepo(&app)
	NewHandlers(r)

	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	//mux.Use(CreateCSRFToken)
	mux.Use(SessionLoad)
	mux.Get("/", Repo.Home)
	mux.Get("/about", Repo.About)
	mux.Get("/lux", Repo.Lux)
	mux.Get("/regular", Repo.Regular)
	mux.Get("/reservation", Repo.Reservation)
	mux.Post("/reservation", Repo.PostReservation)
	mux.Post("/reservation-json", Repo.ReservationJson)

	mux.Get("/make-reservation", Repo.MRes)
	mux.Post("/make-reservation", Repo.PostMRes)
	mux.Get("/reservation-summary", Repo.ReservationSummary)

	fileserver := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileserver))
	return mux
}

func CreateCSRFToken(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.SecureMode,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

func SessionLoad(next http.Handler) http.Handler {
	return sessions.LoadAndSave(next)
}

func CreateTestTemplateCashe() (map[string]*template.Template, error) {
	pageCashe := make(map[string]*template.Template)

	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))
	if err != nil {
		return pageCashe, err
	}

	for _, page := range pages {

		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return pageCashe, err
		}

		layouts, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
		if err != nil {
			return pageCashe, err
		}

		if len(layouts) > 0 {
			ts, err := ts.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
			if err != nil {
				return pageCashe, err
			}

			pageCashe[name] = ts
		}
	}
	return pageCashe, nil
}
