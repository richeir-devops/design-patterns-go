package flyweight

import "testing"

func Test01(t *testing.T) {
	teaMaker := &TeaMaker{AvailableTea: make([]*KarakTea, 10)}

	shop := &TeaShop{orders: make([]*KarakTea, 10), teaMaker: teaMaker}

	shop.TakeOrder("less sugar", 1)
	shop.TakeOrder("more milk", 2)
	shop.TakeOrder("without sugar", 5)

	shop.Serve()
}
