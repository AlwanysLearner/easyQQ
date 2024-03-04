package Router

import (
	"github.com/AlwanysLearner/easyQQ/Middleware"
	"github.com/AlwanysLearner/easyQQ/Service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(r *gin.Engine) {
	r.Static("/static", "./static")
	apirouter := r.Group("/easyQQ")
	login_apiRouter := apirouter.Group("/", Middleware.VerifyToken())
	apirouter.POST("/user/register/", Service.Register)
	apirouter.POST("/user/login/", Service.Login)
	login_apiRouter.POST("/user/tokenlogin", func(c *gin.Context) {
		username, exists := c.Get("username")
		if !exists {
			// 如果无法获取username，返回错误信息
			c.JSON(http.StatusBadRequest, gin.H{"error": "username not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"msg": "hello" + username.(string)})
	})
	apirouter.GET("/chat/getchat", Service.ChatHandle)
	login_apiRouter.POST("/app/exit/", Service.Exit)
	login_apiRouter.POST("/chat/historymessage/", Service.FindHistoryMessage)
	login_apiRouter.POST("/chat/group/", Service.CreateGroup)
	login_apiRouter.POST("/chat/groupmember/", Service.Addmember)
	login_apiRouter.DELETE("/chat/groupmember/", Service.DeleteMemeber)
}
