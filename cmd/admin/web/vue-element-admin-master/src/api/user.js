import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/user/admin/login',
    method: 'post',
    data
  })
}

export function getInfo(data) {
  return request({
    url: '/user/admin/info',
    method: 'post',
    data
  })
}

export function logout() {
  return request({
    url: '/user/admin/logout',
    method: 'post'
  })
}

export function listUser(data) {
  return request({
    url: '/user/admin/list',
    method: 'post',
    data
  })
}

export function createUser(data) {
  return request({
    url: '/user/admin/create',
    method: 'post',
    data
  })
}

export function updateUser(data) {
  return request({
    url: '/user/admin/update',
    method: 'post',
    data
  })
}

export function deleteUser(data) {
  return request({
    url: '/user/admin/delete',
    method: 'post',
    data
  })
}





