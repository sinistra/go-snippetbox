package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// Change the signature of our Home handler so it is defined as a method against *App.
func (app *App) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.NotFound(w) // Use the app.NotFound() helper.
		return
	}

	app.RenderHTML(w, "home.page.html") // Use the app.RenderHTML() helper.

}

// Change the signature of our ShowSnippet handler so it is defined as a method // against App.
func (app *App) ShowSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.NotFound(w) // Use the app.NotFound() helper.
		return
	}

	fmt.Fprintf(w, "Display a specific snippet (ID %d)...", id)

}

// Change the signature of our NewSnippet handler so it is defined as a method against *App.
func (app *App) NewSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display the new snippet form..."))
}
