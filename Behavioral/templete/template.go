package templete

import "fmt"

///////////////////////////////////////////
// 模板方法模式（Template Method）：
// 定义标准流程的抽象类，然后让人实现对应的方法，之后再执行统一流程化的东西
///////////////////////////////////////////

///////////////////////////////////////////

type builder interface {
	build()
	test()
	lint()
	assemble()
	deploy()
}

///////////////////////////////////////////

///////////////////////////////////////////

type AndroidBuilder struct {
}

func (ab *AndroidBuilder) test() {
	fmt.Println("Running android tests")
}

func (ab *AndroidBuilder) lint() {
	fmt.Println("Linting the android code")
}

func (ab *AndroidBuilder) assemble() {
	fmt.Println("Assembling the android build")
}

func (ab *AndroidBuilder) deploy() {
	fmt.Println("Deploying android build to server")
}

func (ab *AndroidBuilder) build() {
	ab.test()
	ab.lint()
	ab.assemble()
	ab.deploy()
}

///////////////////////////////////////////

///////////////////////////////////////////

type IosBuilder struct {
}

func (ib *IosBuilder) test() {
	fmt.Println("Running ios tests")
}

func (ib *IosBuilder) lint() {
	fmt.Println("Linting the ios code")
}

func (ib *IosBuilder) assemble() {
	fmt.Println("Assembling the ios build")
}

func (ib *IosBuilder) deploy() {
	fmt.Println("Deploying ios build to server")
}

func (ib *IosBuilder) build() {
	ib.test()
	ib.lint()
	ib.assemble()
	ib.deploy()
}

///////////////////////////////////////////
