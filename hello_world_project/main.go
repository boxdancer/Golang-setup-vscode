package main

import (
	"fmt"
)

func main() {
	random := 5.9
	strNumber := fmt.Sprintf("%.1f", random)
	strNumberIncr := fmt.Sprintf("%.5f", random*1.1)
	isEven := int(random)%2 == 0
	predLastDigit := (int(random) / 10) % 10
	fmt.Println(
		"Исходное число:", strNumber,
		"\nИсходное число, увеличенное на 10%:", strNumberIncr,
		"\nИсходное число является четным:", isEven,
		"\nПредпоследняя цифра целой части исходного числа:", predLastDigit)
}
