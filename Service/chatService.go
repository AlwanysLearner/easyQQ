package Service

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

// 创建一个全局的连接映射，用于存储每个用户对应的 WebSocket 连接
var connections = make(map[string]*websocket.Conn)
var mu sync.RWMutex // 用于保护连接映射的读写操作
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// 允许所有的WebSocket连接请求
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func ChatHandle(c *gin.Context) {
	// 获取用户名
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Username not found"})
		return
	}

	// 升级 WebSocket 连接
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Could not open websocket connection"})
		return
	}

	// 存储连接信息
	mu.Lock()
	connections[username.(string)] = conn
	mu.Unlock()

	// 处理 WebSocket 连接
	handleWebSocket(conn)
}

func handleWebSocket(conn *websocket.Conn) {
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}
		// 处理收到的消息
		handleMessage(msg)
	}
}

func handleMessage(msg []byte) {
	// 处理收到的消息
	log.Println("Received message:", string(msg))
}

func Exit(c *gin.Context) {
	username := c.PostForm("username")
	connections[username].Close()
	mu.Lock()
	delete(connections, username)
	mu.Unlock()
	c.JSON(http.StatusOK, gin.H{"msg": "退出成功"})
}
