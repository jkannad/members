package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/jkannad/spas/members/pkg/config"
	"github.com/jkannad/spas/members/pkg/models"
)

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Template is not found in the cache")
	}

	buf := new(bytes.Buffer)

	err := t.Execute(buf, td)
	if err != nil {
		log.Println("Error executing template", err)
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println("Error writing to response writer")
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {

	var templateCache = make(map[string]*template.Template)

	templates, err := filepath.Glob("./templates/*.tmpl")
	if err != nil {
		log.Println("Error getting template file names", err)
		return templateCache, err
	}

	for _, tmpl := range templates {
		name := filepath.Base(tmpl)
		t, err := template.New(name).ParseFiles(tmpl)
		if err != nil {
			log.Println("Error parsing template", err)
			return templateCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			log.Println("Error finding layout templates", err)
			return templateCache, err
		}
		if len(matches) > 0 {
			t, err = t.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				log.Println("Error parsing template with layout", err)
				return templateCache, err
			}
		}
		templateCache[name] = t
	}

	return templateCache, nil

}
