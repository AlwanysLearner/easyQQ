package redisModel

import (
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
	"log"
	"time"
)

type Message struct {
	Msg  string
	Time float64
}

func StoreMessage(m *Message, username string) bool {
	rdb := RedisSession()
	ctx := context.Background()
	_, err := rdb.ZAdd(ctx, "chat:"+username, &redis.Z{
		Score:  m.Time,
		Member: m.Msg,
	}).Result()
	if err != nil {
		log.Print("Error adding message to sorted set:", err)
		return false
	}
	//设置redis生命周期
	_, err = rdb.Expire(ctx, "chat:"+username, 60*time.Minute).Result()
	if err != nil {
		log.Print("Error setting expiration for sorted set:", err)
		return false
	}
	return true
}

func HistoryMessage(username string) ([]string, error) {
	ctx := context.Background()
	members, err := rdb.ZRange(ctx, "chat:"+username, 0, -1).Result()
	if err != nil {
		return nil, err
	}
	return members, nil
}
