import request from '@/utils/request'


export function listCustomer(data) {
    return request({
      url: '/crm/customer/list',
      method: 'post',
      data
    })
}

export function createCustomer(data) {
  return request({
    url: '/crm/customer/create',
    method: 'post',
    data
  })
}

export function updateCustomer(data) {
  return request({
    url: '/crm/customer/update',
    method: 'post',
    data
  })
}

export function deleteCustomer(data) {
  return request({
    url: '/crm/customer/delete',
    method: 'post',
    data
  })
}