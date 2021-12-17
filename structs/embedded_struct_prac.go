package main

import "log"

type contactInfo struct {
	phoneNumber string
	zipcode     int
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	var boy person
	boy.firstName = "Surbhit"
	boy.lastName = "Vishwakarma"
	boy.contactInfo.zipcode = 482002
	boy.contactInfo.phoneNumber = "9589228632"
	boyPointer := &boy
	boyPointer.updateName("ShaggY")
	boy.print()
}
func (p person) print() {
	log.Println(p)
}
func (updateNamePtr *person) updateName(name string) {
	updateNamePtr.firstName = name
}
