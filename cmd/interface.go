package cmd

import (
	"fmt"
)

// User ...
type User struct {
	Name  string
	Email string
	Age   int
}

// Notify ...
func (u *User) Notify() {
	fmt.Println("User notified: ", u.Name, u.Email)
}

