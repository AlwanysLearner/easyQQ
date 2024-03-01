package Model

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Username string
	Msg      string
	Time     float64
	//redisModel.Message
}

func (m *Message) Store() bool {
	db := DataBaseSessoin()
	result := db.Create(&m)
	if result.Error != nil {
		return false
	}
	return true
}
