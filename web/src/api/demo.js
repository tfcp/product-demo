import request from '@/utils/request'

export function getList(params) {
  return request({
    url: '/demo/user-list',
    method: 'get',
    params,
    baseURL: process.env.VUE_APP_URL
  })
}

export function getCount(params) {
  return request({
    url: '/demo/user-count',
    method: 'get',
    params,
    baseURL: process.env.VUE_APP_URL
  })
}

export function getDetail(id) {
  return request({
    url: '/demo/user-detail?id=' + id,
    method: 'get',
    baseURL: process.env.VUE_APP_URL
  })
}

export function Delete(id) {
  return request({
    url: '/demo/user-delete?id=' + id,
    method: 'post',
    baseURL: process.env.VUE_APP_URL
  })
}

export function enable(id) {
  const params = {
    id: id,
    status: 1
  }
  return request({
    url: '/demo/user-change',
    method: 'post',
    params,
    baseURL: process.env.VUE_APP_URL
  })
}

export function disable(id) {
  const params = {
    id: id,
    status: 2
  }
  return request({
    url: '/demo/user-change',
    method: 'post',
    params,
    baseURL: process.env.VUE_APP_URL
  })
}

export function save(params) {
  return request({
    url: '/demo/user-save',
    method: 'post',
    params,
    baseURL: process.env.VUE_APP_URL
  })
}
