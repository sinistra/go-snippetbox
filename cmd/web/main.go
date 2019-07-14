package main

import (
	"database/sql"
	"flag"
	"github.com/alexedwards/scs"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sinistra/snippetbox/models"
	"time"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "root:root@/snippetbox?parseTime=true", "MySQL DSN")
	htmlDir := flag.String("html-dir", "./ui/html", "Path to HTML templates")
	// Define a new command-line flag for the session secret (a random key which
	// will be used to encrypt and authenticate session cookies). It should be 32
	// characters long.
	//secret := flag.String("secret", "s6Nd%+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret key")
	staticDir := flag.String("static-dir", "./ui/static", "Path to static assets")
	tlsCert := flag.String("tls-cert", "./tls/cert.pem", "Path to TLS certificate")
	tlsKey := flag.String("tls-key", "./tls/key.pem", "Path to TLS key")

	flag.Parse()

	// To keep the main() function tidy I've put the code for creating a connection
	// pool into the separate connect() function below. We pass connect() the DSN
	// from the command-line flag.
	db := connect(*dsn)
	// We also defer a call to db.Close(), so that the connection pool is closed
	// before the main() function exits.
	defer db.Close()

	// Use the scs.NewCookieManager() function to initialize a new session manager,
	// passing in the secret key as the parameter. Then we configure it so the
	// session always expires after 12 hours and sessions are persisted across
	// browser restarts.
	sessionManager := scs.New()
	sessionManager.Lifetime = 12 * time.Hour
	sessionManager.Cookie.Persist = true

	// Add the *staticDir value to our application dependencies.
	app := &App{
		Addr:      *addr,
		Database:  &models.Database{db},
		HTMLDir:   *htmlDir,
		Sessions:  sessionManager,
		StaticDir: *staticDir,
		TLSCert:   *tlsCert,
		TLSKey:    *tlsKey,
	}

	// Pass the app.Routes() method (which returns a serve mux) to the
	// http.ListenAndServe() function.

	// Call the new RunServer() method to start the server.
	app.RunServer()

}

// The connect() function wraps sql.Open() and returns a sql.DB connection pool for a given DSN.
func connect(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Set the maximum number of concurrently open connections.
	// Setting this to less than or equal to 0 will mean there is no maximum limit.
	// If the maximum number of open connections is reached and a new connection is needed,
	// Go will wait until until one of the connections is freed and becomes idle.
	// From a user perspective, this means their HTTP request will hang until a connection
	// is freed.
	db.SetMaxOpenConns(95)
	// Set the maximum number of idle connections in the pool.
	// Setting this to less than or equal to 0 will mean that no idle connections are retained.
	db.SetMaxIdleConns(5)

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}
