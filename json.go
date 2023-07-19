package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	jsonStr := `{"name":"John", "email":"john@example.com"}`
	var user User
	json.Unmarshal([]byte(jsonStr), &user)
	fmt.Printf("Name: %s, Email: %s", user.Name, user.Email)
}
