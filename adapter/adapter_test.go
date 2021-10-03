package adapter

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	lion := &AfricanLioe{}

	hunter := &Hunter{}
	// hunt a lion
	hunter.hunt(lion)

	fmt.Println("Lion dead.")

	wildDog := &WildDog{}
	wildDogAdapter := NewWildDogAdapter(wildDog)
	// hunt a wilddog
	hunter.hunt(wildDogAdapter)

	fmt.Println("Wilddog dead.")
}
