package main

import (
	"fmt"
)

func main() {
	var weight, height float64

	fmt.Print("Введите ваш вес (кг): ")
	fmt.Scanln(&weight)
	fmt.Print("Введите ваш рост (см): ")
	fmt.Scanln(&height)

	massInsdex := weight / (height * height / 1e4)
	fmt.Printf("Ваш ИМТ: %.2f\n", massInsdex)

	switch {
	case massInsdex < 18.5:
		fmt.Println("Категория: Недостаточный вес")
	case massInsdex >= 25 && massInsdex < 29.9:
		fmt.Println("Категория: Избыточный вес")
	case massInsdex >= 30:
		fmt.Println("Категория: Ожирение")
	default:
		fmt.Println("Категория: Нормальный вес")
	}
}	
