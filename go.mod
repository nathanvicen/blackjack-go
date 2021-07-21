module blackjack-go

go 1.16

replace blackjack-go/deck => ./deck

replace blackjack-go/blackjack => ./blackjack

replace blackjack-go/bank => ./bank

require (
	blackjack-go/bank v0.0.0-00010101000000-000000000000 // indirect
	blackjack-go/blackjack v0.0.0-00010101000000-000000000000 // indirect
	blackjack-go/deck v0.0.0-00010101000000-000000000000 // indirect
	github.com/buger/goterm v1.0.1 // indirect
	github.com/inancgumus/screen v0.0.0-20190314163918-06e984b86ed3 // indirect
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97 // indirect
)
