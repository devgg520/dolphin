// Code generated by dol build. DO NOT EDIT.
// source: auto.go

const axios = require('../axios')


// Add 添加区域
module.exports.Add = (data) => {
	const url = '/api/sys/area/add'
	return axios({
		url: url,
		method: 'post',
		data
    })
}

// Del 删除区域
module.exports.Del = (data) => {
	const url = '/api/sys/area/del'
	return axios({
		url: url,
		method: 'delete',
		data
    })
}

// Update 更新区域
module.exports.Update = (data) => {
	const url = '/api/sys/area/update'
	return axios({
		url: url,
		method: 'put',
		data
    })
}

// Page 区域分页查询
module.exports.Page = (data) => {
	const url = '/api/sys/area/page'
	return axios({
		url: url,
		method: 'get',
		data
    })
}

// Get 获取区域信息
module.exports.Get = (data) => {
	const url = '/api/sys/area/get'
	return axios({
		url: url,
		method: 'get',
		data
    })
}

