package main

import (
	"fmt"
	"strconv"
)

func main() {
	priceStr := "1234.5678"
	quantityStr := "10"

	priceFloat, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		fmt.Println("Error parsing price:", err)
		return
	}

	quantityInt, _ := strconv.Atoi(quantityStr)

	result := priceFloat * float64(quantityInt)
	fmt.Println(strconv.FormatFloat(result, 'f', 2, 64))
}
