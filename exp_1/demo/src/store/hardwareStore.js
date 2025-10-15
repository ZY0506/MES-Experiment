import { defineStore } from 'pinia'
import { ref } from 'vue'
import { createHardware, queryHardware, updateHardware, deleteHardware } from '@/api/hardware'
import { ElMessage } from 'element-plus'

export const useHardwareStore = defineStore('hardware', () => {
  const items = ref([])
  const total = ref(0)
  const loading = ref(false)

  async function fetchList(params = {}) {
    loading.value = true
    try {
      const resp = await queryHardware(params)
      if (resp && resp.data && resp.data.data) {
        items.value = resp.data.data.list || []
        total.value = resp.data.data.total || items.value.length
      } else {
        items.value = []
        total.value = 0
      }
    } catch (err) {
      console.error('fetchList error', err)
    } finally {
      loading.value = false
    }
  }

  async function addItem(payload) {
    try {
      const res = await createHardware(payload)
      if (res && res.data && res.data.code === 1) {
        ElMessage.success(res.data.msg || '添加成功')
        return true
      } else {
        ElMessage.error((res && res.data && res.data.msg) || '添加失败')
        return false
      }
    } catch (err) {
      return false
    }
  }

  async function editItem(payload) {
    try {
      const res = await updateHardware(payload)
      if (res && res.status === 200) {
        ElMessage.success('更新成功')
        return true
      } else {
        ElMessage.error('更新失败')
        return false
      }
    } catch (err) {
      return false
    }
  }

  async function removeItem(id) {
    try {
      const res = await deleteHardware(id)
      if (res && res.data && res.data.code === 1) {
        ElMessage.success(res.data.msg || '删除成功')
        return true
      } else if (res && res.status === 200) {
        ElMessage.success('删除成功')
        return true
      } else {
        ElMessage.error('删除失败')
        return false
      }
    } catch (err) {
      return false
    }
  }

  return { items, total, loading, fetchList, addItem, editItem, removeItem }
})
