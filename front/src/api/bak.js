import request from '@/utils/request'

export function startBak (query) {
  return request({
    url: '/bak/start',
    method: 'get',
    params: query
  })
}

export function stopBak (query) {
  return request({
    url: '/bak/stop',
    method: 'get',
    params: query
  })
}

export function startAllBakByHost (query) {
  return request({
    url: '/bak/start_bak_all_byhost',
    method: 'get',
    params: query
  })
}

export function stopAllBakByHost (hostid) {
  return request({
    url: '/bak/stop_bak_all_byhost',
    method: 'get',
    params: hostid
  })
}

export function startAllBak (query) {
  return request({
    url: '/bak/start_bak_all',
    method: 'get',
    params: query
  })
}

export function stopAllBak (hostid) {
  return request({
    url: '/bak/stop_bak_all',
    method: 'get',
    params: hostid
  })
}

export function getLocalHistoryList (data) {
  console.log(data)
  return request({
    url: '/bak/historylist',
    method: 'get',
    params: data
  })
}

export function getBakHistoryList () {
  return request({
    url: '/bak/findallhistory',
    method: 'get'
  })
}

export function deleteLocalHistory (data) {
  console.log(data)
  return request({
    url: '/bak/history_delete',
    method: 'delete',
    params: data
  })
}

export function getHistoryNumInfo () {
  return request({
    url: '/bak/history_num_info',
    method: 'get'
  })
}
