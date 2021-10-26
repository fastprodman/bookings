package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/stasxgaming/bookings/internal/config"
	"github.com/stasxgaming/bookings/internal/forms"
	"github.com/stasxgaming/bookings/internal/models"
	"github.com/stasxgaming/bookings/internal/render"
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

	render.RenderTemplate(w, r, "main.page.tmpl", &models.TemplateData{
		StrMap: stMap,
	})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	var stMap = make(map[string]string)
	stMap["test"] = "That is about page"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StrMap: stMap,
	})
}

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "reservation.page.tmpl", &models.TemplateData{})
}

func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start-date")
	end := r.Form.Get("end-date")
	w.Write([]byte(fmt.Sprintf("Starting date is %s Ending date is %s", start, end)))
}

func (m *Repository) Regular(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "regular.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Lux(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "lux.page.tmpl", &models.TemplateData{})
}

func (m *Repository) MRes(w http.ResponseWriter, r *http.Request) {
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation
	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

func (m *Repository) PostMRes(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first-name"),
		LastName:  r.Form.Get("last-name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}

	form := forms.New(r.PostForm)
	//form.Has("first-name", r)
	form.Required("first-name", "last-name", "email")
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	m.App.Session.Put(r.Context(), "reservation", reservation)
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

type JsonResponce struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func (m *Repository) ReservationJson(w http.ResponseWriter, r *http.Request) {
	respo := JsonResponce{
		OK:      true,
		Message: "Validated!",
	}

	out, err := json.MarshalIndent(respo, "", "    ")
	if err != nil {
		log.Println(err)
	}

	log.Println(string(out))
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(out))
}

func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		log.Println("cannot get item from session")
		m.App.Session.Put(r.Context(), "error", "Can't get reservation from session")
		http.Redirect(w,r,"/", http.StatusTemporaryRedirect)
		return
	}
	data := make(map[string]interface{})
	data["reservation"] = reservation

	render.RenderTemplate(w, r, "reservation-summary.page.tmpl", &models.TemplateData{
		Data: data,
	})

}
