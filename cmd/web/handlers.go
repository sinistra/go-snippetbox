package main

import (
	"net/http"
	"strconv"
)

// Change the signature of our Home handler so it is defined as a method against *App.
func (app *App) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.NotFound(w) // Use the app.NotFound() helper.
		return
	}

	// Fetch a slice of the latest snippets from the database.
	snippets, err := app.Database.LatestSnippets()
	if err != nil {
		app.ServerError(w, err)
		return
	}
	// Pass the slice of snippets to the "home.page.html" templates.
	app.RenderHTML(w, r, "home.page.html", &HTMLData{
		Snippets: snippets,
	})

}

// Change the signature of our ShowSnippet handler so it is defined as a method // against App.
func (app *App) ShowSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.NotFound(w) // Use the app.NotFound() helper.
		return
	}

	snippet, err := app.Database.GetSnippet(id)
	if err != nil {
		app.ServerError(w, err)
		return
	}
	if snippet == nil {
		app.NotFound(w)
		return
	}
	// Render the show.page.html template, passing in the snippet data wrapped in our HTMLData struct.
	app.RenderHTML(w, r, "show.page.html", &HTMLData{
		Snippet: snippet})
}

// Change the signature of our NewSnippet handler so it is defined as a method against *App.
func (app *App) NewSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display the new snippet form..."))
}
