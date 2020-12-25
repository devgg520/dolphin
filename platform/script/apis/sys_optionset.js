// Code generated by dol build. DO NOT EDIT.
const axios = require('../request').default

// add 添加字典
module.exports.add = (data = {}, opt = {}) => {
  const url = opt.url ||  '/api/sys/optionset/add'
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

// batchAdd 添加字典
module.exports.batchAdd = (data = {}, opt = {}) => {
  const url = opt.url ||  '/api/sys/optionset/batch_add'
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

// del 删除字典
module.exports.del = (data = {}, opt = {}) => {
  const url = opt.url ||  '/api/sys/optionset/del'
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

// batchDel 删除字典
module.exports.batchDel = (data = {}, opt = {}) => {
  const url = opt.url ||  '/api/sys/optionset/batch_del'
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

// update 更新字典
module.exports.update = (data = {}, opt = {}) => {
  const url = opt.url ||  '/api/sys/optionset/update'
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

// batchUpdate 更新字典
module.exports.batchUpdate = (data = {}, opt = {}) => {
  const url = opt.url ||  '/api/sys/optionset/batch_update'
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

// page 字典分页查询
module.exports.page = (data = {}, opt = {}) => {
  let url = opt.url || '/api/sys/optionset/page?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

// get 获取字典信息
module.exports.get = (data = {}, opt = {}) => {
  let url = opt.url || '/api/sys/optionset/get?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

