package main

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"fmt"
)

func signup(ctx *fasthttp.RequestCtx) {
	// create an instance of the user struct
	user := &User{}
	message := ""

	// get the post request body
	formData := ctx.PostBody()
	fmt.Println(string(formData))
	// convert from bytes into a struct as specified by user struct object
	if err := json.Unmarshal(formData, &user); err != nil {
		errorResponseHandler(ctx, message, 500)
	}

	if (invalidCredentials(ctx, *user)) {
		return;
	}
	// Append our existing list of users with a new entry
	users = append(users, *user)
	message = "User account created successfully" 
	userResponseHandler(ctx, user, message, 201)
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
		message := ""
		errorResponseHandler(ctx, message, 500)
		return
	}
	// If all goes well, write the JSON list of birds to the response
	ctx.Write(UsersListBytes)
}
