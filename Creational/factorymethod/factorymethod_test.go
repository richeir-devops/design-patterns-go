package factorymethod

import "testing"

func Test01(t *testing.T) {
	devManager := &DevelopmentManager{}
	devManager.makeInterviewer()
	devManager.takeInterview()
}
