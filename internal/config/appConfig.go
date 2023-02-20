package config

import (
	"log"
	"html/template"
	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the application configuration settings
type AppConfig struct {
	UseCache      	bool
	TemplateCache 	map[string]*template.Template
	IsProduction  	bool
	Session       	*scs.SessionManager
	FormFieldConfig	map[string]map[string]FieldConfig
	InfoLogger		*log.Logger
	ErrorLogger		*log.Logger
}