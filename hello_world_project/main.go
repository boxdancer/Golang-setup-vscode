package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

func main() {
	str, err := UserProfileToString("asd", -30)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(str)
}

func UserProfileToString(name string, age int) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	if age < 0 {
		return "", errors.New("negative age")
	}
	
	name = strings.TrimSpace(name)
	if name == "" {
		return "", errors.New("name cannot contain only spaces")
	}

	result := fmt.Sprintf("Имя человека: %s, возраст: %d.", name, age)
	return result, nil
}
