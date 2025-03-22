package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	price := 10.2

	price = math.Round(price*1000) / 1000
	myFloatToStr := strconv.FormatFloat(price, 'f', 3, 64)

	fmt.Println(myFloatToStr)
}
