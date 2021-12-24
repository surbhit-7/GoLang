package main

type database struct {
	first_name  string
	last_name   string
	roll_number int
	subject     string
}

func GetInstance() database {
	return instance
}

func GetDatabase() map[int]database {
	return records
}

//var serial_number int
var records = make(map[int]database)
var instance database

func (d *database) AddRecord(first_name string, last_name string, roll_number int, subject string) {
	d.first_name = first_name
	d.last_name = last_name
	d.roll_number = roll_number
	d.subject = subject
}
