package main

func main() {
	cards := newDeck()
	/*hand, remainingCards := deal(cards, 5)*/
	/*hand.print()
	remainingCards.print()*/
	/*cards.toString()
	fmt.Println(cards)*/
	cards.saveToFile("my_cards")
}
