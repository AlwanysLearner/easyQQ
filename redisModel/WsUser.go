package redisModel

import "golang.org/x/net/websocket"

type User struct {
	Username string
	Conn     *websocket.Conn
	Message  chan interface{}
}
