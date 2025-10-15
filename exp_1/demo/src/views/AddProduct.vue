<template>
  <div class="card">
    <div class="card-header">
      <h3>添加商品</h3>
    </div>
    <div class="card-body">
      <el-form :model="form" :rules="rules" ref="formRef" label-width="110px" status-icon>
        <el-form-item label="名称" prop="name">
          <el-input v-model="form.name" placeholder="例如：Nvidia 5060" />
        </el-form-item>
        <el-form-item label="分类" prop="category">
          <el-input v-model="form.category" placeholder="例如：显卡" />
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input v-model="form.description" placeholder="可选描述" />
        </el-form-item>
        <el-form-item label="价格" prop="price">
          <el-input-number v-model="form.price" :min="0" :controls="true" />
        </el-form-item>
        <el-form-item label="数量" prop="quantity">
          <el-input-number v-model="form.quantity" :min="0" :controls="true" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSubmit" :loading="submitting">提交</el-button>
          <el-button @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import { reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import { useHardwareStore } from '@/store/hardwareStore'

const router = useRouter()
const store = useHardwareStore()

const formRef = ref(null)
const submitting = ref(false)

const form = reactive({
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

async function onSubmit() {
  formRef.value.validate(async valid => {
    if (!valid) return
    submitting.value = true
    const ok = await store.addItem(form)
    submitting.value = false
    if (ok) {
      router.push('/list')
    }
  })
}

function onReset() {
  form.name = ''
  form.category = ''
  form.description = ''
  form.price = 0
  form.quantity = 0
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
