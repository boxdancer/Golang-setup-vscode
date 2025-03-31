package main

import (
	"fmt"
	"strings"
)

func main() {
	priceList := map[string]int{
		"Клавиатура JZ9": 19200,
		"Наушники N45":   9600,
		"Смартфон S10":   55000}

	fmt.Print("Введите название товара: ")
	var request string
	fmt.Scanln(&request)

	key := findKey(request, priceList)

	price, exists := priceList[key]
	if exists {
		fmt.Printf("%v: %v\n", key, price)
	} else {
		fmt.Printf("Товар \"%v\" не найден.\n", request)
	}
}

func findKey(str string, priceList map[string]int) string {
	for key, _ := range priceList {
		if containsStr(key, str) {
			return key
		}
	}
	return "Не найдено"

}

func containsStr(str string, substr string) bool {
	return strings.Contains(strings.ToLower(str), strings.ToLower(substr))
}
