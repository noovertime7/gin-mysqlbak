import request from '@/utils/request'

export function taskList (data) {
  return request({
    url: '/task/tasklist',
    method: 'get',
    params: data
  })
}

export function taskDelete (query) {
  return request({
    url: '/task/taskdelete',
    method: 'delete',
    params: query
  })
}

export function taskAdd (data) {
  return request({
    url: '/task/taskadd',
    method: 'post',
    data
  })
}

export function taskUpdate (data) {
  return request({
    url: '/task/taskupdate',
    method: 'put',
    data
  })
}

export function taskDetail (query) {
  return request({
    url: '/task/taskdetail',
    method: 'get',
    params: query
  })
}
