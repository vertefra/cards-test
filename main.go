package main

import "fmt"

func main() {
	card := newDeckFromFile("cardgs.txt")
	fmt.Println(card)
}
