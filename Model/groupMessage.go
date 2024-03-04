package Model

type GroupMessage struct {
	GroupID  int     `gorm:"column:group_id"`
	Username string  `gorm:"column:username"`
	Msg      string  `gorm:"column:msg"`
	Time     float64 `gorm:"column:time"`
}

func (GroupMessage) TableName() string {
	return "group_message"
}

func (gmsg *GroupMessage) StoreMessage() bool {
	db := DataBaseSessoin()
	result := db.Create(&gmsg)
	if result.Error != nil {
		return false
	}
	return true
}
