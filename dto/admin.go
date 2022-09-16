package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/public"
	"time"
)

type AdminInfoOutput struct {
	*UserInfoOutPut
	Role *RoleInfo `json:"role"`
}

type UserInfoOutPut struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	LoginTime    time.Time `json:"login_time"`
	Avatar       string    `json:"avatar"`
	Introduction string    `json:"introduction"`
	Status       int       `json:"status"`
	CreatorId    string    `json:"creatorId"`
	GroupName    string    `json:"group_name"`
	RoleName     string    `json:"role_name"`
}

type ChangePwdInput struct {
	OldPassword string `form:"old_password" json:"old_password" comment:"旧密码"   validate:"required" example:"123456"`
	Password    string `form:"password" json:"password" comment:"密码"   validate:"required" example:"123456"`
}

type RoleInfo struct {
	CreateTime  time.Time          `json:"createTime"`
	CreatorId   string             `json:"creatorId"`
	Deleted     int                `json:"deleted"`
	Describe    string             `json:"describe"`
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	Permissions []*PermissionsInfo `json:"permissions"`
}

type PermissionsInfo struct {
	RoleId          string                 `json:"roleId"`
	PermissionId    string                 `json:"permissionId"`
	PermissionName  string                 `json:"permissionName"`
	Actions         string                 `json:"actions"`
	ActionEntitySet []*ActionEntitySetInfo `json:"actionEntitySet"`
	ActionList      string                 `json:"actionList"`
	DataAccess      string                 `json:"dataAccess"`
}

type ActionEntitySetInfo struct {
	Action       string `json:"action"`
	Describe     string `json:"describe"`
	DefaultCheck bool   `json:"defaultCheck"`
}

func (a *ChangePwdInput) BindValidParm(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, a)
}
