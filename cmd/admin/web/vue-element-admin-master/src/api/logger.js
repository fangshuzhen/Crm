import request from '@/utils/request'

export function listLogger(data) {
    return request({
      url: 'http://localhost:8883/log/logger/list',
      method: 'post',
      data
    })
  }