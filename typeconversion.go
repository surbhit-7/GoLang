package main

import (
	"fmt"
	"strings"
)

func main() {
	// String to Byte
	/*cards := []byte("HELLO WORLD")
	fmt.Println(cards)*/

	//Float to INT
	/*num := 4.5
	num1 := int(num)
	fmt.Println(num1)*/

	/*num := []string{"HELLO", "HII"}
	num1 := strings.Join(num, " ")
	fmt.Println(num1)*/

	string1 := "HELLO HII"
	list := strings.Split(string1, "")
	fmt.Println(list)

}
