package main

import (
	"fmt"
	"github.com/AlwanysLearner/easyQQ/Model"
	"github.com/AlwanysLearner/easyQQ/Router"
	"github.com/gin-gonic/gin"
)

func main() {
	Model.InitDatabase()
	r := gin.Default()
	r.Use(func(ctx *gin.Context) { fmt.Println("---------------------路由启动----------------------") })
	Router.InitRouter(r)
	r.Run(":8080")
}
