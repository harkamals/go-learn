package main

import (
	"fmt"
	"github.com/harkamals/go-learn/cmd"
)

func main() {
	fmt.Println("Hello World")
	u := cmd.User{
		Name:  "Siman",
		Email: "simarsingh.sg@gmail.com",
		Age:   6,
	}

	fmt.Println(u)

}
