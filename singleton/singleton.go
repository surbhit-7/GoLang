package main

type singleton struct {
	val int
}
type Singleton interface {
	AddOne() int
}

var instance *singleton

func GetInstance() Singleton {

	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.val++
	return s.val
}
