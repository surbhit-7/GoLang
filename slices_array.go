package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//Practising slice :-
	/*cards := []string{}
	for i := 0; i < 3; i++ {
		var a string
		fmt.Scanln(&a)
		cards = append(cards, a)
	}*/
	/*for i, x := range cards {
		fmt.Println(i, x)
	}*/
	//fmt.Println(cards)
	//fmt.Println(cards[1])
	//ARRAYS
	/*my_array := [10]int{}
	for i := 0; i < 10; i++ {
		my_array[i] = i + 1
	}*/
	//fmt.Println(my_array)

	// Predefining the size of slice:-
	cards := make([]string, 5)
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < len(cards); i++ {
		scanner.Scan()
		line := scanner.Text()
		cards[i] = line
	}
	fmt.Println(cards)
}
