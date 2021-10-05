package bridge

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	darkTheme := &DarkTheme{}
	lightTheme := &LightTheme{}

	aboutPage := &About{}
	aboutPage.Construct(darkTheme)

	careerPage := &Careers{}
	careerPage.Construct(lightTheme)

	fmt.Println(aboutPage.GetContent())
	fmt.Println(careerPage.GetContent())
}
