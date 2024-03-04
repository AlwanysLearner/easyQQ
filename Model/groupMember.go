package Model

type GroupMember struct {
	Username string `gorm:"column:username"`
	GroupID  int    `gorm:"column:group_id"`
}

func (GroupMember) TableName() string {
	return "group_member"
}

func (gm *GroupMember) AddMember() bool {
	db := DataBaseSessoin()
	result := db.Create(&gm)
	if result.Error != nil {
		return false
	}
	return true
}

func (gm *GroupMember) DeleteMember() bool {
	db := DataBaseSessoin()
	result := db.Unscoped().Delete(&gm)
	if result.Error != nil {
		return false
	}
	return true
}

func (gm *GroupMember) IsMember() bool {
	db := DataBaseSessoin()
	err := db.Where(&GroupMember{Username: gm.Username, GroupID: gm.GroupID}).First(&gm).Error
	if err == nil {
		return true
	}
	return false
}

func MemberList(groupid int) []*GroupMember {
	db := DataBaseSessoin()
	var member []*GroupMember
	db.Where("group_id= ?", groupid).Find(&member)
	return member
}
