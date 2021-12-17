package main

import "log"

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, Values := range cardValues {
			cards = append(cards, suit+" of "+Values)
		}
	}
	return cards

}
func (d deck) print() {
	for i, card := range d {
		log.Println(i, card)
	}
}
