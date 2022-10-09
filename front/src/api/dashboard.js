import request from '@/utils/request'

export function getSvcTNum () {
  return request({
    url: 'dashboard/service_task_num',
    method: 'get'
  })
}
