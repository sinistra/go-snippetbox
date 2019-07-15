package main

import (
	"context"
	"net/http"
)

func (app *App) LoggedIn(r *http.Request) bool {
	// Load the session data for the current request, and use the Exists() method // to check if it contains a currentUserID key. This returns true if the
	// key is in the session data; false otherwise.
	session, _ := app.Sessions.Load(context.Background(), "user")
	loggedIn := session.Value("currentUserID")
	if loggedIn == "" {
		return false
	}
	return true
}
