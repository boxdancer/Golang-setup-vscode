package main

import (
	"fmt"
)

func main() {
	city := "Москва"
	temp := 15
	weather := "солнечно"

	weatherStr := fmt.Sprintf("В городе %s температура %d°C, %s.", city, temp, weather)

	fmt.Println(weatherStr)
}
