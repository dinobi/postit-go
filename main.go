// This is the name of our package
// Everything with this package name can see everything
// else inside the same package, regardless of the file they are in
package main

import (
	"log"
	"github.com/valyala/fasthttp"
)

const (
	port string = ":8080"
)

func main() {
	// The "HandleFunc" method accepts a path and a function as arguments
	// (Yes, we can pass functions as arguments, and even treat them like variables in Go)
	// However, the handler function has to have the appropriate signature
	// (as described by the "handler" function below)
	// http.HandleFunc("/", apiHandler) //this is how Go server handles it's endpoints but
	// for application we will be using an external package that helps treat different paths
	// in different ways

	// When an error occurs, we print the error and exit the app (os.Exit(1))
	// ListenAndServe starts an HTTP server with a given address and handler.
	// The handler is usually nil, which means to use DefaultServeMux.
	// but in our case we are using the gorille mux
	log.Fatal(fasthttp.ListenAndServe(port, router))
}
