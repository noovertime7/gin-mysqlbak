import request from '@/utils/request'

export function GetServiceList () {
  return request({
    url: '/public/agentlist',
    method: 'get'
  })
}

export function GetServiceNumInfo () {
  return request({
    url: '/public/service_num_info',
    method: 'get'
  })
}
