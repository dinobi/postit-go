// This is the name of our package
// Everything with this package name can see everything
// else inside the same package, regardless of the file they are in
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// The "HandleFunc" method accepts a path and a function as arguments
	// (Yes, we can pass functions as arguments, and even treat them like variables in Go)
	// However, the handler function has to have the appropriate signature
	// (as described by the "handler" function below)
	// http.HandleFunc("/", apiHandler) //this is how Go server handles it's endpoints but
	// for application we will be using an external package that helps treat different paths
	// in different ways

	// Create a postit router instance for handling routes
	router := newPostitRouter()

	// When an error occurs, we print the error and exit the app (os.Exit(1))
	// ListenAndServe starts an HTTP server with a given address and handler.
	// The handler is usually nil, which means to use DefaultServeMux.
	// but in our case we are using the gorille mux
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Define the apiHandler function to handle (req, res) objects
// It has to follow the function signature of a ResponseWriter
// and Request type as the arguments.
func apiHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Welcome to Postit") // write to res object instead of command line
}
