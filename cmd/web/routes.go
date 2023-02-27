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
	mux.Use(Auth)

	mux.Get(config.ROOT, handlers.Search)
	mux.Get(config.GET_ABOUT, handlers.About)
	mux.Get(config.GET_SEARCH, handlers.Search)
	mux.Get(config.GET_REGISTER, handlers.Register)
	mux.Get(config.GET_MEMBER_BY_ID, handlers.GetMember)
	mux.Get(config.GET_ALL_MEMBERS, handlers.GetAllMembers)
	mux.Get(config.GET_STATES_BY_COUNTRY, handlers.GetStates)
	mux.Get(config.GET_CITIES_BY_COUNTRY_AND_STATE, handlers.GetCities)
	mux.Get(config.GET_DIAL_CODE_BY_COUNTRY, handlers.GetDialCode)
	mux.Get(config.GET_LOGIN, handlers.Login)
	mux.Get(config.GET_LOGOUT, handlers.Logout)

	mux.Post(config.POST_LOGIN, handlers.PostLogin)
	mux.Post(config.POST_UPSERT, handlers.UpsertMember)
	mux.Post(config.POST_SEARCH_RESULT, handlers.SearchResult)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
