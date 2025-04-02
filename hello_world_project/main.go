package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	name := "Михаил"
	fmt.Println(generateCompliment(name))
}

func generateCompliment(name string) string {
	amount := 3
	key := rand.IntN(amount)

	compliments := map[int]string{
		0: "Ты великолепен, ",
		1: "У тебя потрясающая улыбка, ",
		2: "Ты вдохновляешь, "}

	return fmt.Sprintf("%s%s!", compliments[key], name)
}
