package composite

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	john := &Developer{}
	john.Constructor("John", 12000)

	jane := &Designer{}
	jane.Constructor("Jane", 15000)

	org := &Organization{}
	org.AddEmployee(john)
	org.AddEmployee(jane)

	fmt.Println(org.GetNetSalaries())
}
