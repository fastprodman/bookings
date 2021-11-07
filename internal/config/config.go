package config

import (
	"log"
	"text/template"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	TemplateCashe map[string]*template.Template
	SecureMode    bool
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	Session       *scs.SessionManager
}
