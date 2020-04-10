// Code generated by dol build. DO NOT EDIT.
// source: auto.go
const axios = require('@/utils/request')

// add 添加域
module.exports.add = (data) => {
  const url = '/api/sys/domain/add'
  return axios({
    url: url,
    method: 'post',
    data
  })
}

// del 删除域
module.exports.del = (data) => {
  const url = '/api/sys/domain/del'
  return axios({
    url: url,
    method: 'delete',
    data
  })
}

// update 更新域
module.exports.update = (data) => {
  const url = '/api/sys/domain/update'
  return axios({
    url: url,
    method: 'put',
    data
  })
}

// page 域分页查询
module.exports.page = (data) => {
  let url = '/api/sys/domain/page?'
  for (var key in data) {
    url += key + '=' + data[key] + '&'
  }
  return axios({
    url: url,
    method: 'get'
  })
}

// get 获取域信息
module.exports.get = (data) => {
  let url = '/api/sys/domain/get?'
  for (var key in data) {
    url += key + '=' + data[key] + '&'
  }
  return axios({
    url: url,
    method: 'get'
  })
}

