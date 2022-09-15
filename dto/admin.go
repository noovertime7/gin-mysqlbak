package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/public"
	"time"
)

type AdminInfoOutput struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	LoginTime    time.Time `json:"login_time"`
	Avatar       string    `json:"avatar"`
	Introduction string    `json:"introduction"`
	Status       int       `json:"status"`
	CreatorId    string    `json:"creatorId"`
	Role         *RoleInfo `json:"role"`
}

type ChangePwdInput struct {
	Password string `form:"password" json:"password" comment:"密码"   validate:"required" example:"123456"`
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

func InitDemoInfo() *RoleInfo {
	return &RoleInfo{
		CreateTime: time.Now(),
		CreatorId:  "system",
		Deleted:    0,
		Describe:   "拥有所有权限",
		ID:         "admin",
		Name:       "管理员",
		Permissions: []*PermissionsInfo{
			{
				RoleId:         "admin",
				PermissionId:   "dashboard",
				PermissionName: "仪表盘",
				Actions:        "[{\"action\":\"add\",\"defaultCheck\":false,\"describe\":\"新增\"},{\"action\":\"query\",\"defaultCheck\":false,\"describe\":\"查询\"},{\"action\":\"get\",\"defaultCheck\":false,\"describe\":\"详情\"},{\"action\":\"update\",\"defaultCheck\":false,\"describe\":\"修改\"},{\"action\":\"delete\",\"defaultCheck\":false,\"describe\":\"删除\"}]",
				ActionEntitySet: []*ActionEntitySetInfo{
					{
						Action:       "add",
						Describe:     "add",
						DefaultCheck: false,
					},
					{
						Action:       "query",
						Describe:     "查询",
						DefaultCheck: false,
					},
					{
						Action:       "get",
						Describe:     "详情",
						DefaultCheck: false,
					},
					{
						Action:       "update",
						Describe:     "修改",
						DefaultCheck: false,
					},
					{
						Action:       "delete",
						Describe:     "删除",
						DefaultCheck: false,
					},
				},
			},
		},
	}
}
