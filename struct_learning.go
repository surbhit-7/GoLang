package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	boy := person{lastName: "Vishwakarma", firstName: "Surbhit"}
	fmt.Println(boy)
}
