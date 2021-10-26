package models

import "github.com/stasxgaming/bookings/internal/forms"

type TemplateData struct {
	StrMap    map[string]string
	IntMap    map[string]int
	FlotMap   map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
	Form      *forms.Form
}
