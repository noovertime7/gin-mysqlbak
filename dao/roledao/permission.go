package roledao

type PermissionDB struct {
	Id             int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	PermissionName string `gorm:"column:permission_name;type:varchar(20)" json:"permission_name"`
	PermissionId   string `gorm:"column:permission_id;type:varchar(20)" json:"permission_id"`
	Actions        string `gorm:"column:actions;type:text" json:"actions"`
	RoleId         string `gorm:"column:role_id;type:varchar(20)" json:"role_id"`
}

func (p *PermissionDB) TableName() string {
	return "t_permission"
}
