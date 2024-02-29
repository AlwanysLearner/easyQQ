package Router

import (
	"github.com/AlwanysLearner/easyQQ/Service"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.Static("/static", "./static")
	apirouter := r.Group("/easyQQ")
	//login_apiRouter := apirouter.Group("/", Middleware.VerifyToken())
	apirouter.POST("/user/register/", Service.Register)
	apirouter.POST("/user/login/", Service.Login)
	apirouter.GET("/chat/", Service.ChatHandle)
}
