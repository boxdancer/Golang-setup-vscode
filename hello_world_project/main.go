package main

import (
	"fmt"
)

func main() {
	age := 10
	role := "user"
	status := "active"

	var flag bool

	if role == "admin" || role == "moderator" || age >= 18 && role == "user" && status == "active" {
		flag = true
	} else {
		flag = false
	}
	fmt.Println(flag)
}
