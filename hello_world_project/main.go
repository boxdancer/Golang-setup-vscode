package main

import "fmt"

func main() {
	add5 := adder(10)
	fmt.Println(add5(5))
	fmt.Println(add5(10))
}

func adder(n int) func(int) int {
	sum := n
	return func(x int) int {
		sum += x
		return sum
	}
}
