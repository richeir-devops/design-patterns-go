package decorator

///////////////////////////////////////////
// 装饰模式（Decorator）:
// 就像在调制一杯咖啡一样，可以往里面加东西，账单、描述都会自动地做出相应的变更
///////////////////////////////////////////

///////////////////////////////////////////

type Coffee interface {
	GetCost() float32
	GetDescription() string
}

///////////////////////////////////////////

///////////////////////////////////////////

type SimpleCoffee struct {
}

func (sc *SimpleCoffee) GetCost() float32 {
	return 10
}

func (sc *SimpleCoffee) GetDescription() string {
	return "Simple coffee"
}

///////////////////////////////////////////

///////////////////////////////////////////

type MilkCoffee struct {
	coffee Coffee
}

func (mc *MilkCoffee) Construct(c Coffee) {
	mc.coffee = c
}

func (mc MilkCoffee) GetCost() float32 {
	return mc.coffee.GetCost() + 2
}

func (mc MilkCoffee) GetDescription() string {
	return mc.coffee.GetDescription() + ", milk"
}

///////////////////////////////////////////

///////////////////////////////////////////

type WhipCoffee struct {
	coffee Coffee
}

func (wc *WhipCoffee) Construct(c Coffee) {
	wc.coffee = c
}

func (wc WhipCoffee) GetCost() float32 {
	return wc.coffee.GetCost() + 5
}

func (wc WhipCoffee) GetDescription() string {
	return wc.coffee.GetDescription() + ", whip"
}

///////////////////////////////////////////
