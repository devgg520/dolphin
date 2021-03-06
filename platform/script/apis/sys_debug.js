// Code generated by dol build. DO NOT EDIT.
const axios = require('../request').default

// pprof 
module.exports.pprof = (data = {}, opt = {}) => {
  let url = opt.url || '/debug/pprof/?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

// heap 
module.exports.heap = (data = {}, opt = {}) => {
  let url = opt.url || '/debug/pprof/heap?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

// goroutine 
module.exports.goroutine = (data = {}, opt = {}) => {
  let url = opt.url || '/debug/pprof/goroutine?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

// allocs 
module.exports.allocs = (data = {}, opt = {}) => {
  let url = opt.url || '/debug/pprof/allocs?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

// block 
module.exports.block = (data = {}, opt = {}) => {
  let url = opt.url || '/debug/pprof/block?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

// threadcreate 
module.exports.threadcreate = (data = {}, opt = {}) => {
  let url = opt.url || '/debug/pprof/threadcreate?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

// cmdline 
module.exports.cmdline = (data = {}, opt = {}) => {
  let url = opt.url || '/debug/pprof/cmdline?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

// profile 
module.exports.profile = (data = {}, opt = {}) => {
  let url = opt.url || '/debug/pprof/profile?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

// symbol 
module.exports.symbol = (data = {}, opt = {}) => {
  const url = opt.url ||  '/debug/pprof/symbol'
  if ((opt.url || 'get,post') === 'get') {
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
    method: 'get,post',
    data,
    ...opt
  })
}

// trace 
module.exports.trace = (data = {}, opt = {}) => {
  let url = opt.url || '/debug/pprof/trace?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

// mutex 
module.exports.mutex = (data = {}, opt = {}) => {
  let url = opt.url || '/debug/pprof/mutex?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

