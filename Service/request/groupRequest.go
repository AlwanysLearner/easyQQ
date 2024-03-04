package request

type CreateGroupRequest struct {
	Username         string `json:"username" binding:"required,min=4,max=32" form:"username"`
	GroupName        string `json:"groupname" binding:"required,min=4,max=32" form:"groupname"`
	GroupDescription string `json:"groupDescription" binding:"required,min=4,max=256" form:"groupDescription"`
}
