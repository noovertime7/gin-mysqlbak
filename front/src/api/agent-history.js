import request from '@/utils/request'

export function GetAgentHistory (query) {
  return request({
    url: '/agent/historylist',
    method: 'get',
    params: query
  })
}

export function DeleteHistory (query) {
  return request({
    url: '/agent/historydelete',
    method: 'delete',
    params: query
  })
}

export function GetAgentNumInfo (query) {
  return request({
    url: 'agent/history_num_info',
    method: 'get',
    params: query
  })
}
