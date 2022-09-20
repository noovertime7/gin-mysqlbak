package services

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/dao"
	"github.com/noovertime7/gin-mysqlbak/dao/roledao"
	"github.com/noovertime7/gin-mysqlbak/dto"
	"github.com/noovertime7/gin-mysqlbak/public"
	"github.com/noovertime7/gin-mysqlbak/public/database"
)

var RuleService *roleService

type roleService struct{}

func (r *roleService) GetRoleInfo(ctx *gin.Context, uid int) (*dto.RoleInfo, error) {
	tx := database.GetDB()
	adminDB := &dao.Admin{Id: uid}
	admin, err := adminDB.Find(ctx, tx, adminDB)
	//首先查询用户所属的用户组
	groupDB := &roledao.UserGroupDB{Id: admin.GroupId}
	group, err := groupDB.Find(ctx, tx, groupDB)
	if err != nil {
		return nil, err
	}
	//通过用户组role_id查询权限
	roleDB := &roledao.RoleDB{Id: group.RoleId}
	role, err := roleDB.Find(ctx, tx, roleDB)
	if err != nil {
		return nil, err
	}
	//通过role查询Permission
	PermissionDB := &roledao.PermissionDB{RoleId: role.RoleId}
	permissionList, err := PermissionDB.FindPermissions(ctx, tx, PermissionDB)
	if err != nil {
		return nil, err
	}
	//通过permission查询所属的action
	actionDB := &roledao.ActionDB{PermissionId: role.PermissionId}
	actions, err := actionDB.FindActions(ctx, tx, actionDB)
	if err != nil {
		return nil, err
	}
	//数据组装输出
	var Permissions []*dto.PermissionsInfo
	for _, permission := range permissionList {
		var actionList []*dto.ActionEntitySetInfo
		for _, action := range actions {
			a := &dto.ActionEntitySetInfo{
				Action:       action.Action,
				Describe:     action.Describe,
				DefaultCheck: public.IntToBool(action.DefaultCheck),
			}
			actionList = append(actionList, a)
		}
		p := &dto.PermissionsInfo{
			RoleId:          permission.RoleId,
			PermissionId:    permission.PermissionId,
			PermissionName:  permission.PermissionName,
			Actions:         "",
			ActionEntitySet: actionList,
			ActionList:      "",
			DataAccess:      "",
		}
		Permissions = append(Permissions, p)
	}
	return &dto.RoleInfo{
		CreateTime:  role.CreateAt,
		CreatorId:   role.CreateUser,
		Deleted:     role.IsDeleted,
		Describe:    role.Describe,
		ID:          role.RoleId,
		Name:        role.Name,
		Permissions: Permissions,
	}, nil
}
