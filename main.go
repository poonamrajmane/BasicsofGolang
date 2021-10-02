package main

import "fmt"

//var de string = "poo"

//p :="9"

func main() {

	//printState()
	//var card string = "Ace of Spades"
	card := newCard()
	de := newCard1()
	//	var de string
	//de = "poo"

	cards := []string{"Rohit", "Poonam", newCard()}
	cards = append(cards, "pillu")
	fmt.Println(cards)
	fmt.Println(card)
	fmt.Println(de)

	for i, car := range cards {
		fmt.Println(i, car)
	}

}

func newCard() string {
	return "Five of Diamonds"
}

func newCard1() int {
	return 6
}
