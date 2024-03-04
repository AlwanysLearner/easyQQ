package request

type AddMemberRequest struct {
	Username  string `json:"username" binding:"required,min=4,max=32" form:"username"`
	GroupName string `json:"groupname" binding:"required,min=4,max=32" form:"groupname"`
}
