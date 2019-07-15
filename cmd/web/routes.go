package main

import (
	"github.com/bmizerany/pat"
	"net/http"
)

func (app *App) Routes() http.Handler {

	// Declare a serve mux and define the routes in exactly the same as before.
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.Home))
	mux.Get("/snippet/new", app.RequireLogin(http.HandlerFunc(app.NewSnippet)))
	mux.Post("/snippet/new", app.RequireLogin(http.HandlerFunc(app.CreateSnippet)))
	mux.Get("/snippet/:id", NoSurf(app.ShowSnippet))
	mux.Get("/user/signup", NoSurf(app.SignupUser))
	mux.Post("/user/signup", NoSurf(app.CreateUser))
	mux.Get("/user/login", NoSurf(app.LoginUser))
	mux.Post("/user/login", NoSurf(app.VerifyUser))
	mux.Post("/user/logout", app.RequireLogin(http.HandlerFunc(app.LogoutUser)))

	// Use the app.StaticDir field as the location of the static file directory.
	fileServer := http.FileServer(http.Dir(app.StaticDir))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	// Pass the router as the 'next' parameter to the LogRequest middleware.
	// Because LogRequest() is just a function, and the function returns a
	// http.Handler we don't need to do anything else.
	return LogRequest(SecureHeaders(mux))

}
