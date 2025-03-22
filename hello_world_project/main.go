package main

import (
	"fmt"
	"strconv"
)

func main() {
	myInt := 10
	myFloat := 10.2134

	myIntToStr := strconv.Itoa(myInt)
	myFloatToStr := strconv.FormatFloat(myFloat, 'f', 5, 64)

	fmt.Println(myIntToStr, myFloatToStr)
}
