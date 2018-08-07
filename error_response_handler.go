package main

import (
	"github.com/valyala/fasthttp"
	"encoding/json"
	"log"
)

type ResponseType string

// ErrorResponse represents an error object sent back to the client.
// when an error occurs
type ErrorResponse struct {
	Message string `json:"message"`
	Type ResponseType `json:"type"`
}

// errorResponseHandler acts like constructor for an error response object
// to be sent to the client when an error occurs in the request
func errorResponseHandler (ctx *fasthttp.RequestCtx, message string, status int) {
	errorResponse := &ErrorResponse{}
	ctx.SetStatusCode(status)
	errorResponse.Type = "error"

	if status < 500 {
		errorResponse.Message = message
		sendError(ctx, errorResponse)
	} else {
		errorResponse.Message = fasthttp.StatusMessage(status)
		sendError(ctx, errorResponse)
	}
}

// sendError: This function handles sending of json error response
// to the user
func sendError(ctx *fasthttp.RequestCtx, errorResponse *ErrorResponse) {
	res, err := json.Marshal(errorResponse)
		if err != nil {
			log.Fatal(err)
		}
		ctx.Write(res)
}