package Model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
}

func (user *User) Createuser() bool {
	db := DataBaseSessoin()
	result := db.Create(&user)
	if result.Error != nil {
		return false
	}
	return true
}

func (user *User) FinduserByName() *User {
	db := DataBaseSessoin()
	err := db.Where(&User{Username: user.Username}).First(&user).Error
	if err == nil {
		return user
	}
	return nil
}
