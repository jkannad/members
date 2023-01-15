package config

import (
	"html/template"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the application configuration settings
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	IsProduction  bool
	Session       *scs.SessionManager
}
