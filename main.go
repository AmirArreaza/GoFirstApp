package main

func main() {

	/* Type will be infered from the right hand side of the declaration (this is a string)
	var card = "Ace of Spades"
	Go compiler will read the := char and assign the right value. It cannot change like in Javascript!!!
	:= is used only for NEW values
	*/

	cards := newDeck()

	cards.shuffle()

	cards.saveToFile("Cartas.txt")

	cards.print()
}
