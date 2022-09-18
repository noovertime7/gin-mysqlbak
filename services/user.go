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
	//2、利用结构体中的id去读取数据库信息 adminInfo
	//获取数据库连接池
	tx, err := lib.GetGormPool("default")
	if err != nil {
		return err
	}
	adminInfo := &dao.Admin{Id: params.ID}
	adminInfo, err = adminInfo.Find(ctx, tx, adminInfo)
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
		LoginTime:    admin.LoginTime.Format("2006年01月02日15:04:01"),
		Avatar:       userInfo.Avatar,
		Introduction: userInfo.Introduction,
		Status:       admin.Status,
		CreatorId:    userInfo.CreateId,
		GroupName:    group.GroupName,
		RoleName:     roleInfo.Name,
	}, nil
}

func (u *userService) GetUserGroupList(ctx *gin.Context) (*dto.UserGroupOutPut, error) {
	tx, err := lib.GetGormPool("default")
	if err != nil {
		return nil, err
	}
	groupDB := &roledao.UserGroupDB{}
	groups, err := groupDB.FindList(ctx, tx, groupDB)
	if err != nil {
		return nil, err
	}
	var outI []*dto.UserGroupItem
	for _, group := range groups {
		outItem := &dto.UserGroupItem{
			Title: group.GroupName,
			Key:   group.Key,
		}
		outI = append(outI, outItem)
	}
	return &dto.UserGroupOutPut{
		Title:         "用户组",
		Key:           "group",
		UserGroupItem: outI,
	}, nil
}

func (u *userService) FindUserByGroup(ctx *gin.Context, info *dto.GroupUserListInput) (*dto.GroupUserListOutPut, error) {
	tx, err := lib.GetGormPool("default")
	if err != nil {
		return nil, err
	}
	//通过key查询所属的group
	var groupId int
	if info.Key != "" {
		groupDB := &roledao.UserGroupDB{Key: info.Key}
		group, err := groupDB.Find(ctx, tx, groupDB)
		if err != nil {
			return nil, err
		}
		groupId = group.Id
	} else {
		groupId = 0
	}
	//查询在当前组下的所有用户
	userDB := &dao.Admin{}
	list, total, err := userDB.PageList(ctx, tx, info, groupId)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		//当前组没有任何用户，直接return
		return &dto.GroupUserListOutPut{
			Total:    total,
			List:     []dto.UserInfoOutPut{},
			PageNo:   info.PageNo,
			PageSize: info.PageSize,
		}, nil
	}
	var out []dto.UserInfoOutPut
	for _, user := range list {
		// 因为key可能为空，所以要重新查询一下group
		groupDB := &roledao.UserGroupDB{Id: user.GroupId}
		group, err := groupDB.Find(ctx, tx, groupDB)
		if err != nil {
			return nil, err
		}
		//需要用户的userinfo 信息，查询userinfo表
		userInfoDB := &roledao.UserInfo{Id: user.InfoId}
		info, err := userInfoDB.Find(ctx, tx, userInfoDB)
		if err != nil {
			return nil, err
		}
		//需要用户的role信息，查询role表
		roleDB := &roledao.RoleDB{Id: group.RoleId}
		role, err := roleDB.Find(ctx, tx, roleDB)
		if err != nil {
			return nil, err
		}
		outTemp := dto.UserInfoOutPut{
			ID:           user.Id,
			Name:         user.UserName,
			LoginTime:    user.LoginTime.Format("2006年01月02日15:04:01"),
			Avatar:       info.Avatar,
			Introduction: info.Introduction,
			Status:       user.Status,
			CreatorId:    info.CreateId,
			GroupName:    group.GroupName,
			RoleName:     role.Name,
		}
		out = append(out, outTemp)
	}
	return &dto.GroupUserListOutPut{
		Total:    total,
		List:     out,
		PageNo:   info.PageNo,
		PageSize: info.PageSize,
	}, nil
}

func (u *userService) UpdateUserInfo(ctx *gin.Context, info *dto.UpdateUserInfo) error {
	tx, err := lib.GetGormPool("default")
	if err != nil {
		return err
	}
	adminDB := &dao.Admin{Id: info.ID}
	admin, err := adminDB.Find(ctx, tx, adminDB)
	if err != nil {
		return err
	}
	admin.UserName = info.Name
	admin.GroupId = info.GroupID
	//开启事务
	tx = tx.Begin()
	if err := admin.Updates(ctx, tx); err != nil {
		tx.Rollback()
		return err
	}
	userinfoDB := &roledao.UserInfo{Id: admin.InfoId, Introduction: info.Introduction}
	//获取操作用户名
	//从ctx中取出当前操作用户的uid
	claims, exists := ctx.Get("claims")
	if !exists {
		return errors.New("claims不存在,请检查jwt中间件")
	}
	cla, _ := claims.(*public.CustomClaims)
	tempAdminDB := &dao.Admin{Id: cla.Uid}
	tempAdmin, err := tempAdminDB.Find(ctx, tx, tempAdminDB)
	if err != nil {
		userinfoDB.CreateId = "unknown"
	} else {
		userinfoDB.CreateId = tempAdmin.UserName
	}
	if err := userinfoDB.Updates(ctx, tx); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// DeleteUser 删除用户
func (u *userService) DeleteUser(ctx *gin.Context, params *dto.UserIDInput) error {
	tx, err := lib.GetGormPool("default")
	if err != nil {
		return err
	}
	userDB := &dao.Admin{Id: params.ID}
	user, err := userDB.Find(ctx, tx, userDB)
	if err != nil {
		return err
	}
	if user.UserName == "admin" {
		return errors.New("默认admin用户不能删除")
	}
	user.IsDelete = 1
	return user.Updates(ctx, tx)
}

// ResetUserPassword 重置用户密码为  admin@123
func (u *userService) ResetUserPassword(ctx *gin.Context, params *dto.UserIDInput) error {
	tx, err := lib.GetGormPool("default")
	if err != nil {
		return err
	}
	userDB := &dao.Admin{Id: params.ID}
	user, err := userDB.Find(ctx, tx, userDB)
	newHashPassword := public.GenSaltPassword(user.Salt, "admin@123")
	user.Password = newHashPassword
	return user.Updates(ctx, tx)
}
