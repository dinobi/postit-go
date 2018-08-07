package main

import (
	"github.com/valyala/fasthttp"
	"encoding/json"
	"log"
)
// UserResponse represents an success object sent back to the client.
// when an a request is successful
type UserResponse struct {
	Data map[string]string `json:"data"`
	Message string `json:"message"`
	Type ResponseType `json:"type"`
}

// NewUserResponse represents the data sent back to
// a user when request is successful
func NewUserResponse() *UserResponse {
	ur := UserResponse{}
	ur.Data = map[string]string{}
	ur.Type = "success"
	return &ur
}

func userResponseHandler (ctx *fasthttp.RequestCtx, data *User, message string, status int) {
	userResponse := NewUserResponse()
	ctx.SetStatusCode(status)
	// Use a switch statement to match based on error code
	switch status {
	case 200:
		// TODO: add object to return to use on successful events
		userResponse.Message = fasthttp.StatusMessage(200)
		sendData(ctx, userResponse)
	case 201:
		userResponse.Data["username"] = data.Username
		userResponse.Data["email"] = data.Email
		userResponse.Message = message
		sendData(ctx, userResponse)

	default:
		userResponse.Message = fasthttp.StatusMessage(status)
		sendData(ctx, userResponse)
	}
}

// sendError: This function handles sending of json error response
// to the user
func sendData(ctx *fasthttp.RequestCtx, userResponse *UserResponse) {
	res, err := json.Marshal(userResponse)
		if err != nil {
			log.Fatal(err)
		}
		ctx.Write(res)
}