package main

import (
	"regexp"
	"strings"

	"github.com/valyala/fasthttp"
)

func invalidCredentials(ctx *fasthttp.RequestCtx, userData User) (invalid bool) {

	invalid = false

	var (
		name         = strings.TrimSpace(userData.Username)
		email        = strings.TrimSpace(userData.Email)
		password     = strings.TrimSpace(userData.Password)
		emailRE      = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		alphanumeric = regexp.MustCompile("^[a-zA-Z0-9_]*$")
		message      = ""
	)

	if name == "" {
		message = "username field cannot be empty"
		errorResponseHandler(ctx, message, 400)
		invalid = true
		return
	}

	if len(name) < 3 {
		message = "username should be atleast 3 characters long"
		errorResponseHandler(ctx, message, 400)
		invalid = true
		return
	}
	if len(name) > 18 {
		message = "username should not exceed 18 characters"
		errorResponseHandler(ctx, message, 400)
		invalid = true
		return
	}
	if !alphanumeric.MatchString(name) {
		message = "username can contain only alphabets, numbers, and underscore"
		errorResponseHandler(ctx, message, 400)
		invalid = true
		return
	}
	if email == "" {
		message = "email field cannot be empty"
		errorResponseHandler(ctx, message, 400)
		invalid = true
		return
	}
	if !emailRE.MatchString(email) {
		message = "Enter a valid email"
		errorResponseHandler(ctx, message, 400)
		invalid = true
		return
	}
	if password == "" {
		message = "password field cannot be empty"
		errorResponseHandler(ctx, message, 400)
		invalid = true
		return
	}
	if len(password) < 6 {
		message = "password should be up to 6 characters long"
		errorResponseHandler(ctx, message, 400)
		invalid = true
		return
	}
	return
}
