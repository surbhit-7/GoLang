package main

import "log"

func main() {
	// Type 1
	/*var anyNum int
	fmt.Scanln(&anyNum)

	switch i := anyNum; i % 2 {
	case 1, 7, 3, 9:
		log.Println("Number is odd")

	case 0, 2, 4, 8:
		log.Println("Number is even")

	default:
		log.Println("Re enter the number")
	}*/

	//Type 2
	temp := -10

	switch {
	case temp > 0:
		log.Println("Normal")
	case temp < 0:
		log.Println("Cold")
	}
}
