package main

import (
	"log"
	"sync"
	"time"
)

var wg sync.WaitGroup

func helper() {

	wg.Done()

	if r := recover(); r != nil {
		log.Println("Recovered in cleanup")
	}
}

func say(s string) {

	defer helper()

	for i := 1; i < 5; i++ {
		log.Println("say :", s)
		time.Sleep(time.Millisecond * 100)

		if i == 2 {
			panic("oh shit")
		}
	}

}

func main() {

	wg.Add(1)
	go say("hello")

	wg.Add(1)
	go say("hi")

	wg.Wait()

	//time.Sleep(time.Millisecond * 3000)
}
