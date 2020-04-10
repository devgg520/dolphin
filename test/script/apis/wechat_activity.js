// Code generated by dol build. DO NOT EDIT.
// source: auto.go
const axios = require('@/utils/request')

// BatchAdd 添加活动
module.exports.BatchAdd = (data) => {
  const url = '/api/wechat/activity/batch_add'
  return axios({
    url: url,
    method: 'post',
    data
  })
}

// Add 添加活动
module.exports.Add = (data) => {
  const url = '/api/wechat/activity/add'
  return axios({
    url: url,
    method: 'post',
    data
  })
}

// BatchDel 删除活动
module.exports.BatchDel = (data) => {
  const url = '/api/wechat/activity/batch_del'
  return axios({
    url: url,
    method: 'delete',
    data
  })
}

// Del 删除活动
module.exports.Del = (data) => {
  const url = '/api/wechat/activity/del'
  return axios({
    url: url,
    method: 'delete',
    data
  })
}

// BatchUpdate 更新活动
module.exports.BatchUpdate = (data) => {
  const url = '/api/wechat/activity/batch_update'
  return axios({
    url: url,
    method: 'put',
    data
  })
}

// Update 更新活动
module.exports.Update = (data) => {
  const url = '/api/wechat/activity/update'
  return axios({
    url: url,
    method: 'put',
    data
  })
}

// List 活动分页查询
module.exports.List = (data) => {
  let url = '/api/wechat/activity/list?'
  for (var key in data) {
    url += key + '=' + data[key] + '&'
  }
  return axios({
    url: url,
    method: 'get',
  })
}

// One 获取活动
module.exports.One = (data) => {
  let url = '/api/wechat/activity/one?'
  for (var key in data) {
    url += key + '=' + data[key] + '&'
  }
  return axios({
    url: url,
    method: 'get',
  })
}

// Increase 增加次数
module.exports.Increase = (data) => {
  const url = '/api/v1/wechat/activity/increase'
  return axios({
    url: url,
    method: 'post',
    data
  })
}

// IncreaseV2 增加次数
module.exports.IncreaseV2 = (data) => {
  const url = '/api/v2/wechat/activity/increase_v2'
  return axios({
    url: url,
    method: 'post',
    data
  })
}

