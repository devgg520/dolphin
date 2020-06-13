// Code generated by dol build. DO NOT EDIT.
const axios = require('@/utils/request').default

// add 添加字典
module.exports.add = (data) => {
  const url = '/api/sys/optionset/add'
  return axios({
    url: url,
    method: 'post',
    data
  })
}

// del 删除字典
module.exports.del = (data) => {
  const url = '/api/sys/optionset/del'
  return axios({
    url: url,
    method: 'delete',
    data
  })
}

// update 更新字典
module.exports.update = (data) => {
  const url = '/api/sys/optionset/update'
  return axios({
    url: url,
    method: 'put',
    data
  })
}

// page 字典分页查询
module.exports.page = (data) => {
  let url = '/api/sys/optionset/page?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get'
  })
}

// get 获取字典信息
module.exports.get = (data) => {
  let url = '/api/sys/optionset/get?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get'
  })
}
