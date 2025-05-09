package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

// Define an application struct to hold the application-wide dependencies for the web application. For now, we only include the structured logger but we will add more to this as the build proogresses.
type application struct {
	logger *slog.Logger
}

func main() {

	// Define a new command-line flag with the name "addr", a default value of ":4000" and some short help text to explaining what the flag controls. The value of the flag will be stored in the addr variable at runtime.
	addr := flag.String("addr", ":4000", "HTTP network address")

	// Importantly, we use the flag.Parse() funtion to parse the command-line flags. This reads in the command-line flag value and assigns it to the addr variable. You need to call this *before* you use the addr variable, otherwise it wll always contain the default value of ":4000". If any errors are encountered during parsing, the application will be terminated.
	flag.Parse()

	// Use the slog.New() function to initialize a new structured logger, which writes to the standard out stream and uses default settings.
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
		AddSource: true,
	}))

	// Initialize a new instance of the application struct, passing in the logger as a dependency. This allows us to access the logger from within the application struct methods.
	app := &application{
		logger: logger,
	}

	// Use the http.NewServeMux() function to initialize a new servemux, then register the home function as the handler for the "/" url pattern
	mux := http.NewServeMux()

	// Create a file server which serves files out of the "./ui/static" directory.
	// Note that the path given to the http.Dir function is relative to the project
	// directory root.
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.Home)
	mux.HandleFunc("GET /snippet/view/{id}", app.SnippetView)
	mux.HandleFunc("GET /snippet/create", app.SnippetCreate)
	mux.HandleFunc("POST /snippet/create", app.SnippetCreatePost)

	// Print a log message to say that the server is starting
	// log.Printf("starting server on http://localhost%s", *addr)
	// Use the info method to log the starting server message at the Info severity (along with listen address as an attribute).
	logger.Info("starting server", slog.String("addr", *addr))

	// Use the http.ListenAndServe() function to start a new web server. We pass in two parameters: the TCP network address to listen on (in this case ":4000") and the servemux we just created. if http.ListenAndServe() returns an error we use the log.Fatal() function to log the error message and exit. Note that any error returned by http.ListenAndServe() is always non-nil.
	// Pass the dereferrenced addr pointer to http.ListenAndServe() so that the value of the flag is used, rather than the default value.
	err := http.ListenAndServe(*addr, mux)
	// log.Fatal(err)
	// And we also use the Error() method to log any error message returnedd by http.ListenAndServe() at the Error severity (along with the error message as an attribute). and then call os.Exit(1) to terminate the application with exit code 1.
	logger.Error(err.Error())
	os.Exit(1)

}
