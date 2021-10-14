package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/stasxgaming/bookings/pkg/config"
	"github.com/stasxgaming/bookings/pkg/models"
)

var app *config.AppConfig

func GetNewTemplates(a *config.AppConfig) {
	app = a
}

var functions = template.FuncMap{}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	tempCash := app.TemplateCashe

	t, ok := tempCash[tmpl]
	if !ok {
		log.Fatal("not ok")
	}

	byteBuf := new(bytes.Buffer)
	td = AddDefaultData(td)

	err := t.Execute(byteBuf, td)
	if err != nil {
		log.Println("Shit happend")
	}

	_, _ = byteBuf.WriteTo(w)
}

func CreateTemplateCashe() (map[string]*template.Template, error) {
	pageCashe := make(map[string]*template.Template)

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return pageCashe, err
	}

	for _, page := range pages {

		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return pageCashe, err
		}

		layouts, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return pageCashe, err
		}

		if len(layouts) > 0 {
			ts, err := ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return pageCashe, err
			}

			pageCashe[name] = ts
		}
	}
	return pageCashe, nil
}
