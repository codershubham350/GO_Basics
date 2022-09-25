package main

func main() {
	// var card string = "Ace of Spades"
	//	card := newCard() // equivalent to above commented code
	//cards := deck{"Ace of Diamonds", newCard()}
	//cards = append(cards, "Six of Spades") // append does not modify the original slice it returns new slice
	// temp := make([]string, 10, 50)
	// cards := newDeck()
	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// cards := newDeck()
	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()

	//cards.print()

}
