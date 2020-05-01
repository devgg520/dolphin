// Code generated by dol build. DO NOT EDIT.
// source: auto.go
const axios = require('@/utils/request').default

// page 日志分页查询
module.exports.page = (data) => {
  let url = '/api/sys/tracker/page?'
  for (var key in data) {
    url += key + '=' + data[key] + '&'
  }
  return axios({
    url: url,
    method: 'get'
  })
}

// get 获取日志信息
module.exports.get = (data) => {
  let url = '/api/sys/tracker/get?'
  for (var key in data) {
    url += key + '=' + data[key] + '&'
  }
  return axios({
    url: url,
    method: 'get'
  })
}

