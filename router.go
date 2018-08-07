package main

import (
	"fmt"
	"bytes"

	// use the fast version of the net/http package
	"github.com/valyala/fasthttp"
)

var (
	assetsPrefix = []byte("/web")
	// We want to remove the "./web/" prefix when looking for files.
	// For example, if we type "/web/index.html" in our browser, the file server
	// will look for only "index.html" inside the directory declared above.
	// If we did not strip the prefix by 2, the file server would look for
	// "./web/web/index.html", and yield an error
	staticFilesHandler = fasthttp.FSHandler("./web/", 2)
)

// Define the router function to handle request paths and match
// with the appropriate handle
func router(ctx *fasthttp.RequestCtx) {
	path := ctx.Path()

	// The fileserver method is invoked when a user routes to assetsPrefix
	// path("/web") and then routes requests to their respective filename
	if bytes.HasPrefix(path, assetsPrefix) {
		staticFilesHandler(ctx)
	} else {
		switch string(path){
		case "/":
			fmt.Fprintf(ctx, "Welcome to Postit")
		case "/users":
			getUsers(ctx)
		case "/create-user":
			signup(ctx)
		default:
			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		}
	}
}
