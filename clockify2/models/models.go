package models

// User schema of the user table
type User struct {
	user_id   int    `json:"user_id"`
	email     string `json:"email"`
	user_name string `json:"user_name"`
	password  string `json:"password"`
}
