// Code generated by dol build. DO NOT EDIT.
const axios = require('../request').default

// page 日志分页查询
module.exports.page = (data = {}, opt = {}) => {
  let url = opt.url || '/api/sys/tracker/page?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

// get 获取日志信息
module.exports.get = (data = {}, opt = {}) => {
  let url = opt.url || '/api/sys/tracker/get?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

