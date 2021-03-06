// Code generated by dol build. DO NOT EDIT.
const axios = require('../request').default

// add 添加worker
module.exports.add = (data = {}, opt = {}) => {
  const url = opt.url ||  '/api/sys/worker/add'
  if ((opt.url || 'post') === 'get') {
    for (var key in data) {
      url += key + '=' + encodeURIComponent(data[key]) + '&'
    }
    return axios({
      url: url,
      method: 'get',
      ...opt
    })
  }
  return axios({
    url: url,
    method: 'post',
    data,
    ...opt
  })
}

// get 获取worker信息
module.exports.get = (data = {}, opt = {}) => {
  let url = opt.url || '/api/sys/worker/get?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

