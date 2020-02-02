// Code generated by dol build. DO NOT EDIT.
// source: auto.go

const axios = require('../axios')


// Add 添加角色
module.exports.Add = (data) => {
	const url = '/api/sys/role/add'
	return axios({
		url: url,
		method: 'post',
		data
    })
}

// Del 删除角色
module.exports.Del = (data) => {
	const url = '/api/sys/role/del'
	return axios({
		url: url,
		method: 'delete',
		data
    })
}

// Update 更新角色
module.exports.Update = (data) => {
	const url = '/api/sys/role/update'
	return axios({
		url: url,
		method: 'put',
		data
    })
}

// Page 角色分页查询
module.exports.Page = (data) => {
	const url = '/api/sys/role/page'
	return axios({
		url: url,
		method: 'get',
		data
    })
}

// Get 获取角色信息
module.exports.Get = (data) => {
	const url = '/api/sys/role/get'
	return axios({
		url: url,
		method: 'get',
		data
    })
}
