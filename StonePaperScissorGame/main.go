package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	log.Println("Welcome to Stone Paper Scissors Game :")
	log.Println("Choose 1 for Stone\nChoose 2 for Paper\nChoose 3 for Scissor")

	gameSwitch := true

	var choosingNum, switchN int

	var playerHand, computerHand hand

	for gameSwitch != false {

		fmt.Scanln(&choosingNum)
		playerHand = chooseHand(choosingNum)

		if playerHand == "Err" {
			os.Exit(1)
		}

		computerHand = choosingComputersHand()
		log.Println("Computer Chooses :", computerHand)

		PlayerHandVal := playerHand.Result()
		computerHandVal := computerHand.Result()

		result := Decider(PlayerHandVal, computerHandVal)
		fmt.Println("\nAnd The Winner of the Game is : ", result)
		fmt.Println("Do you wish to Play again press 1 :")
		fmt.Scanln(&switchN)

		if switchN != 1 {
			gameSwitch = false
		}

	}
}
