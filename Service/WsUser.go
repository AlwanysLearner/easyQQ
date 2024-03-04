package Service

import (
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
			OutLine(this.Username)
			return
		}
		if string(data) == "who" {
			UserList(this)
		} else if len(data) > 4 && string(data)[:3] == "to|" {
			if len(strings.Split(string(data), "|")) != 3 {
				this.Conn.WriteMessage(websocket.TextMessage, []byte("消息格式不正确，请使用\"to|username|content\"格式。"))
				return
			}
			toname := strings.Split(string(data), "|")[1]
			msg := strings.Split(string(data), "|")[2]
			if toname == "" {
				this.Conn.WriteMessage(websocket.TextMessage, []byte("username不能为空"))
				return
			}
			Chat(this, toname, msg)
		} else if len(data) > 7 && string(data)[:6] == "group|" {
			if len(strings.Split(string(data), "|")) != 3 {
				this.Conn.WriteMessage(websocket.TextMessage, []byte("消息格式不正确，请使用\"group|groupname|content\"格式。"))
				return
			}
			groupname := strings.Split(string(data), "|")[1]
			msg := strings.Split(string(data), "|")[2]
			if groupname == "" {
				this.Conn.WriteMessage(websocket.TextMessage, []byte("groupname不能为空"))
				return
			}
			GroupChat(this, groupname, msg)
		} else {
			this.Conn.WriteMessage(websocket.TextMessage, data)
		}
	}
}
