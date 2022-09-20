package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/dto"
	"github.com/noovertime7/gin-mysqlbak/middleware"
	"github.com/noovertime7/gin-mysqlbak/public"
	"github.com/noovertime7/gin-mysqlbak/services"
	"github.com/noovertime7/mysqlbak/pkg/log"
)

type AdminController struct {
}

func AdminRegister(group *gin.RouterGroup) {
	admininfo := &AdminController{}
	group.DELETE("/user_delete", admininfo.DeleteUser)
	group.GET("/user_reset_pwd", admininfo.ResetUserPasswd)
	group.GET("/role_info", admininfo.RoleInfo)
	group.GET("/user_info", admininfo.UserInfo)
	group.PUT("/user_info_update", admininfo.UpdateUserInfo)
	group.GET("/userinfo_by_group", admininfo.GetUserByGroup)
	group.GET("/user_group_list", admininfo.GetUserGroupList)
	group.POST("/changepwd", admininfo.ChangePwd)
}

func (a *AdminController) UserInfo(ctx *gin.Context) {
	Iuid, exists := ctx.Get("uid")
	if !exists {
		log.Logger.Error("uid不存在,请检查jwt中间件")
	}
	uid := Iuid.(int)
	userinfo, err := services.UserService.GetUserInfo(ctx, uid)
	if err != nil {
		log.Logger.Error("查询用户信息失败", err)
		middleware.ResponseError(ctx, 2002, err)
		return
	}
	middleware.ResponseSuccess(ctx, userinfo)
}

func (a *AdminController) RoleInfo(ctx *gin.Context) {
	Iuid, exists := ctx.Get("uid")
	if !exists {
		log.Logger.Error("uid不存在,请检查jwt中间件")
	}
	uid := Iuid.(int)
	roleInfo, err := services.RuleService.GetRoleInfo(ctx, uid)
	if err != nil {
		log.Logger.Error("查询用户权限失败", err)
		middleware.ResponseError(ctx, 2002, err)
		return
	}
	userinfo, err := services.UserService.GetUserInfo(ctx, uid)
	if err != nil {
		log.Logger.Error("查询用户信息失败", err)
		middleware.ResponseError(ctx, 2002, err)
		return
	}
	out := &dto.AdminInfoOutput{
		UserInfoOutPut: userinfo,
		Role:           roleInfo,
	}
	middleware.ResponseSuccess(ctx, out)
}

// ChangePwd 更改密码
func (a *AdminController) ChangePwd(ctx *gin.Context) {
	params := &dto.ChangePwdInput{}
	if err := params.BindValidParams(ctx); err != nil {
		middleware.ResponseError(ctx, public.ParamsBindErrorCode, err)
		return
	}
	if err := services.UserService.ChangePwd(ctx, params); err != nil {
		log.Logger.Error("修改密码失败", err)
		middleware.ResponseError(ctx, 30002, err)
		return
	}
	middleware.ResponseSuccess(ctx, "更改密码成功")
}

func (a *AdminController) GetUserGroupList(ctx *gin.Context) {
	out, err := services.UserService.GetUserGroupList(ctx)
	if err != nil {
		log.Logger.Error("查询业务组失败", err)
		middleware.ResponseError(ctx, 30002, err)
		return
	}
	middleware.ResponseSuccess(ctx, out)
}

func (a *AdminController) GetUserByGroup(ctx *gin.Context) {
	params := &dto.GroupUserListInput{}
	if err := params.BindValidParams(ctx); err != nil {
		middleware.ResponseError(ctx, public.ParamsBindErrorCode, err)
		return
	}
	out, err := services.UserService.FindUserByGroup(ctx, params)
	if err != nil {
		log.Logger.Error("根据group查询用户信息失败", err)
		middleware.ResponseError(ctx, 30002, err)
		return
	}
	middleware.ResponseSuccess(ctx, out)
}

func (a *AdminController) UpdateUserInfo(ctx *gin.Context) {
	params := &dto.UpdateUserInfo{}
	if err := params.BindValidParams(ctx); err != nil {
		log.Logger.Error("绑定失败", err)
		middleware.ResponseError(ctx, public.ParamsBindErrorCode, err)
		return
	}
	if err := services.UserService.UpdateUserInfo(ctx, params); err != nil {
		log.Logger.Error("更新失败")
		middleware.ResponseError(ctx, 30003, err)
		return
	}
	middleware.ResponseSuccess(ctx, "更新成功")
}

// ResetUserPasswd 重置用户密码
func (a *AdminController) ResetUserPasswd(ctx *gin.Context) {
	params := &dto.UserIDInput{}
	if err := params.BindValidParams(ctx); err != nil {
		log.Logger.Error("绑定失败", err)
		middleware.ResponseError(ctx, public.ParamsBindErrorCode, err)
		return
	}
	if err := services.UserService.ResetUserPassword(ctx, params); err != nil {
		log.Logger.Error("重置密码失败", err)
		middleware.ResponseError(ctx, 30004, err)
		return
	}
	middleware.ResponseSuccess(ctx, "重置密码成功,默认密码:admin@123")
}

// DeleteUser 删除用户
func (a *AdminController) DeleteUser(ctx *gin.Context) {
	params := &dto.UserIDInput{}
	if err := params.BindValidParams(ctx); err != nil {
		log.Logger.Error("绑定失败", err)
		middleware.ResponseError(ctx, public.ParamsBindErrorCode, err)
		return
	}
	if err := services.UserService.DeleteUser(ctx, params); err != nil {
		log.Logger.Error("删除用户失败", err)
		middleware.ResponseError(ctx, 30006, err)
		return
	}
	middleware.ResponseSuccess(ctx, "删除成功")
}
