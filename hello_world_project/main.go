package main

import (
	"fmt"
)

func main() {
	var name, surname, age string
	fmt.Scan(&name, &surname)
	fmt.Scan(&age)
	fmt.Println("Name: ", name, "\nSurname: ", surname, "\nAge: ", age)
	fmt.Println("Приятно познакомиться,", name+".", "Я 5 лет назад познакомился с человеком, у которого тоже фамилия", surname+",", "вам тогда было", age+".", "Как молоды мы были!")
}
