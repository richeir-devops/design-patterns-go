package decorator

import (
	"fmt"
	"testing"
)

func Test01(t *testing.T) {
	simpleCoffee := &SimpleCoffee{}
	fmt.Println(simpleCoffee.GetCost())
	fmt.Println(simpleCoffee.GetDescription())

	milkCoffee := &MilkCoffee{}
	milkCoffee.Construct(simpleCoffee)
	fmt.Println(milkCoffee.GetCost())
	fmt.Println(milkCoffee.GetDescription())

	whipCoffee := &WhipCoffee{}
	whipCoffee.Construct(milkCoffee)
	fmt.Println(whipCoffee.GetCost())
	fmt.Println(whipCoffee.GetDescription())
}
