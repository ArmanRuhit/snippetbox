package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello from Snippetbox"))
	if err != nil {
		log.Fatal("Failed to write response for method home(): ", err)
		return
	}
}

// add a snippetView handler function
func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// Use the fmt.Sprintf() function to interpolate the id value with a message, then write it as the HTTP response
	msg := fmt.Sprintf("Display a specific snippet with ID %d...", id)

	_, err_msg := w.Write([]byte(msg))
	if err_msg != nil {
		log.Fatal("Failed to write response for method snippetView(): ", err_msg)
		return
	}
}

// add a snippetCreate handler function
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Display a form for creating a new snippet..."))
	if err != nil {
		log.Fatal("Failed to write response for method snippetCreate(): ", err)
		return
	}
}

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux, then register the home function as the handler for the "/" url pattern
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/snippet/view/{id}", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// Print a log message to say that the server is starting
	log.Print("starting server on :4000")

	// Use the http.ListenAndServe() function to start a new web server. We pass in two parameters: the TCP network address to listen on (in this case ":4000") and the servemux we just created. if http.ListenAndServe() returns an error we use the log.Fatal() function to log the error message and exit. Note that any error returned by http.ListenAndServe() is always non-nil.
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
