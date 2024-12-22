package main

import "net/http"

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

	// Pass rhe servermux as the 'next' parameter to the commonHeaders middleware.
	// Because commonHeaders is just a function, and the function returns a
	// http.Handler we don't need to do anything else.

	// Wrap the existing chain with the logRequest middleware.

	// Wrap the existing chain with the recoverPanic middleware.
	return app.recoverPanic(app.logRequest(commonHeaders(mux)))

}
