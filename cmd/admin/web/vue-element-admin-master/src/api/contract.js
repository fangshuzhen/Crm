import request from '@/utils/request'

export function listContract(data) {
  return request({
    url: '/crm/contract/list',
    method: 'post',
    data
  })
}

export function uploadContractFile(data) {
  return request({
    url: 'http://localhost:8882/crm/contract/upload',
    method: 'post',
    data
  })
}

export function createContract(data) {
  return request({
    url: '/crm/contract/create',
    method: 'post',
    data
  })
}

export function updateContract(data) {
  return request({
    url: '/crm/contract/update',
    method: 'post',
    data
  })
}

export function deleteContract(data) {
  return request({
    url: '/crm/contract/delete',
    method: 'post',
    data
  })
}

export function addUploadFile(data) {
  return request({
    url: '/crm/contract/addFile',
    method: 'post',
    data
  })
}

export function deleteFile(data) {
  return request({
    url: '/crm/contract/deleteFile',
    method: 'post',
    data
  })
}
