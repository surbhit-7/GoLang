package main

import (
	"fmt"
)

func addition(x, y int) int {
	a := x + y
	return a
}
func substraction(x, y int) int {
	a := x - y
	return a
}
func division(x, y int) int {
	a := x / y
	return a
}
func multiplication(x, y int) int {
	a := x * y
	return a
}
func main() {
	fmt.Println("Welcome to the Calculator : ")
	fmt.Println("Press 1 for addition " + "\n" + "Press 2 for substraction " + "\n" + "Press 3 for multiplication " + "\n" + "Press 4 for division " + "\n")
	flag := true
	for flag != false {
		var i int
		fmt.Println("Choose Option from above :")
		fmt.Scanln(&i)
		fmt.Println("Input 2 numbers")
		var a, b int
		fmt.Scanln(&a, &b)
		if i == 1 {
			c := addition(a, b)
			fmt.Println(c)
		} else if i == 2 {
			c := substraction(a, b)
			fmt.Println(c)
		} else if i == 3 {
			c := multiplication(a, b)
			fmt.Println(c)
		} else if i == 4 {
			c := division(a, b)
			fmt.Println(c)
		} else {
			fmt.Println("Value Entered is Wrong")
		}
		fmt.Println("\n")
		fmt.Println("Do you want to continue press 1 or for exit press 0 ?")
		var ask int
		fmt.Scanln(&ask)
		if ask == 1 {
			flag = true
		} else {
			flag = false
		}
	}
}
