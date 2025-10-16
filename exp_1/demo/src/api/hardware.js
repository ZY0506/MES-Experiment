import apiClient from '../permission'

export function createHardware(data) {
  return apiClient.post('/api/hardware/create', data)
}

export function queryHardware(params) {
  return apiClient.get('/api/hardware/query', { params })
}

// --- START: 新增部分 ---
/**
 * 根据商品ID获取单个商品详情
 * 对应 hardwareStore.js 中的 findHardware API 调用
 * @param {number|string} id - 商品ID
 */
export function findHardware(id) {
  // 假设你的后端获取单个商品的接口是 /api/hardware/get，并且通过查询参数 id 传递
  return apiClient.get('/api/hardware/get', { params: { id } })
}
// --- END: 新增部分 ---

export function updateHardware(data) {
  return apiClient.put('/api/hardware/update', data)
}

export function deleteHardware(id) {
  return apiClient.delete('/api/hardware/delete', { params: { id } })
}