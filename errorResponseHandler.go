package main

import (
	"github.com/valyala/fasthttp"
)

// ErrorResponse represents an error object sent back to the client.
// when an error occurs
type ErrorResponse struct {
	Message string `json:"message"`
	Status int `json:"status"`
}

func errorResponseHandler (status int, ctx *fasthttp.RequestCtx) {
	errorResponse := &ErrorResponse{}
	// Use a switch statement to match based on error code
	switch status {
	case 400:
		errorResponse.Message = fasthttp.StatusMessage(400)
		errorResponse.Status = 400
	case 401:
		errorResponse.Message = fasthttp.StatusMessage(401)
		errorResponse.Status = 401
	case 403:
		errorResponse.Message = fasthttp.StatusMessage(403)
		errorResponse.Status = 401
	case 404:
		errorResponse.Message = fasthttp.StatusMessage(403)
		errorResponse.Status = 401
	case 409:
		errorResponse.Message = fasthttp.StatusMessage(409)
		errorResponse.Status = 401
	case 422:
		errorResponse.Message = fasthttp.StatusMessage(422)
		errorResponse.Status = 401
	default:
		fasthttp.StatusMessage(status)
	}
} 