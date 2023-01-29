package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jkannad/spas/members/internal/config"
	"github.com/jkannad/spas/members/internal/models"
	"github.com/jkannad/spas/members/internal/render"
)

var appConfig *config.AppConfig

func SetAppConfig(app *config.AppConfig) {
	appConfig = app
}

func About(w http.ResponseWriter, r *http.Request) {
	//remoteIP := appConfig.Session.GetString(r.Context(), "remote_ip")
	stringMap := make(map[string]string)
	stringMap["title"] = "About"
	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func Search(w http.ResponseWriter, r *http.Request) {
	//remoteIP := appConfig.Session.GetString(r.Context(), "remote_ip")
	stringMap := make(map[string]string)
	stringMap["title"] = "Search Member"
	render.RenderTemplate(w, r, "search.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func Register(w http.ResponseWriter, r *http.Request) {
	//remoteIP := appConfig.Session.GetString(r.Context(), "remote_ip")
	stringMap := make(map[string]string)
	stringMap["title"] = "Register Member"
	render.RenderTemplate(w, r, "register.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

type upsertResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func UpsertMember(w http.ResponseWriter, r *http.Request) {
	response := upsertResponse{
		Ok:      true,
		Message: "Member's details were saved successfully",
	}
	out, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func GetMember(w http.ResponseWriter, r *http.Request) {
	//remoteIP := appConfig.Session.GetString(r.Context(), "remote_ip")
	stringMap := make(map[string]string)
	stringMap["title"] = "Update Member"
	render.RenderTemplate(w, r, "update.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func SearchMember(w http.ResponseWriter, r *http.Request) {
	//remoteIP := appConfig.Session.GetString(r.Context(), "remote_ip")
	stringMap := make(map[string]string)
	stringMap["title"] = "Update Member"
	render.RenderTemplate(w, r, "update.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
