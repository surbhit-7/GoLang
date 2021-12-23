package main

import "log"

type human struct {
	name string
	age  string
}

type persons struct {
	human
}

func main() {
	s := persons{human{name: "Surbhit", age: "23"}}
	log.Println(s)
}
