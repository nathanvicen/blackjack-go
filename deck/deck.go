package deck

import (
	"fmt"
	"math/rand"
	"time"
)

type Deck[]Card

func (d Deck) Print() {
	for _, card := range d {
		fmt.Printf("%+v", card)
	}
}

func MultipleNewDecks(numDecks int) Deck {
	cards := Deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King" }

	for i := 0; i < numDecks; i++ {
		for _, suit := range cardSuits {
			for _, value := range cardValues {
				cards = append(cards, NewCard(suit, value, AsciiVersionOfCard(suit, value)))
			}
		}
	}
	return cards
}

func NewDeck() Deck {
	cards := Deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, NewCard(suit, value, ""))
		}
	}
	return cards
}

func DealOne(d Deck) (Card, Deck) {
	return d[0], d[1:]
}

func Deal(d Deck, handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}

func (d Deck) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		// Swap Elements
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
