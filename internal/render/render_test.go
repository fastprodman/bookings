package render

import (
	"net/http"
	"testing"

	"github.com/stasxgaming/bookings/internal/models"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	sessions.Put(r.Context(), "flash", "123")
	result := AddDefaultData(&td, r)
	if result.Flash != "123" {
		t.Error("not equal to expected 123 value")
	}

}

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "./../../templates"
	tc, err := CreateTemplateCashe()
	if err != nil {
		t.Error(err)
	}

	app.TemplateCashe = tc
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	var ww myWriter

	err = RenderTemplate(&ww, r, "main.page.tmpl", &models.TemplateData{})
	if err != nil {
		t.Error(err)
	}

	err = RenderTemplate(&ww, r, "not-existing.page.tmpl", &models.TemplateData{})
	if err == nil {
		t.Error("rendered not existing tameplate")
	}
}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "some-url", nil)
	if err != nil {
		return nil, err
	}

	ctx := r.Context()
	ctx, _ = sessions.Load(ctx, r.Header.Get("X-Session"))
	r = r.WithContext(ctx)

	return r, nil
}

func TestGetNewTemplates(t *testing.T) {
	GetNewTemplates(&testApp)
}

func TestCreateTemplateCashe(t *testing.T) {
	_, err := CreateTemplateCashe()
	if err != nil {
		t.Error(err)
	}
}
