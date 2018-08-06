package main

import (
	"github.com/valyala/fasthttp"
	"encoding/json"
	"log"
)

// ErrorResponse represents an error object sent back to the client.
// when an error occurs
type ErrorResponse struct {
	Message string `json:"message"`
	Status int `json:"status"`
}

func errorResponseHandler (ctx *fasthttp.RequestCtx, message string, status int) {
	errorResponse := &ErrorResponse{}
	ctx.SetStatusCode(status)
	// Use a switch statement to match based on error code
	switch status {
	case 400:
		errorResponse.Message = message
		errorResponse.Status = 400
		sendError(ctx, errorResponse)
	
	case 401:
		errorResponse.Message = message
		errorResponse.Status = 401
		sendError(ctx, errorResponse)

	case 403:
		errorResponse.Message = message
		errorResponse.Status = 403
		sendError(ctx, errorResponse)

	case 404:
		errorResponse.Message = message
		errorResponse.Status = 404
		sendError(ctx, errorResponse)

	case 409:
		errorResponse.Message = message
		errorResponse.Status = 409
		sendError(ctx, errorResponse)

	case 422:
		errorResponse.Message = message
		errorResponse.Status = 422
		sendError(ctx, errorResponse)

	default:
		errorResponse.Message = fasthttp.StatusMessage(status)
		errorResponse.Status = status
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