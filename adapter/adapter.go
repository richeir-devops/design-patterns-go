package adapter

import "fmt"

///////////////////////////////////////////

type Lion interface {
	Roar()
}

type AfricanLioe struct {
}

func (lone *AfricanLioe) Roar() {
	fmt.Println("Lion roar!")
}

///////////////////////////////////////////

///////////////////////////////////////////
type WildDog struct {
}

func (dog *WildDog) Bark() {
	fmt.Println("Wilddog bark!")
}

///////////////////////////////////////////

///////////////////////////////////////////

type Hunter struct {
}

func (hunter *Hunter) hunt(lion Lion) {
	lion.Roar()
}

///////////////////////////////////////////

///////////////////////////////////////////

type WildDogAdapter struct {
	wildDog *WildDog
}

func (dog *WildDogAdapter) Roar() {
	dog.wildDog.Bark()
}

func NewWildDogAdapter(dog *WildDog) *WildDogAdapter {
	return &WildDogAdapter{wildDog: dog}
}

///////////////////////////////////////////
