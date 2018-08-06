// This is a test for the main package
package main

import (
	"fmt"
	"io"
	"net"
	"testing"

	"github.com/valyala/fasthttp"
)

func startServerOnPort(t *testing.T, port int, h fasthttp.RequestHandler) io.Closer {
	ln, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		t.Fatalf("cannot start tcp server on port %d: %s", port, err)
	}
	go fasthttp.Serve(ln, h)
	return ln
}

func useResponseBody(body []byte) {
	// Do something with body :)
	fmt.Println(string(body))
}
