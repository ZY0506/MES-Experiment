<template>
  <div class="card">
    <div class="card-header">
      <h3>编辑商品</h3>
    </div>
    <div class="card-body">
      <el-form :model="form" :rules="rules" ref="formRef" label-width="110px" status-icon>
        <el-form-item label="名称" prop="name">
          <el-input v-model="form.name" />
        </el-form-item>
        <el-form-item label="分类" prop="category">
          <el-input v-model="form.category" />
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input v-model="form.description" />
        </el-form-item>
        <el-form-item label="价格" prop="price">
          <el-input-number v-model="form.price" :min="0" />
        </el-form-item>
        <el-form-item label="数量" prop="quantity">
          <el-input-number v-model="form.quantity" :min="0" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSubmit" :loading="saving">保存</el-button>
          <el-button @click="onCancel">取消</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import { reactive, ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useHardwareStore } from '@/store/hardwareStore'

const route = useRoute()
const router = useRouter()
const store = useHardwareStore()

const formRef = ref(null)
const saving = ref(false)

const form = reactive({
  id: null,
  name: '',
  category: '',
  description: '',
  price: 0,
  quantity: 0
})

const rules = {
  name: [{ required: true, message: '请输入名称', trigger: 'blur' }],
  category: [{ required: true, message: '请输入分类', trigger: 'blur' }],
  price: [{ required: true, message: '请输入价格', trigger: 'change' }],
  quantity: [{ required: true, message: '请输入数量', trigger: 'change' }]
}

// -- START: 修改部分 --
onMounted(async () => {
  const id = route.params.id
  
  // 首先，尝试从已经存在的列表中查找商品信息。
  // 如果用户是从列表页点击过来的，通常可以在这里找到，避免一次多余的API请求。
  let item = store.items.find(i => String(i.ID) === String(id))

  // 如果在当前列表中找不到（例如用户直接刷新了编辑页面），
  // 则调用一个专门根据ID获取单个商品信息的action。
  // 注意：你需要在你的 hardwareStore.js 中添加一个类似 fetchItemById 的 action，
  // 这个 action 应该接收一个 id，然后向后端API请求这个商品的详细数据。
  if (!item) {
    // 假设你的store中有一个名为 fetchItemById 的方法
    // 这比原来的 fetchList() 方法更准确、更高效
    item = await store.fetchItemById(id)
  }

  if (item) {
    form.id = item.ID
    form.name = item.name
    form.category = item.category
    form.description = item.description
    form.price = item.price
    form.quantity = item.quantity
  } else {
    // 如果商品不存在，可以给出提示并返回列表页
    console.error('商品未找到!');
    router.push('/list');
  }
})
// -- END: 修改部分 --

async function onSubmit() {
  formRef.value.validate(async valid => {
    if (!valid) return
    saving.value = true
    const payload = {
      id: form.id,
      name: form.name,
      category: form.category,
      description: form.description,
      price: form.price,
      quantity: form.quantity
    }
    const ok = await store.editItem(payload)
    saving.value = false
    if (ok) {
      router.push('/list')
    }
  })
}

function onCancel() {
  router.push('/list')
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