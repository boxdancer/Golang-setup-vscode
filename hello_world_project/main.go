package main

import (
	"fmt"
)

func main() {
	num := 15
	// &num - pointer to num, shows the memory address of num

	num2 := num
	fmt.Println(num2, &num2, num)

	num2 = 16
	fmt.Println(num2, &num2, num)

	num3 := &num
	// * используется для получения значения по указателю
	fmt.Println(num3, &num3, *num3)

	// *num используется для изменения значения по указателю
	*num3 = 17
	fmt.Println(num, num2, *num3)
}
