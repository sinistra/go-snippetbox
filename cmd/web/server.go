package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"time"
)

func (app *App) RunServer() {
	// Declare a tls.Config variable to hold the non-default TLS settings we want the server to use.
	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256}}
	// Initialize a new http.Server struct. We set the Addr and Handler so that
	// the server uses the same network address and routes as before, and we also set
	// the TLSConfig field to use the tlsConfig variable we just created.
	srv := &http.Server{
		Addr:           app.Addr,
		MaxHeaderBytes: 524288,
		Handler:        app.Routes(),
		TLSConfig:      tlsConfig,
		IdleTimeout:    time.Minute,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	// Call the http.Server's ListenAndServeTLS() method to start the server,
	// passing in the paths to the TLS certificate and corresponding private key.
	log.Printf("Starting server on %s", app.Addr)
	err := srv.ListenAndServeTLS(app.TLSCert, app.TLSKey)
	log.Fatal(err)
}
