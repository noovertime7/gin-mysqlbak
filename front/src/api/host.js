import request from '@/utils/request'

export function hostAdd (data) {
  return request({
    url: '/host/hostadd',
    method: 'post',
    data
  })
}

export function hostDelete (id) {
  return request({
    url: '/host/hostdelete',
    method: 'delete',
    params: id
  })
}

export function hostUpdate (data) {
  return request({
    url: '/host/hostupdate',
    method: 'post',
    data
  })
}

export function hostList (data) {
  return request({
    url: '/host/hostlist',
    method: 'get',
    params: data
  })
}
