// Code generated by dol build. DO NOT EDIT.
// source: auto.go

const axios = require('../axios')


// Page 日志分页查询
module.exports.Page = (data) => {
	const url = '/api/sys/tracker/page'
	return axios({
		url: url,
		method: 'get',
		data
    })
}

// Get 获取日志信息
module.exports.Get = (data) => {
	const url = '/api/sys/tracker/get'
	return axios({
		url: url,
		method: 'get',
		data
    })
}
