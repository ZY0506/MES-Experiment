# Vue 硬件管理 (CRUD)

- Vue 3 + Vite + Element Plus + Pinia + Axios
- 使用后端接口（来自你提供的 Apifox 文档）
  - Base URL: http://127.0.0.1:8081
  - 添加: POST /api/hardware/create
  - 查询: GET /api/hardware/query?name=xxx
  - 修改: PUT /api/hardware/update
  - 删除: DELETE /api/hardware/delete?id=xxx

## 快速开始

1. 解压并进入项目目录：
   ```bash
   cd vue-hardware-crud
   ```

2. 安装依赖：
   ```bash
   npm install
   ```

3. 启动：
   ```bash
   npm run dev
   ```

4. 打开浏览器访问 `http://localhost:5173`（Vite 默认端口）

## 注意事项

- 查询接口 `/api/hardware/query` 在 Apifox 文档中要求 `name` 参数。前端允许输入空字符串以拉取全部或匹配项，具体行为取决于后端实现。
- 全局 axios 实例已配置基础地址为 `http://127.0.0.1:8081`，如需修改，请编辑 `src/permission.js`
- 已实现统一的错误提示（使用 Element Plus 的 `ElMessage`）和基础表单校验。
