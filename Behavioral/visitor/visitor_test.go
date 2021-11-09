package visitor

import (
	"fmt"
	"testing"
)

func Test01(t *testing.T) {
	monkey := &Monkey{}
	lion := &Lion{}
	dolphin := &Dolphin{}

	speak := &Speak{}
	monkey.accept(speak)
	lion.accept(speak)
	dolphin.accept(speak)

	fmt.Println("=== fter add some behavior ===")
	jump := &Jump{}
	monkey.accept(jump)
	lion.accept(jump)
	dolphin.accept(jump)
}
