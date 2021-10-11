package chainresponsibilities

import "fmt"

type Account struct {
	Successor *Account
	Balance   float32
	Name      string
}

func (ac *Account) SetNext(account *Account) {
	ac.Successor = account
}

func (ac *Account) Pay(amountToPay float32) {
	if ac.canPay(amountToPay) {
		fmt.Printf("Paid %f using %s \n", amountToPay, ac.Name)
	} else if ac.Successor != nil {
		fmt.Printf("Cannot pay using %s. Processing ..\n", ac.Name)
		ac.Successor.Pay(amountToPay)
	}
}

func (ac *Account) canPay(amountToPay float32) bool {
	return ac.Balance >= amountToPay
}

type Bank struct {
	Acc *Account
}

type Paypal struct {
	Acc *Account
}

type Bitcoin struct {
	Acc *Account
}
