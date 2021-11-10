package main
//import "fmt"

func main(){
	cards := newDeckFromFile(".//data//hand.txt")
	cards.print()
	cards.shuffle()
	cards.print()
	//hand, _ := deal(cards, 5)
	//hand.saveToFile(".//data//hand.txt")
}