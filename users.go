package main

import (
	"time"
)

// User represents Postit's User object.
// This account object has support for Social Auth
type User struct{
	ID string `json:"id"`
	Provider string `json:"provider"`
	UID string `json:"uid"`
	Username string `json:"username"`
	Bio string `json:"bio"`
	Password string `json:"password"`
	Email string `json:"email"`
	Picture string `json:"picture"`
	CreatedAt time.Time `json:"created_at"`
}

// declare users to be an array of user structs
var users []User