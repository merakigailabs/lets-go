package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"
	"text/template"
	"time"

	"snippetbox.mergakigai.com/internal/models"

	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"
	_ "github.com/go-sql-driver/mysql"
)

type config struct {
	addr      string
	staticDir string
	dsn       string
}

// Add a new sessionManager field to the application struct.
type application struct {
	logger         *slog.Logger
	cfg            config
	snippets       *models.SnippetModel
	templateCache  map[string]*template.Template
	formDecoder    *form.Decoder
	sessionManager *scs.SessionManager
}

func main() {

	// Config from flags
	var cfg config
	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.StringVar(&cfg.dsn, "dsn", "user:password@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	// Logging
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	db, err := openDB(cfg.dsn)

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// Initialize a decoder instance...
	formDecoder := form.NewDecoder()

	// Use he scs.New() function to initialze a new session manager. Then we
	// configure it to use our MySQL database as the session store, and set a
	// lifetime of 12 hours (so that sessions automaticaaly expire 12 hours)
	// after first being created).
	sessionManger := scs.New()
	sessionManger.Store = mysqlstore.New(db)
	sessionManger.Lifetime = 12 * time.Hour

	// And add it to the application dependencies.
	app := &application{
		logger:         logger,
		cfg:            cfg,
		snippets:       &models.SnippetModel{DB: db},
		templateCache:  templateCache,
		formDecoder:    formDecoder,
		sessionManager: sessionManger,
	}

	logger.Info("starting server", "addr", cfg.addr)

	err = http.ListenAndServe(cfg.addr, app.routes())

	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
