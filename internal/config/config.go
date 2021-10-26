package config

import (
	"text/template"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	TemplateCashe map[string]*template.Template
	SecureMode    bool
	Session       *scs.SessionManager
}
