package main

import (
	"net/http"
)

func (app *App) Routes() *http.ServeMux {

	// Declare a serve mux and define the routes in exactly the same as before.
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/snippet", app.ShowSnippet)
	mux.HandleFunc("/snippet/new", app.NewSnippet)

	// Use the app.StaticDir field as the location of the static file directory.
	fileServer := http.FileServer(http.Dir(app.StaticDir))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Return the serve mux.
	return mux

}
