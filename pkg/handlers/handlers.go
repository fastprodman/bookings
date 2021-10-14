package handlers

import (
	"net/http"

	"github.com/stasxgaming/bookings/pkg/config"
	"github.com/stasxgaming/bookings/pkg/models"
	"github.com/stasxgaming/bookings/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	var stMap = make(map[string]string)
	stMap["test"] = "That is Home page"

	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)

	render.RenderTemplate(w, "main.page.tmpl", &models.TemplateData{
		StrMap: stMap,
	})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	var stMap = make(map[string]string)
	stMap["test"] = "That is about page"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StrMap: stMap,
	})
}
