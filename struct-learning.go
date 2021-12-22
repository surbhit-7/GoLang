package main

type person struct {
	firstName string
	lastName  string
}

func main() {
	boy := person{lastName: "Vishwakarma", firstName: "Surbhit"}
	log.Println(boy)
}
