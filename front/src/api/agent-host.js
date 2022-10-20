import request from '@/utils/request'
import { Encrypt } from '@/utils/security'

export function GetAgentHostList (query) {
  return request({
    url: '/agent/hostlist',
    method: 'get',
    params: query
  })
}

export function TestAgentHost (query) {
  return request({
    url: '/agent/host_test',
    method: 'get',
    params: query
  })
}

export function GetHostNames (query) {
  return request({
    url: '/agent/host_names',
    method: 'get',
    params: query
  })
}

export function CreateAgentHost (data) {
  data.password = Encrypt(data.password)
  return request({
    url: '/agent/hostadd',
    method: 'post',
    data
  })
}

export function UpdateAgentHost (query) {
  return request({
    url: '/agent/hostupdate',
    method: 'put',
    params: query
  })
}

export function DeleteAgentHost (query) {
  return request({
    url: '/agent/hostdelete',
    method: 'delete',
    params: query
  })
}
