package main

import (
	"github.com/bmizerany/pat"
	"net/http"
)

func (app *App) Routes() http.Handler {

	// Declare a serve mux and define the routes in exactly the same as before.
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.Home))
	mux.Get("/snippet/new", http.HandlerFunc(app.NewSnippet))
	mux.Post("/snippet/new", http.HandlerFunc(app.CreateSnippet))
	mux.Get("/snippet/:id", http.HandlerFunc(app.ShowSnippet))

	// Use the app.StaticDir field as the location of the static file directory.
	fileServer := http.FileServer(http.Dir(app.StaticDir))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	// Return the serve mux.
	return mux

}
