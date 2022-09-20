import request from '@/utils/request'

export function panelGroupData() {
  return request({
    url: 'dashboard/panel_group_data',
    method: 'get'
  })
}

export function pieChartData() {
  return request({
    url: 'dashboard/pie_chart_data',
    method: 'get'
  })
}

export function GetBarData() {
  return request({
    url: '/agent/barchart',
    method: 'get'
  })
}
