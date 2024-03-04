package Service

import (
	"github.com/AlwanysLearner/easyQQ/Model"
	"github.com/AlwanysLearner/easyQQ/Service/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateGroup(c *gin.Context) {
	var req request.CreateGroupRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	group := &Model.Group{GroupName: req.GroupName, GroupDescription: req.GroupDescription}
	if group.FindGroupByName() != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"err": "该群已存在，请重新命名"})
	}
	if group.Create() {
		gm := &Model.GroupMember{GroupID: int(group.ID), Username: req.Username}
		gm.AddMember()
		c.JSON(http.StatusOK, gin.H{"msg": "创建成功"})
	} else {
		c.JSON(http.StatusExpectationFailed, gin.H{"msg": "创建失败"})
	}
}

func Addmember(c *gin.Context) {
	var req request.CreateGroupRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	group := &Model.Group{GroupName: req.GroupName}
	if group.FindGroupByName() == nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"err": "该群不存在"})
		return
	}
	gm := &Model.GroupMember{GroupID: int(group.ID), Username: req.Username}
	gm.AddMember()
	c.JSON(http.StatusOK, gin.H{"msg": "添加成功"})
}

func DeleteMemeber(c *gin.Context) {
	var req request.CreateGroupRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	group := &Model.Group{GroupName: req.GroupName}
	if group.FindGroupByName() == nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"err": "该群不存在"})
		return
	}
	gm := &Model.GroupMember{GroupID: int(group.ID), Username: req.Username}
	gm.DeleteMember()
	c.JSON(http.StatusOK, gin.H{"msg": "删除成功"})
}
