package main

import (
	"fmt"
	"log"
)

func main() {
	str, err := userProfile("a12")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(str)
}

func fetchUserInfo(id string) (int, error) {
	return 1123, nil
	// return 0, fmt.Errorf("bad data")
}

func userProfile(id string) (string, error) {
	data, err := fetchUserInfo(id)
	if err != nil {
		return "", fmt.Errorf("fetch error: %w", err)
	}

	money := float64(data) / 100
	return fmt.Sprintf("Пользователь с id %s имеет на счету %.2f руб.", id, money), nil

}
