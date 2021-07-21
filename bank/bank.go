package bank

import "fmt"

type Bank struct {
	Balance float32
}

func (b *Bank) Deposit(depositAmount float32) {
	b.Balance += depositAmount
}

func (b *Bank) Win(winAmount float32) {
	b.Balance += winAmount
}

func (b *Bank) Lose(loseAmount float32) {
	b.Balance -= loseAmount
}

func (b *Bank) Print() {
	fmt.Printf("Balance: $%.2f\n", b.Balance)
}

