package deck

import "strings"

type Card struct {
	Suit  string
	Value string
	Name  string
	Show  bool
	Symbol string
}

func NewCard(suit string, value string, symbol string) Card {
	return Card{
		Suit:  suit,
		Value: value,
		Name:  value + " of " + suit,
		Show: true,
		Symbol: symbol,
	}
}



func (c Card) GetCardNumericalValue() int {
	numericalValues := map[string]int {
		"ace": 1,
		"two": 2,
		"three": 3,
		"four": 4,
		"five": 5,
		"six": 6,
		"seven": 7,
		"eight": 8,
		"nine": 9,
		"ten": 10,
		"jack": 10,
		"queen": 10,
		"king": 10,
	}
	return numericalValues[strings.ToLower(c.Value)]
}

func AsciiHiddenCard() string {
	var lines[]string
	lines = append(lines, "\n")
	lines = append(lines, "┌─────────┐")
	lines = append(lines,"│░░░░░░░░░│")
	lines = append(lines, "│░░░░░░░░░│")
	lines = append(lines, "│░░░░░░░░░│")
	lines = append(lines, "│░░░░░░░░░│")
	lines = append(lines, "│░░░░░░░░░│")
	lines = append(lines, "│░░░░░░░░░│")
	lines = append(lines, "│░░░░░░░░░│")
	lines = append(lines, "└─────────┘")

	return strings.Join(lines, "\n")
}

func AsciiVersionOfCard(suit string, value string) string {
	var suitSymbol string
	var suitValue string

	var lines[]string

	switch suit {
	case "Spades":
		suitSymbol = "♠"
	case "Diamonds":
		suitSymbol = "♦"
	case "Hearts":
		suitSymbol = "♥"
	case "Clubs":
		suitSymbol = "♣"
	}

	switch value {
	case "Ace":
		suitValue = " A "
	case "Two":
		suitValue = " 2 "
	case "Three":
		suitValue = " 3 "
	case "Four":
		suitValue = " 4 "
	case "Five":
		suitValue = " 5 "
	case "Six":
		suitValue = " 6 "
	case "Seven":
		suitValue = " 7 "
	case "Eight":
		suitValue = " 8 "
	case "Nine":
		suitValue = " 9 "
	case "Ten":
		suitValue = "10 "
	case "Jack":
		suitValue = " J "
	case "Queen":
		suitValue = " Q "
	case "King":
		suitValue = " K "
	}

	lines = append(lines, "\n")
	lines = append(lines, "┌─────────┐")
	lines = append(lines, strings.Replace(strings.Replace("│{}{}     │", "{}", suitValue, 1), "{}", suitSymbol, 1))
	lines = append(lines, "│         │")
	lines = append(lines, "│         │")
	lines = append(lines, strings.Replace("│    {}    │", "{}", suitSymbol, 1))
	lines = append(lines, "│         │")
	lines = append(lines, "│         │")
	lines = append(lines, strings.Replace(strings.Replace("│     {}{}│", "{}", suitValue, 1), "{}", suitSymbol, 1))
	lines = append(lines, "└─────────┘")

	return strings.Join(lines, "\n")
}


