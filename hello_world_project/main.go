package main

import (
	"fmt"
	"math"
)

func main() {
	num := 2
	printNumberInfo(num)

}

func printNumberInfo(num int) {
	fmt.Println(checkSign(num))
	fmt.Println(checkParity(num))
	if num > 0 {
		fmt.Println(checkSqrt(num))
	}
}

func checkSign(num int) string {
	if num > 0 {
		return fmt.Sprintf("Число %d положительное.", num)
	} else if num < 0 {
		return fmt.Sprintf("Число %d отрицательное.", num)
	} else {
		return "Число равно 0."
	}
}

func checkParity(num int) string {
	if num%2 == 0 {
		return fmt.Sprintf("Число %d четное.", num)
	} else {
		return fmt.Sprintf("Число %d нечетное.", num)
	}
}

func checkSqrt(num int) string {
	sqrt := math.Sqrt(float64(num))
	if sqrt == float64(int(sqrt)) {
		return fmt.Sprintf("Квадратный корень числа %v является целым числом и равен %v.", num, sqrt)
	} else {
		return fmt.Sprintf("Квадратный корень числа %v не является целым числом и равен %.5f.", num, sqrt)
	}
}
