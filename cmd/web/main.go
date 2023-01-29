package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/jkannad/spas/members/internal/config"
	"github.com/jkannad/spas/members/internal/handlers"
	"github.com/jkannad/spas/members/internal/render"
)

const PortNumber = ":1709"

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

	if err != nil {
		log.Fatal("Cannot create template cache", err)
	}

	app.TemplateCache = templateCache
	app.UseCache = false

	render.NewTemplate(&app)
	handlers.SetAppConfig(&app)

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
