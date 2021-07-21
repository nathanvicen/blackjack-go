package blackjack

import (
	"blackjack-go/deck"
	"fmt"
	"strings"
)

type Result struct {
	Naturals bool
	PlayerWin bool
	DealerWin bool
	Push bool
	PlayerBust bool
	DealerBust bool
}

type Hand struct {
	Cards []deck.Card
	Value int
	ValueAlt int
	DisplayValue int
	DisplayValueAlt int
	Result Result
	Ace bool
}

func (h *Hand) CalculateValue() {
	totalValue := 0
	totalValueAlt := 0
	displayValue := 0
	displayValueAlt := 0
	containsAce := false
	for _, card := range h.Cards {
		numericalValue := card.GetCardNumericalValue()
		totalValue = totalValue + numericalValue
		if strings.EqualFold(card.Value, "ace") {
			containsAce = true
		}

		if card.Show == true {
			displayValue = displayValue + numericalValue
		}

		h.Ace = containsAce
	}
	if containsAce == true {
		totalValueAlt = (totalValue) + 10
	}

	if totalValueAlt > 21 && totalValue <= 21 {
		totalValueAlt = 0
		displayValueAlt = 0
	}

	if totalValueAlt == 21 {
		totalValue = 21
		totalValueAlt = 0
		displayValueAlt = 0
	}

	h.Value = totalValue
	h.ValueAlt = totalValueAlt
	h.DisplayValue = displayValue
	h.DisplayValueAlt = displayValueAlt
}

func InitDeal(d deck.Deck) (Hand, Hand, deck.Deck) {
	var playerHand Hand
	var dealerHand Hand

	playerCard1, d := deck.DealOne(d)
	dealerCard1, d := deck.DealOne(d)
	playerCard2, d := deck.DealOne(d)
	dealerCard2, d := deck.DealOne(d)
	dealerCard2.Show = false

	playerHand.Cards = append(playerHand.Cards, playerCard1)
	playerHand.Cards = append(playerHand.Cards, playerCard2)

	dealerHand.Cards = append(dealerHand.Cards, dealerCard1)
	dealerHand.Cards = append(dealerHand.Cards, dealerCard2)

	return playerHand, dealerHand, d
}

func (h *Hand) ShowAllDealerCards() {
	fmt.Println("Showing all cards")
	for _, dealerCard := range h.Cards {
		dealerCard.Show = true
		fmt.Println(dealerCard)
	}
}

func (h *Hand) Hit(d deck.Deck) deck.Deck {
	card, d := deck.DealOne(d)

	h.Cards = append(h.Cards, card)

	h.CalculateValue()

	return d
}

func (h *Hand) DealerHit(d deck.Deck) deck.Deck {
	dealerCard, d := deck.DealOne(d)

	h.Cards = append(h.Cards, dealerCard)

	h.CalculateValue()

	return d
}

func (h *Hand) CalculateResult()  {

	if h.Value > 21 {
		h.Result.PlayerBust = true
		h.Result.PlayerWin = false
	}
}

func (h *Hand) CheckNaturals() Result {
	result := Result{
		Naturals: false,
		PlayerWin: false,
		DealerWin: false,
		Push: false,
	}

	if h.Ace == true {
		if h.Value == 21 || h.ValueAlt == 21 {
			result.Naturals = true
		}
	}

	return result
}

func (h *Hand) PrintShowAll() {
	for _, card := range h.Cards {
		fmt.Println(card.Name)
	}
	fmt.Printf("Hand Total: %d", h.Value)
	fmt.Println("")

	if h.Ace == true {
		fmt.Printf("Alternate Hand Total: %d", h.Value + 10)
		fmt.Println("")
	}
}

func (h *Hand) Print() {
	for _, card := range h.Cards {
		if card.Show == true {
			fmt.Println(card.Name)
		} else {
			fmt.Println("???")
		}
	}
	fmt.Printf("Hand Total: %d", h.DisplayValue)
	fmt.Println("")

	if h.Ace == true {
		fmt.Printf("Alternate Hand Total: %d", h.DisplayValue + 10)
		fmt.Println("")
	}
}

func (h *Hand) PrintDealer() {
	for _, card := range h.Cards {
		if card.Show == true {
			fmt.Println(card.Name)
		} else {
			fmt.Println("???")
		}
	}
	fmt.Printf("Hand Total: %d", h.DisplayValue)
	fmt.Println("")
}


