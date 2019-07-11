package main

import (
	"github.com/alexedwards/scs"
	"sinistra/snippetbox/models"
)

// Define an App struct to hold the application-wide dependencies and configuration
// settings for our web application. For now we'll only include a HTMLDir field
// for the path to the HTML templates directory, but we'll add more to it as our build progresses.
// Add a new StaticDir field to our application dependencies.

type App struct {
	Database  *models.Database
	HTMLDir   string
	Sessions  *scs.SessionManager
	StaticDir string
}
