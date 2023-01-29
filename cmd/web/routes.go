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
	mux.Post("/member/upsert/v1", handlers.UpsertMember)
	mux.Post("/member/search/v1", handlers.SearchMember)
	mux.Get("/member/get/v1", handlers.GetMember)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
