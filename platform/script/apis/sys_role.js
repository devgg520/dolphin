// Code generated by dol build. DO NOT EDIT.
const axios = require('../request').default

// add 添加角色
module.exports.add = (data = {}, opt = {}) => {
  const url = opt.url ||  '/api/sys/role/add'
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

// batchAdd 添加角色
module.exports.batchAdd = (data = {}, opt = {}) => {
  const url = opt.url ||  '/api/sys/role/batch_add'
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

// del 删除角色
module.exports.del = (data = {}, opt = {}) => {
  const url = opt.url ||  '/api/sys/role/del'
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

// batchDel 删除角色
module.exports.batchDel = (data = {}, opt = {}) => {
  const url = opt.url ||  '/api/sys/role/batch_del'
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

// update 更新角色
module.exports.update = (data = {}, opt = {}) => {
  const url = opt.url ||  '/api/sys/role/update'
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

// batchUpdate 更新角色
module.exports.batchUpdate = (data = {}, opt = {}) => {
  const url = opt.url ||  '/api/sys/role/batch_update'
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

// page 角色分页查询
module.exports.page = (data = {}, opt = {}) => {
  let url = opt.url || '/api/sys/role/page?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

// roleMenuTree 角色菜单树形结构
module.exports.roleMenuTree = (data = {}, opt = {}) => {
  let url = opt.url || '/api/sys/role/role_menu_tree?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

// roleAppFunTree 角色App功能树形结构
module.exports.roleAppFunTree = (data = {}, opt = {}) => {
  let url = opt.url || '/api/sys/role/role_app_fun_tree?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

// get 获取角色信息
module.exports.get = (data = {}, opt = {}) => {
  let url = opt.url || '/api/sys/role/get?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

