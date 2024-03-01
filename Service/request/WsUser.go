package request

import (
	"github.com/AlwanysLearner/easyQQ/Service"
	"github.com/gorilla/websocket"
	"strings"
)

type User struct {
	Username string
	Conn     *websocket.Conn
}

func NewUser(username string, conn *websocket.Conn) *User {
	user := &User{
		Username: username,
		Conn:     conn,
	}
	return user
}

func (this *User) ListenMessage() {
	for {
		_, data, err := this.Conn.ReadMessage()
		if err != nil {
			Service.OutLine(this.Username)
			return
		}
		if string(data) == "who" {
			Service.UserList(this)
		} else if len(data) > 4 && string(data)[:3] == "to|" {
			toname := strings.Split(string(data), "|")[1]
			msg := strings.Split(string(data), "|")[2]
			if toname == "" {
				this.Conn.WriteMessage(websocket.TextMessage, []byte("消息格式不正确，请使用\"to|username|content\"格式。"))
				return
			}
			Service.Chat(this, toname, msg)
		} else {
			this.Conn.WriteMessage(websocket.TextMessage, data)
		}
	}
}
