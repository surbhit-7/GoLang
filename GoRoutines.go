package main

import (
	"log"
	"sync"
	"time"
)

var wg sync.WaitGroup

func say(s string) {

	for i := 1; i < 5; i++ {
		log.Println("say :", s)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done()
}

func main() {
	wg.Add(1)
	go say("hello")
	wg.Add(1)
	go say("hi")
	wg.Wait()

	//time.Sleep(time.Millisecond * 3000)
}
