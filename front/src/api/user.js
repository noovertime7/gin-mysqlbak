import request from '@/utils/request'

const api = {
  groupList: '/admin/user_group_list',
  getUserByGroup: '/admin/userinfo_by_group'
}

export default api

export function getgroupList () {
  return request({
    url: api.groupList,
    method: 'get'
  })
}

export function getUserByGroup (parameter) {
  return request({
    url: api.getUserByGroup,
    method: 'get',
    params: parameter
  })
}
