package main

import (
	"fmt"
	"strings"
)

func main() {
	var message string = " Go - это не просто язык, это СТИЛЬ ЖИЗНИ!         "

	messageNoSpace := strings.TrimSpace(message)
	fmt.Println(messageNoSpace)

	messageLower := strings.ToLower(messageNoSpace)
	fmt.Println(messageLower)

	if strings.HasPrefix(messageLower, "go") {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
