package state

import "testing"

func Test01(t *testing.T) {
	editor := &TextEditor{state: &DefaultTest{}}

	editor.typing("First line")

	editor.state = &UpperCase{}
	editor.typing("Second line")
	editor.typing("Third line")

	editor.state = &LowerCase{}
	editor.typing("Fourth line")
	editor.typing("Fifth line")
}
