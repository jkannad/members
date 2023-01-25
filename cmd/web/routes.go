package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jkannad/spas/members/pkg/config"
	"github.com/jkannad/spas/members/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Home)
	mux.Get("/about", handlers.About)
	mux.Get("/member/search", handlers.MemberSearch)
	mux.Get("/member/register", handlers.MemberRegister)
	return mux
}
