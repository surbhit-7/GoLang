package main

import (
	"fmt"
)

func addition(num1, num2 int) int {
	temp := num1 + num2
	return temp
}
func substraction(num1, num2 int) int {
	temp := num1 - num2
	return temp
}
func division(num1, num2 int) int {
	temp := num1 / num2
	return temp
}
func multiplication(num1, num2 int) int {
	temp := num1 * num2
	return temp
}
func main() {
	fmt.Println("Welcome to the Calculator : ")
	fmt.Println("Press 1 for addition " + "\n" + "Press 2 for substraction " + "\n" + "Press 3 for multiplication " + "\n" + "Press 4 for division " + "\n")
	calculatorSwitch := true
	for calculatorSwitch != false {
		var i int
		fmt.Println("Choose Option from above :")
		fmt.Scanln(&i)
		fmt.Println("Input 2 numbers")
		var num1, num2 int
		fmt.Scanln(&num1, &num2)
		if i == 1 {
			ans := addition(num1, num2)
			fmt.Println(ans)
		} else if i == 2 {
			ans := substraction(num1, num2)
			fmt.Println(ans)
		} else if i == 3 {
			ans := multiplication(num1, num2)
			fmt.Println(ans)
		} else if i == 4 {
			ans := division(num1, num2)
			fmt.Println(ans)
		} else {
			fmt.Println("Value Entered is Wrong")
		}
		fmt.Println("\n")
		fmt.Println("Do you want to continue press 1 or for exit press 0 ?")
		var ask int
		fmt.Scanln(&ask)
		if ask == 1 {
			calculatorSwitch = true
		} else {
			calculatorSwitch = false
		}
	}
}
