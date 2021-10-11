package facade

import "testing"

func Test01(t *testing.T) {
	computer := &ComputerFacade{computer: &Computer{}}
	computer.TurnOn()
	computer.TurnOff()
}
