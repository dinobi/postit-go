package main

import (
	"net/http"
	// package for handling routing via mutliplexer
	"github.com/gorilla/mux"
)

// Lets create a router constructor that creates the router and
// returns it to us. We can now use this function
// to instantiate and test the router outside of the main function
func newPostitRouter() *mux.Router {
	// The mux routher is useful because it allows us to declare
	// paths that handler methods will be valid for
	r := mux.NewRouter()
	r.HandleFunc("/", apiHandler).Methods("GET")

	// Declare the static file directory and point it to the
	// directory we just made
	staticFileDirectory := http.Dir("./assets/")
	// Declare the handler, that routes requests to their respective filename.
	// The fileserver is wrapped in the `stripPrefix` method, because we want to
	// remove the "/assets/" prefix when looking for files.
	// For example, if we type "/assets/index.html" in our browser, the file server
	// will look for only "index.html" inside the directory declared above.
	// If we did not strip the prefix, the file server would look for
	// "./assets/assets/index.html", and yield an error
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	// The "PathPrefix" method acts as a matcher, and matches all routes starting
	// with "/assets/", instead of the absolute route itself
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	return r
}