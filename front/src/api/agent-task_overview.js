import request from '@/utils/request'

export function GetAgentTaskOverViewList (query) {
  return request({
    url: '/agent/overview_task_list',
    method: 'get',
    params: query
  })
}
