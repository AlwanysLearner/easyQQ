package Service

import (
	"github.com/AlwanysLearner/easyQQ/Middleware"
	"github.com/AlwanysLearner/easyQQ/Model"
	"github.com/AlwanysLearner/easyQQ/redisModel"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
	"time"
)

// 创建一个全局的连接映射，用于存储每个用户对应的 WebSocket 连接
var OnlineMap = make(map[string]*websocket.Conn)
var mapLock sync.RWMutex // 用于保护连接映射的读写操作
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// 允许所有的WebSocket连接请求
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
var Message = make(chan string)
var MysqlMessage = make(chan *Model.Message, 100)

func ChatHandle(c *gin.Context) {
	token := c.Query("token")
	if ok, claims := Middleware.CheckToken(token); ok {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println("WebSocket upgrade error:", err)
			c.JSON(http.StatusBadRequest, gin.H{"msg": "Could not open websocket connection"})
			return
		}
		user := NewUser(claims.Username, conn)
		// 存储连接信息
		mapLock.Lock()
		OnlineMap[user.Username] = user.Conn
		mapLock.Unlock()

		// 处理 WebSocket 连接
		go user.ListenMessage()
		BroadCast(user.Username, "已上线")
		select {}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "未登录请先登录"})
	}
}

func StoreMessageInMysql() {
	for {
		m := <-MysqlMessage
		for !m.Store() {
		}
	}
}

func ListenMessage() {
	for {
		msg := <-Message
		mapLock.RLock()
		for k, cli := range OnlineMap {
			cli.WriteMessage(websocket.TextMessage, []byte(msg))
			for !redisModel.StoreMessage(&redisModel.Message{Time: float64(time.Now().Unix()), Msg: msg}, k) {
			}
			m := &Model.Message{Msg: msg, Time: float64(time.Now().Unix()), Username: k}
			MysqlMessage <- m
		}
		mapLock.RUnlock()
	}
}

func BroadCast(username string, msg string) {
	sendMsg := "[" + username + "]" + username + ":" + msg
	Message <- sendMsg
}
func Exit(c *gin.Context) {
	username := c.PostForm("username")
	OutLine(username)
	c.JSON(http.StatusOK, gin.H{"msg": username + "已下线"})
}
func OutLine(username string) {
	BroadCast(username, "已下线")
	mapLock.Lock()
	OnlineMap[username].Close()
	delete(OnlineMap, username)
	mapLock.Unlock()

}
func UserList(user *User) {
	msg := ""
	mapLock.RLock()
	for k, _ := range OnlineMap {
		if k != user.Username {
			msg += "[" + k + "]" + k + ":" + "在线...\n"
		}
	}
	mapLock.RUnlock()
	user.Conn.WriteMessage(websocket.TextMessage, []byte(msg))
}

func Chat(user *User, username string, msg string) {
	mapLock.RLock()
	u, ok := OnlineMap[username]
	mapLock.RUnlock()
	msg = user.Username + ":" + msg
	if ok {
		u.WriteMessage(websocket.TextMessage, []byte(msg))
	}
	for !redisModel.StoreMessage(&redisModel.Message{Time: float64(time.Now().Unix()), Msg: msg}, username) {
	}
	m := &Model.Message{Msg: msg, Time: float64(time.Now().Unix()), Username: username}
	MysqlMessage <- m
}

func GroupChat(user *User, groupName string, msg string) {
	group := &Model.Group{GroupName: groupName}
	group = group.FindGroupByName()
	if group == nil {
		user.Conn.WriteMessage(websocket.TextMessage, []byte("群不存在..."))
	}
	member := Model.MemberList(int(group.ID))
	for _, v := range member {
		if v.Username != user.Username {
			Chat(user, v.Username, msg)
		}
	}
}

func FindHistoryMessage(c *gin.Context) {
	username := c.PostForm("username")
	msg, err := redisModel.HistoryMessage(username)
	if err != nil {
		log.Print("查询出错", err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": "查询错误"})
	}
	c.JSON(http.StatusOK, gin.H{"msg": msg})
}
