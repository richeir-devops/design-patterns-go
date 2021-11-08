package memento

import (
	"fmt"
	"testing"
)

func Test01(t *testing.T) {
	editor := &Editor{}

	editor.write("This is the first sentence.")
	editor.write("This is second.")

	saved := editor.save()

	editor.write("And this is third.")

	fmt.Println(editor.content)

	editor.resotre(saved)

	fmt.Println(editor.content)
}
