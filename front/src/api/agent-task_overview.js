import request from '@/utils/request'

export function GetAgentTaskOverViewList (query) {
  return request({
    url: '/agent/overview_task_list',
    method: 'get',
    params: query
  })
}

export function stopOverviewBak (query) {
  return request({
    url: '/agent/overview_task_stop',
    method: 'get',
    params: query
  })
}

export function startOverviewBak (query) {
  return request({
    url: '/agent/overview_task_start',
    method: 'get',
    params: query
  })
}

export function batchStartOverviewBak (data) {
  return request({
    url: '/agent/overview_task_batch_start',
    method: 'post',
    data
  })
}

export function batchStopOverviewBak (data) {
  return request({
    url: '/agent/overview_task_batch_stop',
    method: 'post',
    data
  })
}

export function deleteOverviewBak (query) {
  return request({
    url: '/agent/overview_task_delete',
    method: 'delete',
    params: query
  })
}

export function restoreOverviewBak (query) {
  return request({
    url: '/agent/overview_task_restore',
    method: 'put',
    params: query
  })
}

export function syncOverviewBak () {
  return request({
    url: '/agent/overview_task_sync',
    method: 'get'
  })
}
