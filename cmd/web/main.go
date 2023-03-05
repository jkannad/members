package main

import (
	"os" 
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/jkannad/spas/members/internal/config"
	"github.com/jkannad/spas/members/internal/handlers"
	"github.com/jkannad/spas/members/internal/render"
	"github.com/jkannad/spas/members/internal/helper"
)

const PortNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	app.IsProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.IsProduction

	app.Session = session

	templateCache, err := render.CreateTemplateCache()
	formFieldConfig := config.BuildFormFieldConfigs()

	if err != nil {
		log.Fatal("Cannot create template cache", err)
	}

	app.TemplateCache = templateCache
	app.UseCache = false
	app.FormFieldConfig = formFieldConfig
	app.InfoLogger = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.ErrorLogger = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	render.NewTemplate(&app)
	handlers.SetAppConfig(&app)
	helper.New(&app)


	srv := &http.Server{
		Addr:    PortNumber,
		Handler: routes(&app),
	}
	fmt.Println("Application started at port :1709")
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server", err)
	}
}
