package main

import (
	"fmt"
	"bytes"

	// use the fast version of the net/http package
	"github.com/valyala/fasthttp"
)

var (
	assetsPrefix = []byte("/web")
	staticFilesHandler = fasthttp.FSHandler("./web/", 2)
)

// Define the router function to handle request paths and match
// with the appropriate handle
func router(ctx *fasthttp.RequestCtx) {
	// Declare the handler, that routes requests to their respective filename.
	// The fileserver is wrapped in the `stripPrefix` method, because we want to
	// remove the "/assets/" prefix when looking for files.
	// For example, if we type "/assets/index.html" in our browser, the file server
	// will look for only "index.html" inside the directory declared above.
	// If we did not strip the prefix, the file server would look for
	// "./assets/assets/index.html", and yield an error
	// staticFileHandler := fasthttp.PathRewriteFunc(ctx.Path())

	path := ctx.Path()

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
