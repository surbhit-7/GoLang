package main

import (
	"log"
	"os"
)

func main() {
	//log.Println("Hello World")
	// creating our own Log
	logFile, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		log.Println(err)
	}

	defer logFile.Close()

	newStdLogger := log.New(logFile, "New Logger : ", log.LstdFlags|log.Llongfile)
	newStdLogger.Println("HEllo World")
}
