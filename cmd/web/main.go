package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	"snippetbox.mergakigai.com/internal/models"

	_ "github.com/go-sql-driver/mysql"
)

type config struct {
	addr      string
	staticDir string
	dsn       string
}

// Add a snippets field to the application struct. This will allow us to
// make the SnippetModel object available to our handlers.
type application struct {
	logger   *slog.Logger
	cfg      config
	snippets *models.SnippetModel
}

func main() {

	// Config from flags
	var cfg config
	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Path to static assets")
	// Define a new command-line flag for the MySQL DSN string.
	flag.StringVar(&cfg.dsn, "dsn", "user:password@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	// Logging
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	// To keep the main() function tidy, I've put the code for creating a connection pool
	// into the separate openDB() function below. We pass openDB() the DSN - Database string name from the cmd line flag

	db, err := openDB(cfg.dsn)

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// We also defer a call to db.Close(), so that the connection pool is close,
	// before the main() function exists.
	defer db.Close()

	// Initialize a models.SnippetModel instance containing the connection pool
	// and add it to the application dependencies
	app := &application{
		logger:   logger,
		cfg:      cfg,
		snippets: &models.SnippetModel{DB: db},
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
