// Code generated by dol build. DO NOT EDIT.
// source: auto.go
const axios = require('@/utils/request')

// add 添加附件
module.exports.add = (data) => {
  const url = '/api/sys/attachment/add'
  return axios({
    url: url,
    method: 'post',
    data
  })
}

// upload 上传附件
module.exports.upload = (data) => {
  const url = '/api/sys/attachment/upload'
  return axios({
    url: url,
    method: 'post',
    data
  })
}

// del 删除附件
module.exports.del = (data) => {
  const url = '/api/sys/attachment/del'
  return axios({
    url: url,
    method: 'delete',
    data
  })
}

// update 更新附件
module.exports.update = (data) => {
  const url = '/api/sys/attachment/update'
  return axios({
    url: url,
    method: 'put',
    data
  })
}

// page 附件分页查询
module.exports.page = (data) => {
  let url = '/api/sys/attachment/page?'
  for (var key in data) {
    url += key + '=' + data[key] + '&'
  }
  return axios({
    url: url,
    method: 'get'
  })
}

// get 获取附件信息
module.exports.get = (data) => {
  let url = '/api/sys/attachment/get?'
  for (var key in data) {
    url += key + '=' + data[key] + '&'
  }
  return axios({
    url: url,
    method: 'get'
  })
}

