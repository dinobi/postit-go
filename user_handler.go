package main

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

func createUser(ctx *fasthttp.RequestCtx) {
	// create an instance of the user struct
	user := &User{}

	// get the post request body
	formData := ctx.PostBody()

	// convert from bytes into a struct as specified by user struct object
	if err := json.Unmarshal(formData, &user); err != nil {
		errorResponseHandler(ctx, 500)
	}

	// Append our existing list of birds with a new entry
	users = append(users, *user)
}

func getUsers(ctx *fasthttp.RequestCtx) {
	/*
	 TODO: check if user has a valid token
	*/

	//Convert the "users" variable to json
	UsersListBytes, err := json.Marshal(users)
	// If there is an error, print it to the console, and return a server
	// error response to the user
	if err != nil {
		errorResponseHandler(ctx, 500)
		return
	}
	// If all goes well, write the JSON list of birds to the response
	ctx.Write(UsersListBytes)
}