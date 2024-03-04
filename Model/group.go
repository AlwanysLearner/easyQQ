package Model

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	GroupName        string
	GroupDescription string
}

func (g *Group) Create() bool {
	db := DataBaseSessoin()
	result := db.Create(&g)
	if result.Error != nil {
		return false
	}
	return true
}

func (g *Group) FindGroupByName() *Group {
	db := DataBaseSessoin()
	err := db.Where(&Group{GroupName: g.GroupName}).First(&g).Error
	if err == nil {
		return g
	}
	return nil
}
