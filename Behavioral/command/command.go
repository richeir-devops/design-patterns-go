package command

import "fmt"

///////////////////////////////////////////
// 命令行模式（Command）：
// 可以理解正在操作遥控器，上面的每一个按键实际上都实现了相同的接口；
// 这样就可以通过“按下按键”这个统一的动作来实现各种操作
///////////////////////////////////////////

///////////////////////////////////////////

type Command interface {
	execute()
	undo()
	redo()
}

///////////////////////////////////////////

///////////////////////////////////////////

type Bulb struct {
}

func (b *Bulb) trunOn() {
	fmt.Println("Bulb has been lit")
}

func (b *Bulb) turnOff() {
	fmt.Println("Darkness!")
}

///////////////////////////////////////////

///////////////////////////////////////////

type TurnOn struct {
	Bulb
}

func (to *TurnOn) execute() {
	to.Bulb.trunOn()
}

func (to *TurnOn) undo() {
	to.Bulb.turnOff()
}

func (to *TurnOn) redo() {
	to.execute()
}

///////////////////////////////////////////

///////////////////////////////////////////

type TurnOff struct {
	Bulb
}

func (to *TurnOff) execute() {
	to.Bulb.turnOff()
}

func (to *TurnOff) undo() {
	to.Bulb.trunOn()
}

func (to *TurnOff) redo() {
	to.execute()
}

///////////////////////////////////////////

///////////////////////////////////////////

type RemoteControl struct {
}

func (rc *RemoteControl) submit(cmd Command) {
	cmd.execute()
}

///////////////////////////////////////////
