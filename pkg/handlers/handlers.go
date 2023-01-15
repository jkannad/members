package handlers

import (
	"net/http"

	"github.com/jkannad/spas/members/pkg/config"
	"github.com/jkannad/spas/members/pkg/models"
	"github.com/jkannad/spas/members/pkg/render"
)

var appConfig *config.AppConfig

func SetAppConfig(app *config.AppConfig) {
	appConfig = app
}

func Home(w http.ResponseWriter, r *http.Request) {
	appConfig.Session.Put(r.Context(), "remote_ip", r.RemoteAddr)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func About(w http.ResponseWriter, r *http.Request) {
	remoteIP := appConfig.Session.GetString(r.Context(), "remote_ip")
	stringMap := make(map[string]string)
	stringMap["father"] = "Janarthanan Kannadhasan"
	stringMap["mother"] = "Vaidegi Rajagopalan"
	stringMap["son"] = "Ashwin Janarthanan"
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
