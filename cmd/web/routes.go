package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/jkannad/spas/members/internal/config"
	"github.com/jkannad/spas/members/internal/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Search)
	mux.Get("/member/about", handlers.About)
	mux.Get("/member/search", handlers.Search)
	mux.Get("/member/register", handlers.Register)
	mux.Get("/member/getmember/{id:[0-9]+}", handlers.GetMember)
	mux.Get("/member/getallmembers", handlers.GetAllMembers)
	mux.Get("/member/getstates/{id}", handlers.GetStates)
	mux.Get("/member/getcities/{country}/{state}", handlers.GetCities)
	mux.Get("/member/getdialcode/{id}", handlers.GetDialCode)
	mux.Get("/member/login", handlers.Login)


	
	mux.Post("/member/login", handlers.PostLogin)
	mux.Post("/member/upsert", handlers.UpsertMember)
	mux.Post("/member/search/result", handlers.SearchResult)
	

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
