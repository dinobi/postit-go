// This is a test for the main package
package main

import (
	"log"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestRouter(t *testing.T) {
	port := 3000
	defer startServerOnPort(t, port, router).Close()

	// Perpare a client, which fetches webpages via HTTP proxy listening
	// on the localhost:8080.
	c := &fasthttp.HostClient{
		Addr: "localhost:3000",
	}

	statusCode, body, err := c.Get(nil, "http://localhost:3000/")

	if err != nil {
		log.Fatalf("Error when visiting root route: %s", err)
	}

	// Start our assertions
	// 1) We want our status to be 200 (ok)
	if statusCode != fasthttp.StatusOK {
		log.Fatalf("Unexpected status code: %d. Expecting %d", statusCode, fasthttp.StatusOK)
	}

	// 2) Check that the response body is what we expect.
	expected := `Welcome to Postit`
	actual := string(body)
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestRouterForNonExistentRoute(t *testing.T) {
	port := 3002
	defer startServerOnPort(t, port, router).Close()

	// Perpare a client, which fetches webpages via HTTP proxy listening
	// on the localhost:8080.
	c := &fasthttp.HostClient{
		Addr: "localhost:3002",
	}

	statusCode, body, err := c.Get(nil, "http://google.com/foo/bar")

	if err != nil {
		log.Fatalf("Error when visiting root route: %s", err)
	}

	// Start our assertions
	// 1) We want our status to be 404 (Not found)
	if statusCode != 404 {
		log.Fatalf("Unexpected status code: %d. Expecting %d", 404, statusCode)
	}

	// 2) Check that the response body is what we expect.
	expected := `Unsupported path`
	actual := string(body)
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}

	// Create an instance of the args struct to client post request
	myArgs := &fasthttp.Args{}
	// Fetch foobar page via local proxy. Reuse body buffer.
	statusCode, body, err = c.Post(nil, "http://foobar.com/google/com", myArgs)

	if err != nil {
		log.Fatalf("Error when visiting root route: %s", err)
	}

	// Start our assertions
	// 1) We want our status to be 404 (Not found)
	if statusCode != 404 {
		log.Fatalf("Unexpected status code: %d. Expecting %d", 404, statusCode)
	}
}

func TestStaticFileServer(t *testing.T) {
	port := 3002
	defer startServerOnPort(t, port, router).Close()

	// Perpare a client, which fetches webpages via HTTP proxy listening
	// on the localhost:8080.
	c := &fasthttp.HostClient{
		Addr: "localhost:3002",
	}

	statusCode, _, err := c.Get(nil, "http://localhost:3002/web")

	if err != nil {
		log.Fatalf("Error when visiting root route: %s", err)
	}

	// Start our assertions
	// 1) We want our status to be 200 (ok)
	if statusCode != fasthttp.StatusOK {
		log.Fatalf("Unexpected status code: %d. Expecting %d", statusCode, fasthttp.StatusOK)
	}

	// 2) It isn't wise to test the entire content of the HTML file.
	// Instead, we test that the content-type header is "text/html; charset=utf-8"
	// so that we know that an html file has been served
	// contentType := c.Name
	// expectedContentType := "text/html; charset=utf-8"

	// if contentType != expectedContentType {
	// 	t.Errorf("Wrong content type, expected %s, got %s", expectedContentType, contentType)
	// }
}
