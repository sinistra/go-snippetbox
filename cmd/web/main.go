package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	htmlDir := flag.String("html-dir", "./ui/html", "Path to HTML templates")
	staticDir := flag.String("static-dir", "./ui/static", "Path to static assets")

	flag.Parse()

	// Add the *staticDir value to our application dependencies.
	app := &App{
		HTMLDir:   *htmlDir,
		StaticDir: *staticDir,
	}

	// Pass the app.Routes() method (which returns a serve mux) to the // http.ListenAndServe() function.

	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, app.Routes())
	log.Fatal(err)

}
