import axios from 'axios'
import { ElMessage } from 'element-plus'

const apiClient = axios.create({
  baseURL: '/',  
  timeout: 8000,
  headers: {
    'Content-Type': 'application/json'
  }
})

apiClient.interceptors.request.use(config => {
  return config
}, error => {
  return Promise.reject(error)
})

apiClient.interceptors.response.use(response => {
  return response
}, error => {
  let msg = '网络错误'
  if (error.response && error.response.data && error.response.data.msg) {
    msg = error.response.data.msg
  } else if (error.message) {
    msg = error.message
  }
  ElMessage.error(msg)
  return Promise.reject(error)
})

export default apiClient
