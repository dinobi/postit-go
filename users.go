package main

import (
	"time"
)

// User represents Postit's User object.
// This account object has support for Social Auth
type User struct {
	ID       string `json:"id"`
	Provider string `json:"provider"`
	UID      string `json:"uid"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Password  string
	Email     string      `json:"email"`
	Picture   string      `json:"picture"`
	CreatedAt time.Time   `json:"created_at"`
}
// Password interface represents group of methods used to
// confirm password entires, encode and decode password
// supplied by users
type Password interface {
	// encodePassword() (string, error)
	// decodePassword() (string, error)
}

// func (up *userPassword) encodePassword() (string, error) {

// }

// func (up *userPassword) decodePassword() (string, error) {
	
// }

// declare users to be an array of user structs
var users []User