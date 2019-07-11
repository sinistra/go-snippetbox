package main

import (
	"fmt"
	"net/http"
	"sinistra/snippetbox/pkg/forms"
	"strconv"
)

// Change the signature of our Home handler so it is defined as a method against *App.
func (app *App) Home(w http.ResponseWriter, r *http.Request) {
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
	// Pat doesn't strip the colon from the named capture key, so we need to
	// get the value of ":id" from the query string instead of "id".
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
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

	// Load the session data then use the PopString() method to retrieve the value
	// for the "flash" key. PopString() also deletes the key and value from the
	// session data, so it acts like a one-time fetch. If there is no matching
	// key in the session data it will return the empty string. If you want to
	// retrieve a string from the session and not delete it you should use the
	// GetString() method instead.
	session, _ := app.Sessions.Load(r.Context(), "flash")
	flash := app.Sessions.PopString(session, "flash")
	//io.WriteString(w, flash)

	// Render the show.page.html template, passing in the snippet data wrapped in our HTMLData struct.
	app.RenderHTML(w, r, "show.page.html", &HTMLData{
		Snippet: snippet,
		Flash:   flash,
	})
}

func (app *App) CreateSnippet(w http.ResponseWriter, r *http.Request) {
	// First we call r.ParseForm() which adds any POST (also PUT and PATCH) data
	// to the r.PostForm map. If there are any errors we use our
	// app.ClientError helper to send a 400 Bad Request response to the user.
	err := r.ParseForm()
	if err != nil {
		app.ClientError(w, http.StatusBadRequest)
		return
	}
	// We initialize a *forms.NewSnippet object and use the r.PostForm.Get() method
	// to assign the data to the relevant fields.
	form := &forms.NewSnippet{
		Title:   r.PostForm.Get("title"),
		Content: r.PostForm.Get("content"),
		Expires: r.PostForm.Get("expires"),
	}
	// Check if the form passes the validation checks. If not, then use the
	// fmt.Fprint function to dump the failure messages to the response body.
	if !form.Valid() {
		// Re-display the new.page.html template passing in the *forms.NewSnippet
		// object (which contains the validation failure messages and previously // submitted data).
		app.RenderHTML(w, r, "new.page.html", &HTMLData{Form: form})
		return
	}

	// If the validation checks have been passed, call our database model's
	// InsertSnippet() method to create a new database record and return it's ID value.
	id, err := app.Database.InsertSnippet(form.Title, form.Content, form.Expires)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	// Use session manager's Load() method to fetch the session data for the current
	// request. If there's no existing session for the current user (or their
	// session has expired) then a new, empty, session will be created. Any errors
	// are deferred until the session is actually used.
	session, _ := app.Sessions.Load(r.Context(), "flash")
	// Use the PutString() method to add a string value ("Your snippet was saved
	// successfully!") and the corresponding key ("flash") to the the session data.
	app.Sessions.Put(session, "flash", "Your snippet was saved successfully!")

	// If successful, send a 303 See Other response redirecting the user to the
	// page with their new snippet.
	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}

func (app *App) NewSnippet(w http.ResponseWriter, r *http.Request) {
	// Pass an empty *forms.NewSnippet object to the new.page.html template.
	// Because it's empty, it won't contain any previously submitted data or validation
	// failure messages.
	app.RenderHTML(w, r, "new.page.html", &HTMLData{
		Form: &forms.NewSnippet{},
	})
}
