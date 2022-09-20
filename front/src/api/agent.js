import request from '@/utils/request'

export function GetServiceList() {
  return request({
    url: '/public/agentlist',
    method: 'get'
  })
}

