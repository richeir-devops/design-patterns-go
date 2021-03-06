package chainresponsibilities

import "fmt"

///////////////////////////////////////////
// 责任链模式（Chain Of Responsibilities）：
// 可以理解为按照付款方式排序来付款的业务场景，如果当前付款方式的钱不够，自动跳到下一个；
// 也可以理解为自动化处理的流程，当前处理不了，自动转移到下一个既定备用流程
///////////////////////////////////////////

///////////////////////////////////////////

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

///////////////////////////////////////////

///////////////////////////////////////////

type Bank struct {
	Acc *Account
}

type Paypal struct {
	Acc *Account
}

type Bitcoin struct {
	Acc *Account
}

///////////////////////////////////////////
