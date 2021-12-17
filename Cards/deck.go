package main

import (
	"io/ioutil"
	"log"
	"strings"
)

type deck []string

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
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

func (d deck) toString() string {
	var result = strings.Join([]string(d), " , ")
	return result
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
