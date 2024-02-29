package Service

import (
	"github.com/AlwanysLearner/easyQQ/Middleware"
	"github.com/AlwanysLearner/easyQQ/Model"
	"github.com/AlwanysLearner/easyQQ/Service/request"
	"github.com/AlwanysLearner/easyQQ/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var req request.RegisterLoginRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := &Model.User{Username: req.Username}
	user.FinduserByName()
	if utils.VerifyPassword(req.Password, user.Password) == false {
		c.JSON(http.StatusExpectationFailed, gin.H{"error": "用户名密码错误"})
		return
	}
	token, err := Middleware.GenerateToken(req.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "token生成失败"})
	}
	c.JSON(http.StatusOK, gin.H{"msg": "登录成功", "token": token})
}

func Register(c *gin.Context) {
	var req request.RegisterLoginRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := &Model.User{Username: req.Username}
	user.FinduserByName()
	if user.Password != "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "用户名已存在"})
		return
	}
	user = &Model.User{Username: req.Username, Password: utils.BcryptEncode(req.Password)}
	if user.Createuser() {
		c.JSON(http.StatusOK, gin.H{"msg": "注册成功"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "注册失败"})
	}
}
