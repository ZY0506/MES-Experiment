<template>
  <div class="card">
    <div class="card-header" style="display:flex;align-items:center;justify-content:space-between">
      <div>
        <h3 style="margin:0">商品列表</h3>
      </div>
      <div style="display:flex;gap:8px;align-items:center">
        <el-input v-model="filters.name" placeholder="按名称查询" clearable @clear="onSearch" @keyup.enter.native="onSearch" />
        <el-button type="primary" @click="onSearch" :loading="loading">查询</el-button>
        <el-button @click="onReset">重置</el-button>
      </div>
    </div>

    <div class="card-body">
      <el-table :data="items" v-loading="loading" stripe border style="width:100%">
        <el-table-column prop="ID" label="ID" width="90" align="center" />
        <el-table-column prop="name" label="名称" min-width="150" show-overflow-tooltip />
        <el-table-column prop="category" label="分类" min-width="120" show-overflow-tooltip />
        <el-table-column prop="description" label="描述" min-width="250" show-overflow-tooltip />
        <el-table-column prop="price" label="价格" width="110" align="right">
            <template #default="scope">
                {{ formatPrice(scope.row.price) }}
            </template>
        </el-table-column>
        <el-table-column prop="quantity" label="数量" width="110" align="center" />
        <el-table-column label="操作" width="180" align="center" fixed="right">
          <template #default="scope">
            <el-button type="primary" size="small" @click="onEdit(scope.row.ID)">编辑</el-button>
            <el-button type="danger" size="small" @click="onDelete(scope.row.ID)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <div style="margin-top:12px;display:flex;justify-content:flex-end;align-items:center">
        <el-pagination
          background
          :page-size="pageSize"
          :page-sizes="[5,10, 20, 50, 100]"
          :current-page="page"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="onPageChange"
          @size-change="onSizeChange"
        />
      </div>
    </div>
  </div>
</template>

<script setup>
import { reactive, computed, onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import { useHardwareStore } from '@/store/hardwareStore'
import { ElMessageBox } from 'element-plus'

const router = useRouter()
const store = useHardwareStore()

const filters = reactive({
  name: ''
})

const page = ref(1)
const pageSize = ref(10)

const loading = computed(() => store.loading)
const items = computed(() => store.items)
const total = computed(() => store.total)

onMounted(() => {
  fetchData()
})

async function fetchData() {
  const params = {
    name: filters.name || '',
    page: page.value,
    pageSize: pageSize.value
  }
  await store.fetchList(params)
}

function formatPrice(price) {
  if (price === null || price === undefined) return 'N/A';
  return `¥ ${parseFloat(price).toFixed(2)}`;
}

function onSearch() {
  page.value = 1
  fetchData()
}

function onReset() {
  filters.name = ''
  fetchData()
}

function onEdit(id) {
  router.push(`/edit/${id}`)
}

async function onDelete(id) {
  try {
    await ElMessageBox.confirm('确认删除该商品吗？', '删除确认', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
  } catch {
    return
  }
  const ok = await store.removeItem(id)
  if (ok) {
    if (items.value.length === 1 && page.value > 1) {
      page.value--
    }
    fetchData()
  }
}

function onPageChange(p) {
  page.value = p
  fetchData()
}

function onSizeChange(size) {
  pageSize.value = size
  page.value = 1
  fetchData()
}
</script>

<style scoped>
.card {
  background: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 6px rgba(60,66,87,0.04);
  overflow: hidden;
}
.card-header { padding: 16px; border-bottom: 1px solid #f0f0f0; }
.card-body { padding: 20px; }
</style>