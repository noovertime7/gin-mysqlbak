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

type UserGroupOutPut struct {
	Title         string           `json:"title"`
	Key           string           `json:"key"`
	UserGroupItem []*UserGroupItem `json:"children"`
}

type UserGroupItem struct {
	Title string `json:"title"`
	Key   string `json:"key"`
}

//根据组查询所属用户

type GroupUserListInput struct {
	Info     string `form:"info" json:"info" comment:"关键词"   validate:"" example:""`
	PageNo   int    `form:"page_no" json:"page_no" comment:"每页条数"   validate:"" example:"1"`
	PageSize int    `form:"page_size" json:"page_size" comment:"页数"   validate:"" example:"20"`
	Key      string `form:"key" json:"key" comment:"组key" validate:"" `
}

type GroupUserListOutPut struct {
	Total    int               `form:"total" json:"total" comment:"总数"   validate:"" example:""`
	PageNo   int               `form:"page_no" json:"page_no" comment:"当前页数"   validate:"" example:"1"`
	PageSize int               `form:"page_size" json:"page_size" comment:"页数"   validate:"" example:"20"`
	List     []*UserInfoOutPut `json:"list"`
}

type UserInfoOutPut struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	LoginTime    string `json:"login_time"`
	Avatar       string `json:"avatar"`
	Introduction string `json:"introduction"`
	Status       int    `json:"status"`
	CreatorId    string `json:"creatorId"`
	GroupName    string `json:"group_name"`
	RoleName     string `json:"role_name"`
}

//绑定参数

func (a *ChangePwdInput) BindValidParams(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, a)
}

func (a *GroupUserListInput) BindValidParams(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, a)
}
