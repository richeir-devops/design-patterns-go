package factorymethod

import "fmt"

///////////////////////////////////////////
// 工厂方法模式（Factory Method）：
// 基于抽象类和构造函数来实现示例的快速创建
// go 里面并没有这两个概念，所以没法用
///////////////////////////////////////////

///////////////////////////////////////////

type Interviewer interface {
	askQuestions()
}

///////////////////////////////////////////

///////////////////////////////////////////

type Developer struct {
}

func (d *Developer) askQuestions() {
	fmt.Println("Asking about design patterns!")
}

///////////////////////////////////////////

///////////////////////////////////////////

type CommunityExecutive struct {
}

func (ce *CommunityExecutive) askQuestions() {
	fmt.Println("Asking about community building")
}

///////////////////////////////////////////

///////////////////////////////////////////

type HiringManager struct {
	iv Interviewer
}

func (hr *HiringManager) makeInterviewer(iv Interviewer) {
	hr.iv = iv
}

func (hr *HiringManager) takeInterview() {
	hr.iv.askQuestions()
}

///////////////////////////////////////////

///////////////////////////////////////////

type DevelopmentManager struct {
	HiringManager
}

func (dm *DevelopmentManager) makeInterviewer() {
	dm.HiringManager.makeInterviewer(&Developer{})
}

///////////////////////////////////////////
