package main

import (
	"net/http"
	"github.com/justinas/nosurf"
	"github.com/jkannad/spas/members/internal/helper"
)
//NoSurf adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.IsProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

//SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

//Auth validates user session, if user doesn't have valid session. It routes back to login page.
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		if !helper.IsLoginRoute(r) && !helper.IsAuthenticated(r) {
			session.Put(r.Context(), "error", "Log in first")
			http.Redirect(w, r, "/member/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}