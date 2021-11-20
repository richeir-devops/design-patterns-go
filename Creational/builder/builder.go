package builder

type Burger struct {
	size      int
	cheese    bool
	pepperoni bool
	lettuce   bool
	tomato    bool
}

func (b *Burger) contrust(bb *BurgerBuilder) *Burger {
	b.size = bb.size
	b.cheese = bb.cheese
	b.pepperoni = bb.pepperoni
	b.lettuce = bb.lettuce
	b.tomato = bb.tomato

	return b
}

type BurgerBuilder struct {
	Burger
}

func (bb *BurgerBuilder) addPepperoni() {
	bb.pepperoni = true
}

func (bb *BurgerBuilder) addLettuce() {
	bb.lettuce = true
}

func (bb *BurgerBuilder) addCheese() {
	bb.cheese = true
}

func (bb *BurgerBuilder) addTomato() {
	bb.tomato = true
}

func (bb *BurgerBuilder) build() *Burger {
	return bb.contrust(bb)
}
