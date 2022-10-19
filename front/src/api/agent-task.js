import request from '@/utils/request'

export function GetAgentTaskList (query) {
  return request({
    url: '/agent/tasklist',
    method: 'get',
    params: query
  })
}

export function GetAgentTaskDetail (query) {
  return request({
    url: '/agent/taskdetail',
    method: 'get',
    params: query
  })
}

export function DeleteAgentTask (query) {
  return request({
    url: '/agent/taskdelete',
    method: 'delete',
    params: query
  })
}

export function AddAgentTask (data) {
  return request({
    url: '/agent/taskadd',
    method: 'post',
    data
  })
}

export function AutoAddAgentTask (data) {
  return request({
    url: '/agent/task_auto_add',
    method: 'post',
    data
  })
}

export function UpdateAgentTask (query) {
  return request({
    url: '/agent/taskupdate',
    method: 'put',
    params: query
  })
}

export function StartAgentTask (query) {
  return request({
    url: '/agent/bakstart',
    method: 'put',
    params: query
  })
}

export function StopAgentTask (query) {
  return request({
    url: '/agent/bakstop',
    method: 'put',
    params: query
  })
}

export function StartAgentHostTask (query) {
  return request({
    url: '/agent/bakhoststart',
    method: 'put',
    params: query
  })
}

export function StopAgentHostTask (query) {
  return request({
    url: '/agent/bakhoststop',
    method: 'put',
    params: query
  })
}
