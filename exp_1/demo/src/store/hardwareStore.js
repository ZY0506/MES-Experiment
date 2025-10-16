import { defineStore } from 'pinia'
import { ref } from 'vue'
// --- START: 修改部分 1 ---
// 假设你有一个名为 findHardware 的 API 函数，用于根据 ID 获取单个商品
// 你需要从你的 API 文件中导入它
import { createHardware, queryHardware, updateHardware, deleteHardware, findHardware } from '@/api/hardware'
// --- END: 修改部分 1 ---
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

  // --- START: 新增部分 ---
  // 根据 ID 获取单个商品的 Action
  async function fetchItemById(id) {
    loading.value = true
    try {
      // 调用 API 函数，传入 ID
      const resp = await findHardware(id)
      if (resp && resp.data && resp.data.data) {
        // 成功获取数据后，直接返回该商品的数据
        return resp.data.data
      }
      return null // 如果没有找到，返回 null
    } catch (err) {
      console.error('fetchItemById error', err)
      return null
    } finally {
      loading.value = false
    }
  }
  // --- END: 新增部分 ---

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

  // --- START: 修改部分 2 ---
  // 将新增的 fetchItemById 函数导出，以便在组件中可以调用
  return { items, total, loading, fetchList, fetchItemById, addItem, editItem, removeItem }
  // --- END: 修改部分 2 ---
})