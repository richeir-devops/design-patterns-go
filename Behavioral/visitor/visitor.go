package visitor

import "fmt"

///////////////////////////////////////////
// 访问者模式（Visitor）：
// 其实就像一个小型的外挂部件，跟结构体本身没啥关联，只是执行预留的动作
// 但是在给结构体加东西的时候，可以去实现预留的动作，然后去调用
///////////////////////////////////////////

///////////////////////////////////////////

type Animal interface {
	accept(operation AnimalOperation)
}

type AnimalOperation interface {
	visitMonkey(*Monkey)
	visitLion(*Lion)
	visitDolphin(*Dolphin)
}

///////////////////////////////////////////

///////////////////////////////////////////

type Monkey struct {
}

func (m *Monkey) shout() {
	fmt.Println("Ooh oo aa aa!")
}

func (m *Monkey) accept(operation AnimalOperation) {
	operation.visitMonkey(m)
}

///////////////////////////////////////////

///////////////////////////////////////////

type Lion struct {
}

func (l *Lion) roar() {
	fmt.Println("Roaaar!")
}

func (l *Lion) accept(operation AnimalOperation) {
	operation.visitLion(l)
}

///////////////////////////////////////////

///////////////////////////////////////////

type Dolphin struct {
}

func (d *Dolphin) speak() {
	fmt.Println("Tuut tuttu tuutt!")
}

func (d *Dolphin) accept(operation AnimalOperation) {
	operation.visitDolphin(d)
}

///////////////////////////////////////////

///////////////////////////////////////////

type Speak struct {
}

func (s *Speak) visitMonkey(monkey *Monkey) {
	monkey.shout()
}

func (s *Speak) visitLion(lion *Lion) {
	lion.roar()
}

func (s *Speak) visitDolphin(dolphin *Dolphin) {
	dolphin.speak()
}

///////////////////////////////////////////

///////////////////////////////////////////

type Jump struct {
}

func (j *Jump) visitMonkey(monkey *Monkey) {
	fmt.Println("Jumped 20 feet high! on to the tree!")
}

func (j *Jump) visitLion(lion *Lion) {
	fmt.Println("Jumped 7 feet high! Back on the ground!")
}

func (j *Jump) visitDolphin(dolphin *Dolphin) {
	fmt.Println("Walked on water a little and disappeared!")
}

///////////////////////////////////////////
