import request from '@/utils/request'

const api = {
  groupList: '/admin/user_group_list',
  getUserByGroup: '/admin/userinfo_by_group',
  UpdateUserInfo: '/admin/user_info_update',
  ChangePwd: '/admin/changepwd',
  DeleteUser: '/admin/user_delete',
  ResetPassword: '/admin/user_reset_pwd'
}

export default api

// 获取用户组列表
export function getgroupList () {
  return request({
    url: api.groupList,
    method: 'get'
  })
}

// 根据用户组获取用户列表
export function getUserByGroup (parameter) {
  return request({
    url: api.getUserByGroup,
    method: 'get',
    params: parameter
  })
}

// 更新用户信息
export function UpdateUserInfo (parameter) {
  return request({
    url: api.UpdateUserInfo,
    method: 'put',
    data: parameter
  })
}

// 更改用户密码
export function ChangePwd (parameter) {
  return request({
    url: api.ChangePwd,
    method: 'post',
    data: parameter
  })
}

// 删除用户
export function DeleteUser (parameter) {
  return request({
    url: api.DeleteUser,
    method: 'delete',
    params: parameter
  })
}

// 重置用户密码
export function ResetUserPassword (parameter) {
  return request({
    url: api.ResetPassword,
    method: 'get',
    params: parameter
  })
}
