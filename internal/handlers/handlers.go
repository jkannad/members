package handlers

import (
	"fmt"
	"encoding/json"
	//"log"
	"net/http"
	//"errors"
	"strconv"

	"github.com/jkannad/spas/members/internal/config"
	"github.com/jkannad/spas/members/internal/models"
	"github.com/jkannad/spas/members/internal/render"
	"github.com/jkannad/spas/members/internal/helper"
	"github.com/jkannad/spas/members/internal/service"
	"github.com/go-chi/chi/v5"
)

const(
	COUNTRIES = "countries"
	STATES = "states"
	CITIES = "cities"
	DIAL_CODES ="dialcodes"
)

var appConfig *config.AppConfig


func SetAppConfig(app *config.AppConfig) {
	appConfig = app
}

func About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["title"] = "About"
	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["title"] = "Login"
	render.RenderTemplate(w, r, "login.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func Search(w http.ResponseWriter, r *http.Request) {
	//remoteIP := appConfig.Session.GetString(r.Context(), "remote_ip")
	stringMap := make(map[string]string)
	stringMap["title"] = "Search Member"
	countries, err := service.GetCountries()
	if err != nil {
		helper.ServerError(w, err)
		return
	}
	render.RenderTemplate(w, r, "search.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
		Countries: countries,
	})
}

func SearchResult(w http.ResponseWriter, r *http.Request) {
	//remoteIP := appConfig.Session.GetString(r.Context(), "remote_ip")
	stringMap := make(map[string]string)
	stringMap["title"] = "Search Result"
	search := models.Search {
		FirstName: r.FormValue("first_name"),
		LastName: r.FormValue("last_name"),
		Area: r.FormValue("area"),
		PostalCode: r.FormValue("postal_code"),
		Country: r.FormValue("country"),
		State: r.FormValue("state"),
		City: r.FormValue("city"),
		ContactNumber: r.FormValue("contact_number"),
		Email: r.FormValue("email"),
	}
	members, err := service.SearchMembers(search)
	if err != nil {
		helper.ServerError(w, err)
	}
	data := make(map[string]interface{})
	for _, member := range members {
		data[strconv.Itoa(member.Id)] = member
	}
	render.RenderTemplate(w, r, "search.result.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
		Data: data,
	})
}

func Register(w http.ResponseWriter, r *http.Request) {
	//remoteIP := appConfig.Session.GetString(r.Context(), "remote_ip")
	stringMap := make(map[string]string)
	stringMap["title"] = "Register Member"
	countries, err := service.GetCountries()
	if err != nil {
		helper.ServerError(w, err)
		return
	}

	dialCodes, err := service.GetDialCodes()
	if err != nil {
		helper.ServerError(w, err)
		return
	}

	render.RenderTemplate(w, r, "register.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
		Data: nil,
		Countries: countries,
		DialCodes: dialCodes,
	})
}

type response struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func UpsertMember(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	err := r.ParseForm()
	if err != nil {
		helper.ServerError(w, err)
		return
	}
	fmt.Println("Dial Code:", r.FormValue("dial_code"))
	member := models.Member {
		Title: r.FormValue("title"),
		FirstName: r.FormValue("first_name"),
		LastName: r.FormValue("last_name"),
		Gender: r.FormValue("gender"),
		Dob: r.FormValue("dob"),
		Doj: r.FormValue("doj"),
		Address1: r.FormValue("address1"),
		Address2: r.FormValue("address2"),
		Area: r.FormValue("area"),
		Country: r.FormValue("country"),
		State: r.FormValue("state"),
		City: r.FormValue("city"),
		PostalCode: r.FormValue("postal_code"),
		DialCode: r.FormValue("dial_code"),
		ContactNumber: r.FormValue("contact_number"),
		Email: r.FormValue("email"),
	}


	if r.FormValue("id") != "" {
		id, _ := strconv.Atoi(r.FormValue("id"))
		member.Id = id
	}

	validationErrors := config.ValidateFormData(&member, appConfig.FormFieldConfig["member"])
	
	var status bool
	var msg string
	if len(validationErrors) != 0 {
		status = false
		msg = getFormatedErrorResponse(validationErrors)
	} else {
		err = service.Upsert(&member)
		if err != nil {
			status = false
			msg = "Internal service error. Please try after sometime or get in touch with your adminstrators"
			helper.ServerError(w, err)
			return
		} else {
			status = true
			msg = "Member details were saved successfully"
		}
	}

	res := response{
		Ok:      status,
		Message: msg,
	}

	out, e := json.MarshalIndent(res, "", "    ")

	if e != nil {
		helper.ServerError(w, e)
		return
	}
	w.Write(out)
}

func GetMember(w http.ResponseWriter, r *http.Request) {
	//remoteIP := appConfig.Session.GetString(r.Context(), "remote_ip")
	stringMap := make(map[string]string)
	data := make(map[string]interface{})
	stringMap["title"] = "Update Member"
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	member, err := service.GetMember(id)
	if err != nil {
		helper.ServerError(w, err)
		return
	}
	countries, err := service.GetCountries()
	if err != nil {
		helper.ServerError(w, err)
		return
	}

	dialCodes, err := service.GetDialCodes()
	if err != nil {
		helper.ServerError(w, err)
		return
	}

	states, err := service.GetStates(member.Country)
	if err != nil {
		helper.ServerError(w, err)
		return
	}

	cities, err := service.GetCities(member.Country, member.State)
	if err != nil {
		helper.ServerError(w, err)
		return
	}

	data["member"] = member
	render.RenderTemplate(w, r, "register.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
		Data: data,
		Countries: countries,
		States: states,
		Cities: cities,
		DialCodes: dialCodes,
	})
}

func GetAllMembers(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["title"] = "Search Result"
	members, err := service.GetAllMembers()
	if err != nil {
		helper.ServerError(w, err)
	}
	data := make(map[string]interface{})
	for _, member := range members {
		data[strconv.Itoa(member.Id)] = member
	}

	render.RenderTemplate(w, r, "search.result.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
		Data: data,
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

func GetStates(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	states, err := service.GetStates(id)
	if err != nil {
		helper.ServerError(w, err)
		return
	}

	out, err := json.MarshalIndent(states, "", "    ")

	if err != nil {
		helper.ServerError(w, err)
		return
	}
	w.Write(out)
}

func GetCities(w http.ResponseWriter, r *http.Request) {
	countryCd := chi.URLParam(r, "country")
	stateCd := chi.URLParam(r, "state")
	cities, err := service.GetCities(countryCd, stateCd)
	if err != nil {
		helper.ServerError(w, err)
		return
	}
	out, err := json.MarshalIndent(cities, "", "    ")

	if err != nil {
		helper.ServerError(w, err)
		return
	}
	w.Write(out)
}

func GetDialCode(w http.ResponseWriter, r *http.Request) {
	countryCd := chi.URLParam(r, "id")
	dialCode, err := service.GetDialCode(countryCd)
	if err != nil {
		helper.ServerError(w, err)
		return
	}
	out, err := json.MarshalIndent(dialCode, "", "    ")

	if err != nil {
		helper.ServerError(w, err)
		return
	}
	w.Write(out)
}


func getFormatedErrorResponse(error map[string][]string) string {
	var msg string
	if len(error) != 0 {
		for _, value := range error {
			for _, val := range value {
				msg += fmt.Sprintf("%s\n", val)
			}
		}
	} 
	return msg
}
