package main

import (
	"fmt"
	"github.com/AlwanysLearner/easyQQ/Model"
	"github.com/AlwanysLearner/easyQQ/Router"
	"github.com/AlwanysLearner/easyQQ/Service"
	"github.com/AlwanysLearner/easyQQ/redisModel"
	"github.com/gin-gonic/gin"
)

func main() {
	Model.InitDatabase()
	redisModel.InitRedis()
	r := gin.Default()
	r.Use(func(ctx *gin.Context) { fmt.Println("---------------------路由启动----------------------") })
	Router.InitRouter(r)
	go Service.ListenMessage()
	go Service.StoreMessageInMysql()
	r.Run(":8080")
}
