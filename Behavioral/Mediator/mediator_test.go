package mediator

import "testing"

func Test01(t *testing.T) {
	mediator := &ChatRoom{}
	john := &User{name: "John", chatMediator: mediator}
	jane := &User{name: "Jane", chatMediator: mediator}

	john.send("Hi there!")
	jane.send("Hey!")
}
