<template>
  <el-container style="height: 100vh">
    <el-aside width="220px" class="app-aside">
      <div class="logo">硬件管理系统</div>
      <el-menu
        :default-active="active"
        router
        class="el-menu-vertical-demo"
        background-color="#ffffff"
        text-color="#333"
        active-text-color="#409EFF"
      >
        <el-menu-item index="/list">查询商品</el-menu-item>
        <el-menu-item index="/add">添加商品</el-menu-item>
      </el-menu>
    </el-aside>
    <el-container>
      <!-- <el-header class="app-header">
        <div class="header-left"></div>
      </el-header> -->
      <el-main class="app-main">
        <router-view />
      </el-main>
      <el-footer class="app-footer">2023 工业软件</el-footer>
    </el-container>
  </el-container>
</template>

<script setup>
import { computed } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const active = computed(() => route.path)
</script>

<style scoped>
.app-aside {
  border-right: 1px solid #ebeef5;
  padding: 16px 8px;
  background: linear-gradient(180deg, #fafafa, #fff);
}
.logo {
  font-weight: 700;
  padding: 12px 8px;
  margin-bottom: 8px;
  text-align: center;
  color: #409EFF;
  border-bottom: 1px solid #f0f0f0;
}
.app-header {
  background: #fff;
  border-bottom: 1px solid #ebeef5;
  padding: 10px 20px;
  font-size: 16px;
}
.app-main {
  padding: 20px;
  background: #f5f7fa;
}
.app-footer {
  text-align: center;
  padding: 10px;
  color: #888;
}
</style>
