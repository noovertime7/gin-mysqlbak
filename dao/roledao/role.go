package roledao

import "time"

// RoleDB TRole 角色表
type RoleDB struct {
	Id           int       `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Name         string    `gorm:"column:name;type:varchar(20)" json:"name"`
	RoleId       string    `gorm:"column:role_id;type:varchar(20)" json:"role_id"`
	Describe     string    `gorm:"column:describe;type:varchar(20)" json:"describe"`
	CreateAt     time.Time `gorm:"column:create_at;type:datetime" json:"create_at"`
	IsDeleted    int       `gorm:"column:is_deleted;type:int(11)" json:"is_deleted"`
	PermissionId int       `gorm:"column:permission_id;type:int(11)" json:"permission_id"`
}

func (t *RoleDB) TableName() string {
	return "t_role"
}
