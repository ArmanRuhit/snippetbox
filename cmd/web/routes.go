package main

import "net/http"

// The routes() method returns a servemux containing our application routes
func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.Home)
	mux.HandleFunc("GET /snippet/view/{id}", app.SnippetView)
	mux.HandleFunc("GET /snippet/create", app.SnippetCreate)
	mux.HandleFunc("POST /snippet/create", app.SnippetCreatePost)

	return mux
}