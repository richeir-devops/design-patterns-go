package proxy

import "testing"

func Test01(t *testing.T) {
	door := &SecureDoor{door: &LabDoor{}}
	door.Open("invalid")

	door.Open("123456")
	door.Close()
}
