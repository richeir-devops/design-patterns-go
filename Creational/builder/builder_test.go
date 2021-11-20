package builder

import "testing"

func Test01(t *testing.T) {
	burderBuilder := &BurgerBuilder{}
	burderBuilder.size = 14
	burderBuilder.addTomato()
	burger := burderBuilder.build()

	t.Logf("Burger's tomato: %v", burger.tomato)
}
