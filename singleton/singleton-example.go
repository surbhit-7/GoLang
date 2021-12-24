package main

import "log"

func main() {
	gi := GetInstance()

	gi.AddRecord("surbhit", "vishwakarma", 69, "CSE")

	log.Println(gi)

}
