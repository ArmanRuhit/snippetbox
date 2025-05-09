package main

import (
	"fmt"
	"html/template"
	"log"
	"log/slog"
	"net/http"
	"strconv"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "GO")
	// Initialize a slice containing the paths to the two files. It's important to note that the file containing our base template must be the *fist* file in the slice.
	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		// log.Print("Failed to parse template for method home(): ", err.Error())
		// Beacause this home handler is now a method against the application struct it can access its fields, including the structured logger. We will use this to create a log entry at the error level containing the error message, also including the request method and URI as attributes to assist with debugging.
		app.logger.Error("Failed to parse template for method home()", slog.String("method", r.Method), slog.String("uri", r.URL.RequestURI()), slog.String("error", err.Error()))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		// log.Print("Failed to execute template for method home(): ", err.Error())
		// New log process
		app.logger.Error("Failed to execute template for method home()", slog.String("method", r.Method), slog.String("uri", r.URL.RequestURI()), slog.String("error", err.Error()))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// SnippetView add a snippetView handler function
// Change the signature of the snippetView handler function to accept a pointer to the application struct as a parameter. This allows us to access the application-wide dependencies (like the logger) from within the handler function.
func (app *application) SnippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		app.logger.Error("Invalid snippet ID", slog.String("method", r.Method), slog.String("uri", r.URL.RequestURI()), slog.String("error", err.Error()))
		http.NotFound(w, r)
		return
	}

	// Use the fmt.Sprintf() function to interpolate the id value with a message, then write it as the HTTP response
	//msg := fmt.Sprintf("Display a specific snippet with ID %d...", id)

	_, errMsg := fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
	if errMsg != nil {
		// log.Fatal("Failed to write response for method snippetView(): ", errMsg)
		// New log process
		app.logger.Error("Failed to write response for method snippetView()", slog.String("method", r.Method), slog.String("uri", r.URL.RequestURI()), slog.String("error", errMsg.Error()))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// SnippetCreate add a snippetCreate handler function
func (app *application) SnippetCreate(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Display a form for creating a new snippet..."))
	if err != nil {
		log.Fatal("Failed to write response for method snippetCreate(): ", err)
		return
	}
}

// SnippetCreatePost Add a snippetCreatePost handler function
func (app *application) SnippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "GO")
	w.WriteHeader(http.StatusCreated)
	_, err := w.Write([]byte("Save a new snippet..."))
	if err != nil {
		// log.Fatal("Failed to write response for method snippetCreatePost(): ", err)
		// New log process
		app.logger.Error("Failed to write response for method snippetCreatePost()", slog.String("method", r.Method), slog.String("uri", r.URL.RequestURI()), slog.String("error", err.Error()))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
