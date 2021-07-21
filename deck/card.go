package deck

import "strings"

type Card struct {
	Suit  string
	Value string
	Name  string
	Show  bool
}

func NewCard(suit string, value string) Card {
	return Card{
		Suit:  suit,
		Value: value,
		Name:  value + " of " + suit,
		Show: true,
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


