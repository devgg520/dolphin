// Code generated by dol build. DO NOT EDIT.
const axios = require('../request').default

// add 添加区域
module.exports.add = (data = {}, opt = {}) => {
  const url = opt.url ||  '/api/sys/area/add'
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

// batchAdd 添加区域
module.exports.batchAdd = (data = {}, opt = {}) => {
  const url = opt.url ||  '/api/sys/area/batch_add'
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

// del 删除区域
module.exports.del = (data = {}, opt = {}) => {
  const url = opt.url ||  '/api/sys/area/del'
  if ((opt.url || 'delete') === 'get') {
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
    method: 'delete',
    data,
    ...opt
  })
}

// batchDel 删除文章
module.exports.batchDel = (data = {}, opt = {}) => {
  const url = opt.url ||  '/api/sys/area/batch_del'
  if ((opt.url || 'delete') === 'get') {
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
    method: 'delete',
    data,
    ...opt
  })
}

// update 更新区域
module.exports.update = (data = {}, opt = {}) => {
  const url = opt.url ||  '/api/sys/area/update'
  if ((opt.url || 'put') === 'get') {
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
    method: 'put',
    data,
    ...opt
  })
}

// batchUpdate 更新文章
module.exports.batchUpdate = (data = {}, opt = {}) => {
  const url = opt.url ||  '/api/sys/area/batch_update'
  if ((opt.url || 'put') === 'get') {
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
    method: 'put',
    data,
    ...opt
  })
}

// page 区域分页查询
module.exports.page = (data = {}, opt = {}) => {
  let url = opt.url || '/api/sys/area/page?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

// get 获取区域信息
module.exports.get = (data = {}, opt = {}) => {
  let url = opt.url || '/api/sys/area/get?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

