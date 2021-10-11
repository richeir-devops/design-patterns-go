package facade

import "fmt"

///////////////////////////////////////////
// 门面模式（Facade）:
// 就是把一大堆方法放到另外一个类的一个方法里面一起调用，就是个简单的集成
///////////////////////////////////////////

///////////////////////////////////////////

type Computer struct {
}

func (c *Computer) GetElectricShock() {
	fmt.Println("Ouch!")
}

func (c *Computer) MakeSound() {
	fmt.Println("Beep beep!")
}

func (c *Computer) ShowLoadingScreen() {
	fmt.Println("Loading...")
}

func (c *Computer) Bam() {
	fmt.Println("Ready to be uesd!")
}

func (c *Computer) CloseEverything() {
	fmt.Println("Bup bup bup buzzzz!")
}

func (c *Computer) Sooth() {
	fmt.Println("Zzzzz")
}

func (c *Computer) PullCurrent() {
	fmt.Println("Haaah!")
}

///////////////////////////////////////////

///////////////////////////////////////////

type ComputerFacade struct {
	computer *Computer
}

func (cf *ComputerFacade) TurnOn() {
	fmt.Println("==========Turn On=====")
	cf.computer.GetElectricShock()
	cf.computer.MakeSound()
	cf.computer.ShowLoadingScreen()
	cf.computer.Bam()
	fmt.Println("==========Done========")
}

func (cf *ComputerFacade) TurnOff() {
	fmt.Println("==========Turn Off====")
	cf.computer.CloseEverything()
	cf.computer.PullCurrent()
	cf.computer.Sooth()
	fmt.Println("==========Done========")
}

///////////////////////////////////////////
