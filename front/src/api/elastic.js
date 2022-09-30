import request from '@/utils/request'

export function getEsTaskList (query) {
  return request({
    url: 'agent/es/tasklist',
    method: 'get',
    params: query
  })
}

export function addEsTask (data) {
  return request({
    url: 'agent/es/taskadd',
    method: 'post',
    data
  })
}

export function deleteEsTask (query) {
  return request({
    url: 'agent/es/taskdelete',
    method: 'delete',
    params: query
  })
}

export function updateEsTask (data) {
  return request({
    url: 'agent/es/taskupdate',
    method: 'put',
    data
  })
}

export function getEsTaskDetail (query) {
  return request({
    url: 'agent/es/taskdetail',
    method: 'get',
    params: query
  })
}

export function startEsTask (query) {
  return request({
    url: 'agent/es/start',
    method: 'put',
    params: query
  })
}

export function stopEsTask (query) {
  return request({
    url: 'agent/es/stop',
    method: 'put',
    params: query
  })
}

export function getEsHistoryList (query) {
  return request({
    url: 'agent/es/historylist',
    method: 'get',
    params: query
  })
}

export function deleteEsHistory (query) {
  return request({
    url: 'agent/es/historydelete',
    method: 'delete',
    params: query
  })
}
