package services

import (
	"errors"
	"fmt"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/dao"
	"github.com/noovertime7/gin-mysqlbak/dao/roledao"
	"github.com/noovertime7/gin-mysqlbak/dto"
	"github.com/noovertime7/gin-mysqlbak/public"
	"time"
)

var UserService *userService

type userService struct{}

func (u *userService) Login(ctx *gin.Context, params *dto.AdminLoginInput) (string, error) {
	//获取数据库连接池
	tx, err := lib.GetGormPool("default")
	if err != nil {
		return "", err
	}
	//进行密码校验
	admin := &dao.Admin{}
	admin, err = admin.LoginCheck(ctx, tx, params)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	//生成token
	token, err := public.JWTToken.GenerateToken(&admin.Id)
	if err != nil {
		return "", err
	}
	//更新用户在线状态
	admin.Status = 1
	admin.LoginTime = time.Now()
	if err = admin.UpdateStatus(ctx, tx, admin); err != nil {
		return "", err
	}
	return token, nil
}

func (u *userService) LoginOut(ctx *gin.Context) error {
	//获取数据库连接池
	tx, err := lib.GetGormPool("default")
	if err != nil {
		return err
	}
	claims, exists := ctx.Get("claims")
	if !exists {
		return err
	}
	cla, _ := claims.(*public.CustomClaims)
	adminDB := &dao.Admin{Id: cla.Uid}
	admin, err := adminDB.Find(ctx, tx, adminDB)
	admin.Status = 0
	admin.LoginTime = time.Now()
	if err := admin.UpdateStatus(ctx, tx, admin); err != nil {
		return err
	}
	return nil
}

// ChangePwd 修改密码
func (u *userService) ChangePwd(ctx *gin.Context, params *dto.ChangePwdInput) error {
	//1、通过claims解析用户id
	claims, exists := ctx.Get("claims")
	if !exists {
		return errors.New("claims不存在,请检查jwt中间件")
	}
	cla, _ := claims.(*public.CustomClaims)
	//2、利用结构体中的id去读取数据库信息 adminInfo
	//获取数据库连接池
	tx, err := lib.GetGormPool("default")
	if err != nil {
		return err
	}
	adminInfo := &dao.Admin{}
	adminInfo, err = adminInfo.Find(ctx, tx, &dao.Admin{Id: cla.Uid})
	if err != nil {
		return err
	}
	//使用旧密码进行登录测试
	if _, err := adminInfo.LoginCheck(ctx, tx, &dto.AdminLoginInput{UserName: adminInfo.UserName, Password: params.OldPassword}); err != nil {
		return errors.New("原密码不正确，请重新输入")
	}
	//3、加盐 params.Password + admin-info.salt sha256 saltPassword
	saltPassword := public.GenSaltPassword(adminInfo.Salt, params.Password)
	//4、保存新的password到数据库中
	adminInfo.Password = saltPassword
	if err := adminInfo.Save(ctx, tx); err != nil {
		return err
	}
	return nil
}

// GetUserInfo 获取用户信息
func (u *userService) GetUserInfo(ctx *gin.Context, uid int) (*dto.UserInfoOutPut, error) {
	tx, err := lib.GetGormPool("default")
	if err != nil {
		return nil, err
	}
	adminDB := &dao.Admin{Id: uid}
	admin, err := adminDB.Find(ctx, tx, adminDB)
	//2、取出数据然后封装输出
	roleInfo, err := RuleService.GetRoleInfo(ctx, uid)
	if err != nil {
		return nil, err
	}
	userDB := &roledao.UserInfo{Id: admin.InfoId}
	userInfo, err := userDB.Find(ctx, tx, userDB)
	if err != nil {
		return nil, err
	}
	groupDB := &roledao.UserGroupDB{Id: admin.GroupId}
	group, err := groupDB.Find(ctx, tx, groupDB)
	if err != nil {
		return nil, err
	}
	return &dto.UserInfoOutPut{
		ID:           admin.Id,
		Name:         admin.UserName,
		LoginTime:    admin.LoginTime,
		Avatar:       userInfo.Avatar,
		Introduction: userInfo.Introduction,
		Status:       admin.Status,
		CreatorId:    userInfo.CreateId,
		GroupName:    group.GroupName,
		RoleName:     roleInfo.Name,
	}, nil
}

//
//func (u *userService) DeleteUser(ctx *gin.Context) error {}
//
//func (u *userService) ChangeUserGroup(ctx *gin.Context) error {}
