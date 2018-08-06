package main

// User represents Postit's User object.
// This account object has support for Social Auth
type User struct{
	Id float64
	Provider string `json:"provider"`
	UID interface{} `json:"uid"`
	Username string `json:"username"`
	Bio string `json:"bio"`
	Password interface{} `json:"password"`
	Email string `json:"email"`
	Picture string `json:"picture"`
}

// declare users to be an array of user structs
var users []User