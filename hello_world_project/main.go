package main

import "fmt"

type Day int

const (
	_ Day = iota // пропускаем 0
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func main() {
	fmt.Println(isWeekend(Monday))
	fmt.Println(isWeekend(Tuesday))
	fmt.Println(isWeekend(Saturday))
	fmt.Println(isWeekend(Sunday))
	fmt.Println(isWeekend(10))
}

func isWeekend(day Day) bool {
	switch day {

	case Saturday, Sunday:
		return true

	default:
		return false
	}
}
