package simplefactory

import "testing"

func Test01(t *testing.T) {
	doorFactory := &DoorFactory{}
	door1 := doorFactory.makeDoor(100, 200)
	door2 := doorFactory.makeDoor(50, 100)

	t.Logf("Door1: %.f width, %.f width \n", door1.Width, door1.Height)
	t.Logf("Door2: %.f width, %.f width \n", door2.Width, door2.Height)
}
