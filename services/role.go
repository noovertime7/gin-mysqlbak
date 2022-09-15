package services

import (
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/dao"
	"github.com/noovertime7/gin-mysqlbak/dao/roledao"
	"github.com/noovertime7/gin-mysqlbak/dto"
	"github.com/noovertime7/gin-mysqlbak/public"
	"github.com/noovertime7/mysqlbak/pkg/log"
)

var RuleService *roleService

type roleService struct{}

func (r *roleService) GetRoleInfo(ctx *gin.Context, uid int) (*dto.RoleInfo, error) {
	tx, err := lib.GetGormPool("default")
	if err != nil {
		return nil, err
	}
	//从ctx中获取用户id
	claims, exists := ctx.Get("claims")
	if !exists {
		log.Logger.Error("claims不存在,请检查jwt中间件")
	}
	cla, _ := claims.(*public.CustomClaims)
	adminDB := &dao.Admin{Id: cla.Uid}
	admin, err := adminDB.Find(ctx, tx, adminDB)
	roleDB := &roledao.RoleDB{Id: admin.Role}
	role, err := roleDB.Find(ctx, tx, roleDB)
	if err != nil {
		return nil, err
	}
	PermissionDB := &roledao.PermissionDB{RoleId: role.RoleId}
	permissionList, err := PermissionDB.FindPermissions(ctx, tx, PermissionDB)
	if err != nil {
		return nil, err
	}
	actionDB := &roledao.ActionDB{PermissionId: role.PermissionId}
	actions, err := actionDB.FindActions(ctx, tx, actionDB)
	if err != nil {
		return nil, err
	}
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
