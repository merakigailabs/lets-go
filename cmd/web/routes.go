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

	// Leave the static files route unchanged.
	fileServer := http.FileServer(http.Dir(app.cfg.staticDir))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// Create a new middleware chain containing the middleware specific to our
	// dynamic application routes. For now, this chain will only contain the
	// LoadAndSave session middleware but we'll add more to it later.
	dynamic := alice.New(app.sessionManager.LoadAndSave)

	mux.Handle("GET /{$}", dynamic.ThenFunc(app.home))
	mux.Handle("GET /snippet/view/{id}", dynamic.ThenFunc(app.snippetView))
	mux.Handle("GET /snippet/create", dynamic.ThenFunc(app.snippetCreate))
	mux.Handle("POST /snippet/create", dynamic.ThenFunc(app.snippetCreatePost))

	// Create a middleware chain containing our 'standard' middleware
	// which will be used for every request our application receives.
	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)
	// Return the 'standard' middleware chain followed by the servemux.
	return standard.Then(mux)

}
