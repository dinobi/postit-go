package main

import (
	"github.com/valyala/fasthttp"
	"encoding/json"
)

func userHandler(ctx *fasthttp.RequestCtx) {
	//Convert the "users" variable to json
	UsersListBytes, err := json.Marshal(users)

	// If there is an error, print it to the console, and return a server
	// error response to the user
	if err != nil {
		errorResponseHandler(500, ctx)
		return
	}
	// If all goes well, write the JSON list of birds to the response
	ctx.Write(UsersListBytes)
}