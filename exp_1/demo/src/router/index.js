import { createRouter, createWebHistory } from 'vue-router'
import ListProduct from '@/views/ListProduct.vue'
import AddProduct from '@/views/AddProduct.vue'
import EditProduct from '@/views/EditProduct.vue'

const routes = [
  { path: '/', redirect: '/list' },
  { path: '/list', component: ListProduct },
  { path: '/add', component: AddProduct },
  { path: '/edit/:id', component: EditProduct, props: true }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
