import request from '@/utils/request'

export function GetServiceList (query) {
  return request({
    url: '/public/agentlist',
    method: 'get',
    params: query
  })
}

export function GetServiceNumInfo () {
  return request({
    url: '/public/service_num_info',
    method: 'get'
  })
}
