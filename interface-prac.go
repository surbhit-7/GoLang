package main

import "log"

type dog struct {
	name string
}

type cat struct {
	name string
}

func GetCredentials(a animal) {
	log.Println(a.GetName())
}

type animal interface {
	GetName() string
}

func (d dog) GetName() string {
	//d.name = name
	return "WOOF WOOF my name is " + d.name
}
func (c cat) GetName() string {
	//c.name = name
	return "MEOW MEOW my name is " + c.name
}

func main() {
	dog1 := dog{name: "Raghav"}
	cat1 := cat{name: "Prattek"}
	var inter animal
	inter = dog1
	//var inter animal
	GetCredentials(inter)
	GetCredentials(cat1)

}
