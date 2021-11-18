package abstractfactory

import "testing"

func Test01(t *testing.T) {
	woodenFactory := &WoodenDoorFactroy{}
	door := woodenFactory.makeDoor()
	expert := woodenFactory.makeFittingExpert()

	door.getDescription()
	expert.getDescription()
}
