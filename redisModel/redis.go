package redisModel

import (
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"os"
)

type RedisConfig struct {
	Addr     string `json:"Addr"`
	Password string `json:"Password"`
	DB       int    `json:"DB"`
}

var rdb *redis.Client

func InitRedis() {
	conf, _ := os.Open("config/redis.json")
	defer conf.Close() //执行完毕后关闭连接
	var config RedisConfig
	jsonParser := json.NewDecoder(conf)
	if err := jsonParser.Decode(&config); err != nil {
		panic(err)
	}
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: string(config.Password), // 如果没有设置密码，则留空
		DB:       config.DB,               // 使用默认的数据库
	})
}
func RedisSession() *redis.Client {
	return rdb
}
