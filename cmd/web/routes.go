package main

import (
	"net/http"

	"github.com/justinas/alice"
)

// Update the signature for the routes() method so that it return a
// http.Handler instead of *http.ServeMux
func (app *application) routes() http.Handler {

	// HTTP Handlers
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(app.cfg.staticDir))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	// Create a middleware chain containing our 'standard' middleware
	// which will be used for every request our application receives.
	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)
	// Return the 'standard' middleware chain followed by the servemux.
	return standard.Then(mux)

}
