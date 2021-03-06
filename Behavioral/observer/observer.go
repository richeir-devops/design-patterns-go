package observer

import "fmt"

///////////////////////////////////////////
// 观察者模式（Observer）：
// 观察者模式其实就是把观察者都放一起，然后统一给消息
///////////////////////////////////////////

///////////////////////////////////////////

type Observable interface {
	onJobPosted(*JobPost)
}

type JobPost struct {
	title string
}

///////////////////////////////////////////

///////////////////////////////////////////

type JobSeeker struct {
	name string
}

func (js *JobSeeker) onJobPosted(job *JobPost) {
	fmt.Printf("Hi, %s! New job posted: %s \n", js.name, job.title)
}

///////////////////////////////////////////

///////////////////////////////////////////

type EmploymentAgentcy struct {
	observers []Observable
}

func (ea *EmploymentAgentcy) notify(jobPosting *JobPost) {
	for _, obs := range ea.observers {
		obs.onJobPosted(jobPosting)
	}
}

func (ea *EmploymentAgentcy) attach(observer Observable) {
	ea.observers = append(ea.observers, observer)
}

func (ea *EmploymentAgentcy) addJob(jobPosting *JobPost) {
	ea.notify(jobPosting)
}

///////////////////////////////////////////
