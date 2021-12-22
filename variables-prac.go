package main

import (
	"fmt"
	"log"
)

var car int

func main() {
	// Declerations
	//Type 1 Dec
	var hello = "hello"
	//type 2 Dec

	var hello1 string = "Hello"
	// type 3 dec
	hello2 := "HELLO"

	log.Println(hello+" ", hello2, hello1)

	// scanf
	// type 1
	var d, c (int)
	fmt.Scanln(&d, &c)

	log.Println(d + c)

	//type 2
	var b (int)
	fmt.Scanf("%d", &b)

	fmt.Println(b)
}
