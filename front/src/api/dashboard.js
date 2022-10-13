import request from '@/utils/request'

export function getSvcTNum () {
  return request({
    url: 'dashboard/service_task_num',
    method: 'get'
  })
}

export function clusterDataByDate (query) {
  return request({
    url: '/dashboard/service_info_by_date',
    method: 'get',
    params: query
  })
}

export function getSvcFinishNum () {
  return request({
    url: 'dashboard/service_finish_num',
    method: 'get'
  })
}
