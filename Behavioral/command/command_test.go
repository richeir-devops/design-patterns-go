package command

import "testing"

func Test01(t *testing.T) {
	bulb := &Bulb{}

	turnOn := &TurnOn{Bulb: *bulb}
	turnOff := &TurnOff{Bulb: *bulb}

	remote := &RemoteControl{}
	remote.submit(turnOn)
	remote.submit(turnOff)
}
