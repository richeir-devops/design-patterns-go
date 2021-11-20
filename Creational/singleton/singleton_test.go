package singleton

import "testing"

func Test01(t *testing.T) {
	president := &President{}
	president1 := president.getInstance()
	president2 := president.getInstance()

	if president1 == president2 {
		t.Log("They are same.")
	}
}
