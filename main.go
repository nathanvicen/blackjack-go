package main

import (
	"blackjack-go/bank"
	"blackjack-go/blackjack"
	"blackjack-go/deck"
	"fmt"
	"github.com/inancgumus/screen"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var playerBank bank.Bank
var cardDeck deck.Deck
var wager float32

func printHeader() {
	fmt.Println("----------------------- BLACKJACK -----------------------")
}

func printFooter() {
	fmt.Println("----------------------------------------------------------")
}

func resetScreen() {
	screen.Clear()
	screen.MoveTopLeft()
	printHeader()
	playerBank.Print()

	if wager != 0 {
		fmt.Printf("Wager: $%.2f\n", wager)
	}
}


func deckSizePrompt() int {
	printHeader()
	var deckSizeResponse string
	for {
		//resetScreen()
		fmt.Println("Enter Number of Decks You'd Like to Play With")
		fmt.Scan(&deckSizeResponse)

		validDeckSizeRegex := regexp.MustCompile("\\d+")
		validDeckSize := validDeckSizeRegex.MatchString(deckSizeResponse)
		deckSize, _ := strconv.ParseInt(deckSizeResponse, 10, 32)

		if validDeckSize == true && deckSize >= 1 {
			break
		}

		fmt.Println("Oops... Please enter deck size as a whole number...")
	}

	deckSize, _ := strconv.ParseInt(deckSizeResponse, 10, 32)
	return int(deckSize)
}



func nextMove(c chan string) {
	var nextMoveResponse string
	for {
		fmt.Println("Stand / Hit")
		fmt.Scan(&nextMoveResponse)

		standRegex := regexp.MustCompile("(?i)Stand")
		stand := standRegex.MatchString(nextMoveResponse)
		hitRegex := regexp.MustCompile("(?i)Hit")
		hit := hitRegex.MatchString(nextMoveResponse)

		if stand == true {
			c <- "stand"
			break
		}

		if hit == true {
			c <- "hit"
			break
		}
		fmt.Println("Oops... Please choose to Stand or Hit...")
	}
}

func printHands(playerHand blackjack.Hand, dealerHand blackjack.Hand) {
	fmt.Println("----------------------- YOUR HAND -----------------------")
	fmt.Println("")
	playerHand.Print()
	fmt.Println("")
	printFooter()

	fmt.Println("")

	fmt.Println("----------------------- DEALER HAND ----------------------")
	fmt.Println("")
	dealerHand.PrintDealer()
	fmt.Println("")
	printFooter()

}

func printHandsShowAll(playerHand blackjack.Hand, dealerHand blackjack.Hand) {
	fmt.Println("----------------------- YOUR HAND -----------------------")
	fmt.Println("")
	playerHand.Print()
	fmt.Println("")
	printFooter()

	fmt.Println("")

	fmt.Println("----------------------- DEALER HAND ----------------------")
	fmt.Println("")
	dealerHand.PrintShowAll()
	fmt.Println("")
	printFooter()

}

func continuePrompt(c chan bool) {
	var continueResponse string
	for {
		fmt.Println("Play Again? Yes / No")
		fmt.Scan(&continueResponse)

		yesInputRegex := regexp.MustCompile("(?i)Yes")
		yesInput := yesInputRegex.MatchString(continueResponse)

		noInputRegex := regexp.MustCompile("(?i)No")
		noInput := noInputRegex.MatchString(continueResponse)

		if yesInput == true {
			c <- true
			break
		}

		if noInput == true {
			c <- false
			break
		}

	}
}

func wagerPrompt() {
	var betResponse string
	for {
		fmt.Println("Enter Bet Amount")
		fmt.Print("$")
		fmt.Scan(&betResponse)

		validBetAmountRegex := regexp.MustCompile("\\d+")
		validBetAmount := validBetAmountRegex.MatchString(betResponse)

		betAmount, _ := strconv.ParseFloat(betResponse, 10)

		if validBetAmount == true {
			if float32(betAmount) > playerBank.Balance {
				fmt.Printf("You don't have enough money! Your balance is: $%.2f\n", playerBank.Balance)
			} else {
				break
			}
		} else {
			fmt.Println("Oops... Please place a bet as a whole dollar amount...")
		}
	}
	betAmount, _ := strconv.ParseFloat(betResponse, 10)
	wager = float32(betAmount)
}

func playHand(c chan blackjack.Result)  {

	handComplete := false

	result := blackjack.Result{
		PlayerWin: false,
		DealerWin: false,
		Naturals: false,
		Push: false,
		PlayerBust: false,
		DealerBust: false,
	}

	playerHand, dealerHand, newDeck := blackjack.InitDeal(cardDeck)
	cardDeck = newDeck

	playerHand.CalculateValue()
	playerHand.Result = playerHand.CheckNaturals()

	if playerHand.Result.Naturals == true {
		result.Naturals = true
		result.PlayerWin = true
		handComplete = true
	}

	dealerHand.CalculateValue()
	dealerHand.Result = dealerHand.CheckNaturals()
	if dealerHand.Result.Naturals == true {
		result.Naturals = true
		result.DealerWin = true
		handComplete = true
	}

	if handComplete == true {
		printHandsShowAll(playerHand, dealerHand)
		c <- result
		return
	}

	printHands(playerHand, dealerHand)

	nextC := make(chan string)

	for {

		go nextMove(nextC)

		nextMove := <-nextC

		// Hit
		if strings.EqualFold(nextMove, "hit") {
			d := playerHand.Hit(cardDeck)
			cardDeck = d
			resetScreen()
			printHands(playerHand, dealerHand)
			if playerHand.Value > 21 {
				result.PlayerBust = true
				result.DealerWin = true
				nextMove = "stand"
			}
		}

		// Stand
		if strings.EqualFold(nextMove, "stand") {
			if result.PlayerBust == false && playerHand.Ace == true && playerHand.ValueAlt > playerHand.Value {
				playerHand.Value = playerHand.ValueAlt
			}

			resetScreen()
			printHandsShowAll(playerHand, dealerHand)

			if dealerHand.Value >= 17 {
				break
			}
			for {
				d := dealerHand.Hit(cardDeck)
				cardDeck = d
				resetScreen()
				printHandsShowAll(playerHand, dealerHand)
				if dealerHand.Value > 21 {
					result.DealerBust = true
					break
				}
				if dealerHand.Value >= 17 {
					break
				}
				time.Sleep(1 * time.Second)
			}
			break
		}
	}

	if playerHand.Value == dealerHand.Value {
		result.Push = true
	} else if playerHand.Value > dealerHand.Value && result.PlayerBust == false {
		result.PlayerWin = true
	} else if dealerHand.Value > playerHand.Value {
		result.DealerWin = true
	}

	c <- result
}

func main() {
	screen.Clear()
	screen.MoveTopLeft()

	playerBank = bank.Bank{Balance: 100.00}
	deckSize := deckSizePrompt()

	cardDeck = deck.MultipleNewDecks(deckSize)
	cardDeck.Shuffle()

	playC := make(chan blackjack.Result)
	continueC := make(chan bool)

	for {

		if len(cardDeck) <= 19 {
			fmt.Println("Shuffling Cards...")
			cardDeck = deck.MultipleNewDecks(deckSize)
			cardDeck.Shuffle()
		}

		resetScreen()

		wagerPrompt()

		fmt.Printf("Bet Amount: $%.2f\n", wager)

		playerBank.Lose(wager)

		resetScreen()

		go playHand(playC)

		handResult := <- playC

		fmt.Println("")

		if handResult.Push == true {
			fmt.Println("Push")
			playerBank.Win(wager)
		}

		if handResult.PlayerWin == true  && handResult.Naturals == true {
			fmt.Println("Blackjack!")
			fmt.Printf("You've won $%.2f\n", wager * 2)
			playerBank.Win(wager * 2)
		} else if handResult.PlayerWin == true  && handResult.Naturals == false {
			fmt.Println("You won!")
			fmt.Printf("You've won $%.2f\n", wager * 1.5)
			playerBank.Win(wager * 1.5)
		} else if handResult.PlayerBust == false && handResult.DealerBust == true {
			fmt.Println("You won!")
			fmt.Printf("You've won $%.2f\n", wager * 1.5)
			playerBank.Win(wager * 1.5)
		} else if handResult.DealerWin == true {
			if handResult.PlayerBust == true {
				fmt.Println("BUST!")
			}
			fmt.Println("You lost")
		} else if playerBank.Balance == 0 {
			fmt.Println("You're all out of money! Bye!")
			break
		}

		fmt.Println("")

		go continuePrompt(continueC)

		continueResponse := <- continueC

		if continueResponse == false {
			break
		}

		wager = 0

	}

}
