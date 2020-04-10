// Code generated by dol build. DO NOT EDIT.
// source: auto.go
const axios = require('@/utils/request')

// add 添加客户端
module.exports.add = (data) => {
  const url = '/api/sys/client/add'
  return axios({
    url: url,
    method: 'post',
    data
  })
}

// del 删除客户端
module.exports.del = (data) => {
  const url = '/api/sys/client/del'
  return axios({
    url: url,
    method: 'delete',
    data
  })
}

// update 更新客户端
module.exports.update = (data) => {
  const url = '/api/sys/client/update'
  return axios({
    url: url,
    method: 'put',
    data
  })
}

// page 客户端分页查询
module.exports.page = (data) => {
  let url = '/api/sys/client/page?'
  for (var key in data) {
    url += key + '=' + data[key] + '&'
  }
  return axios({
    url: url,
    method: 'get'
  })
}

// get 获取客户端信息
module.exports.get = (data) => {
  let url = '/api/sys/client/get?'
  for (var key in data) {
    url += key + '=' + data[key] + '&'
  }
  return axios({
    url: url,
    method: 'get'
  })
}

