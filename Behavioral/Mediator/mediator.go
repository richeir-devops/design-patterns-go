package mediator

import (
	"fmt"
	"time"
)

///////////////////////////////////////////
// 中介者模式（Mediator）：
// 就像聊天室，有一个中枢来处理多方发过来的信息并显示出来
///////////////////////////////////////////

///////////////////////////////////////////

type ChatRoomMediator interface {
	showMessage(*User, string)
}

///////////////////////////////////////////

///////////////////////////////////////////

type ChatRoom struct {
}

func (cr *ChatRoom) showMessage(user *User, message string) {
	fmt.Printf("[%s][%s]: %s \n", time.Now().Format(time.RFC822), user.name, message)
}

///////////////////////////////////////////

///////////////////////////////////////////

type User struct {
	name         string
	chatMediator ChatRoomMediator
}

func (u *User) send(message string) {
	u.chatMediator.showMessage(u, message)
}

///////////////////////////////////////////
