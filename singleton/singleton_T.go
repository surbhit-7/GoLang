package main

import "log"

func main() {

	instance := GetInstance()

	t := instance.AddOne()

	log.Println(t)

}
