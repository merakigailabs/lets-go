package main

import "net/http"

func (app *application) routes() *http.ServeMux {

	// HTTP Handlers
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(app.cfg.staticDir))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// Swap the route declarations to use the application struct's methods as the
	// handler functions. (because we defined the methods against the struct in handlers.go)
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	return mux
}
