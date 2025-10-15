import apiClient from '../permission'

export function createHardware(data) {
  return apiClient.post('/api/hardware/create', data)
}

export function queryHardware(params) {
  return apiClient.get('/api/hardware/query', { params })
}

export function updateHardware(data) {
  return apiClient.put('/api/hardware/update', data)
}

export function deleteHardware(id) {
  return apiClient.delete('/api/hardware/delete', { params: { id } })
}
