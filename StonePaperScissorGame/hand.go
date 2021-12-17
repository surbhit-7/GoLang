package main

import (
	"math/rand"
	"time"
)

type hand string

func chooseHand(i int) hand {
	if i == 1 {
		return "Stone"
	} else if i == 2 {
		return "Paper"
	} else if i == 3 {
		return "Scissor"
	} else {
		return "Err"
	}
}
func choosingComputersHand() hand {
	rand.Seed(time.Now().UnixNano())
	RandNum := rand.Intn(100)
	if (RandNum%3 + 1) == 1 {
		return "Stone"
	} else if (RandNum%3 + 1) == 2 {
		return "Paper"
	} else {
		return "Scissor"
	}
}
func (h hand) Result() int {
	if h == "Stone" {
		return 1
	} else if h == "Paper" {
		return 2
	} else {
		return 3
	}
}
func Decider(i, j int) hand {
	if (i == 1 && j == 2) || (i == 2 && j == 3) || (i == 3 && j == 1) {
		return "Computer"
	} else if (j == 1 && i == 2) || (j == 2 && i == 3) || (j == 3 && i == 1) {
		return "Player"
	} else {
		return "Draw"
	}
}
