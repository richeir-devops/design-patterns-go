package observer

import "testing"

func Test01(t *testing.T) {
	john := &JobSeeker{name: "John"}
	jane := &JobSeeker{name: "Jane"}

	jobPostings := &EmploymentAgentcy{}
	jobPostings.attach(john)
	jobPostings.attach(jane)

	jobPostings.addJob(&JobPost{title: "Software Engineer"})
}
