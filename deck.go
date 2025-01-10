package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {

	cards := deck{}

	//Creates string slices that are
	//used to form the cards
	cardSuite := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValue := []string{"Ace", "Two", "Three", "Four", "Five", "Six",
		"Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	//Creates a card of each value
	//for every unique suite
	for _, suite := range cardSuite {
		for _, value := range cardValue {
			cards = append(cards, value+" of "+suite)
		}
	}

	return cards
}

func (d deck) print() {
	//Prints all cards in receiver deck
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {

	//Returns number of specified cards into first deck(hand)
	//Returns remaining deck
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	//Converts receiver deck into a string
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	//Converts receiver deck into a byte slice and
	//stores it into a file
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {

	//Reads from provided file and stores
	//a byte slice
	byteSlice, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	//Turns byteslice into a string
	deckString := string(byteSlice)

	//Turns string into a string slice and converts
	//into a deck type which is returned
	return deck(strings.Split(deckString, ","))
}

func (d deck) shuffle() {

	//Generates a random seed based on current local time
	//to be used for random number generation
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
