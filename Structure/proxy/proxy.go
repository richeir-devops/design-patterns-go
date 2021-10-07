package proxy

import "fmt"

///////////////////////////////////////////
// 代理模式（Proxy）:
// 在另一个类的基础方法上进行了整合，加了很多额外的操作
///////////////////////////////////////////

///////////////////////////////////////////

type Door interface {
	Open()
	Close()
}

///////////////////////////////////////////

///////////////////////////////////////////

type LabDoor struct {
}

func (ld *LabDoor) Open() {
	fmt.Println("Opening lab door")
}

func (ld *LabDoor) Close() {
	fmt.Println("Closing lab door")
}

///////////////////////////////////////////

///////////////////////////////////////////

type SecureDoor struct {
	door Door
}

func (sd *SecureDoor) Open(passowrd string) {
	if passowrd == "123456" {
		sd.door.Open()
	} else {
		fmt.Println("Big no! It ain't possible.")
	}
}

func (sd *SecureDoor) Close() {
	sd.door.Close()
}

///////////////////////////////////////////
