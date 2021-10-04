package bridge

import (
	"fmt"
)

///////////////////////////////////////////
// 桥梁模式（Bridge）:
// 最大的感觉就是这算是依赖注入的最基本的原理；
// 在构造函数里面加入很多内容，这样就可以使用简单的组合来灵活地控制行为
///////////////////////////////////////////

///////////////////////////////////////////

type WebPage interface {
	Construct(string)
	GetContent() string
}

type Theme interface {
	GetColor() string
}

///////////////////////////////////////////

///////////////////////////////////////////

type DarkTheme struct {
}

func (t *DarkTheme) GetColor() string {
	return "Dark Black"
}

type LightTheme struct {
}

func (t *LightTheme) GetColor() string {
	return "Star Light"
}

///////////////////////////////////////////

///////////////////////////////////////////

type About struct {
	theme string
}

func (p *About) Construct(theme Theme) {
	p.theme = theme.GetColor()
}

func (p *About) GetContent() string {
	return fmt.Sprint("About page in ", p.theme)
}

///////////////////////////////////////////

///////////////////////////////////////////

type Careers struct {
	theme string
}

func (p *Careers) Construct(theme Theme) {
	p.theme = theme.GetColor()
}

func (p *Careers) GetContent() string {
	return fmt.Sprint("Careers  page in ", p.theme)
}

///////////////////////////////////////////
