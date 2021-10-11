package chainresponsibilities

import "testing"

func Test01(t *testing.T) {
	bank := &Bank{Acc: &Account{Name: "Bank", Balance: 100}}
	paypal := &Paypal{Acc: &Account{Name: "Paypal", Balance: 200}}
	bitcoin := &Bitcoin{Acc: &Account{Name: "Bitcoin", Balance: 300}}

	bank.Acc.Successor, _ = interface{}(paypal.Acc).(*Account)
	paypal.Acc.Successor, _ = interface{}(bitcoin.Acc).(*Account)

	bank.Acc.Pay(259)
}
