// This is a test for the main package
package main

import (
	"log"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestGetUser(t *testing.T) {
	port := 3000
	defer startServerOnPort(t, port, router).Close()

	// Perpare a client, which fetches webpages via HTTP proxy listening
	// on the localhost:8080.
	c := &fasthttp.HostClient{
		Addr: "localhost:3000",
	}

	statusCode, body, err := c.Get(nil, "http://localhost:3000/users")

	if err != nil {
		log.Fatalf("Error when visiting root route: %s", err)
	}

	// Start our assertions
	// 1) We want our status to be 200 (ok)
	if statusCode != fasthttp.StatusOK {
		log.Fatalf("Unexpected status code: %d. Expecting %d", statusCode, fasthttp.StatusOK)
	}

	// 2) Check that the response body is what we expect.
	expected := "null"
	actual := string(body)
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

// func TestCreateUser(t *testing.T) {
// 	port := 3001
// 	defer startServerOnPort(t, port, router).Close()

// 	// Perpare a client, which fetches webpages via HTTP proxy listening
// 	// on the localhost:8080.
// 	c := &fasthttp.HostClient{
// 		Addr: "localhost:3001",
// 	}

// 	// Create an instance of the args struct to client post request
// 	userData := &fasthttp.Args{}
// 	statusCode, body, err := c.Post(nil, "http://localhost:3001/create-user", userData)

// 	if err != nil {
// 		log.Fatalf("Error when visiting root route: %s", err)
// 	}

// 	fmt.Println(userData)
// 	// Start our assertions
// 	// 1) We want our status to be 200 (ok)
// 	if statusCode != fasthttp.StatusOK {
// 		log.Fatalf("Unexpected status code: %d. Expecting %d", statusCode, fasthttp.StatusOK)
// 	}

// 	// 2) Check that the response body is what we expect.
// 	// expected := "null"
// 	// actual := string(body)
// 	// if actual != expected {
// 	// 	t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
// 	// }

// 	useResponseBody(body)
// }
