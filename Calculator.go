package main

import (
	"fmt"
	"log"
)

func add1(num1, num2 int) int {
	return num1 + num2
}
func substract(num1, num2 int) int {
	return num1 - num2
}
func division(num1, num2 int) int {
	return num1 / num2
}
func multiply(num1, num2 int) int {
	return num1 * num2
}
func main() {

	log.Println("Welcome to the Calculator : ")
	log.Println("Press 1 for addition " + "\n" + "Press 2 for substraction " + "\n" + "Press 3 for multiplication " + "\n" + "Press 4 for division " + "\n")

	calculatorSwitch := true

	for calculatorSwitch != false {

		var i int
		log.Println("Choose Option from above :")
		fmt.Scanln(&i)

		log.Println("Input 2 numbers")
		var num1, num2 int
		fmt.Scanln(&num1, &num2)

		if i == 1 {
			ans := add1(num1, num2)
			log.Println(ans)
		} else if i == 2 {
			ans := substract(num1, num2)
			log.Println(ans)
		} else if i == 3 {
			ans := multiply(num1, num2)
			log.Println(ans)
		} else if i == 4 {
			ans := division(num1, num2)
			log.Println(ans)
		} else {
			log.Println("Value Entered is Wrong")
		}

		log.Println("\n")

		log.Println("Do you want to continue press 1 or for exit press 0 ?")
		var ask int
		fmt.Scanln(&ask)

		if ask == 1 {
			calculatorSwitch = true
		} else {
			calculatorSwitch = false
		}
	}
}
