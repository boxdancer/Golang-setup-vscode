package main

import (
	"fmt"
)

func main() {
	num := 15

	numStr := fmt.Sprintf("Запись числа %d в разных системах счисления:\nДесятичная: %d\nДвоичная: %b\nВосьмеричная: %o\nШестнадцатеричная: %X", num, num, num, num, num)
	fmt.Println(numStr)
}
