import request from '@/utils/request'

// 用于获取下载文件的url，并且检测文件是否在本地存在
export function checkDownloadBakFile(id) {
  return request({
    url: '/public/check_file_exists',
    method: 'get',
    params: { id }
  })
}

// 用于程序异常终止后。调用函数可批量启动运行中任务
export function initBakTask() {
  return request({
    url: '/public/initbak',
    method: 'get'
  })
}

