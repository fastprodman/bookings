package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/justinas/nosurf"
	"github.com/stasxgaming/bookings/internal/config"
	"github.com/stasxgaming/bookings/internal/models"
)

var app *config.AppConfig
var pathToTemplates = "./templates"

func GetNewTemplates(a *config.AppConfig) {
	app = a
}

var functions = template.FuncMap{}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.CSRFToken = nosurf.Token(r)
	return td
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {
	tempCash := app.TemplateCashe

	t, ok := tempCash[tmpl]
	if !ok {
		log.Fatal("not ok")
	}

	byteBuf := new(bytes.Buffer)
	td = AddDefaultData(td, r)

	err := t.Execute(byteBuf, td)
	if err != nil {
		log.Println("error executing bytes buffer")
	}

	_, _ = byteBuf.WriteTo(w)
}

func CreateTemplateCashe() (map[string]*template.Template, error) {
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
