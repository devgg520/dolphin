// Code generated by dol build. DO NOT EDIT.
// source: auto.go
const axios = require('axios')

const request = axios.create({
  baseURL: '/',
  timeout: 6000,
  headers: { 'X-Custom-Header': 'xxx' }
});

module.exports = request