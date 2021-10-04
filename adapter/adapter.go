package adapter

import "fmt"

///////////////////////////////////////////
// 可以理解为“HDMI转DVI”的转接器(Adapter)：
// 1、显示器只有 DVI 接口
// 2、显卡有 HDMI 接口和连接线(HDMI)
// 3、需要一个转接器(Adapter)来进行转换，转换器同时实现了 DVI 和 HDMI，
//   也提供了一个转换方法（最重要的），这样就可以达到功能需求
///////////////////////////////////////////

///////////////////////////////////////////

type Lion interface {
	Roar()
}

type AfricanLioe struct {
}

func (lone *AfricanLioe) Roar() {
	fmt.Println("Lion roar!")
}

///////////////////////////////////////////

///////////////////////////////////////////
type WildDog struct {
}

func (dog *WildDog) Bark() {
	fmt.Println("Wilddog bark!")
}

///////////////////////////////////////////

///////////////////////////////////////////

type Hunter struct {
}

func (hunter *Hunter) hunt(lion Lion) {
	lion.Roar()
}

///////////////////////////////////////////

///////////////////////////////////////////

type WildDogAdapter struct {
	wildDog *WildDog
}

// 转换方法
func (dog *WildDogAdapter) Roar() {
	dog.wildDog.Bark()
}

func NewWildDogAdapter(dog *WildDog) *WildDogAdapter {
	return &WildDogAdapter{wildDog: dog}
}

///////////////////////////////////////////
