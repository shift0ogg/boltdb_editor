# golang实现BoltDB 数据库管理工具
![效果图](https://github.com/shift0ogg/boltdb_editor/blob/main/docs/jm1.PNG?raw=true)
## BoltDB 数据库管理工具的完整架构：
```
boltdb_editor/
├── backend/                    # Golang 后端
│   ├── go.mod
│   ├── main.go                 # 服务器入口
│   ├── handlers/
│   │   └── handlers.go         # HTTP 处理器
│   ├── models/
│   │   └── models.go           # 数据模型
│   └── services/
│       └── db_service.go       # 数据库服务层
└── frontend/                   # Vue3 前端
    ├── package.json
    ├── vite.config.js
    ├── tailwind.config.js
    ├── postcss.config.js
    ├── index.html
    └── src/
        ├── main.js
        ├── App.vue
        ├── style.css
        └── components/
            ├── DatabaseSelector.vue    # 数据库选择器
            ├── BucketManager.vue       # 桶管理器
            ├── BucketViewer.vue        # 桶查看器
            ├── CreateBucketModal.vue   # 创建桶模态框
            └── KeyValueModal.vue       # 键值对模态框
        
```
## 功能特性
数据库管理：打开任意 BoltDB 数据库文件
桶管理：创建、删除和查看桶（相当于表）
数据管理：在桶中添加、编辑、删除键值对
可视化界面：使用 Vue3 + Tailwind CSS 的现代单页应用

## 运行说明
### 1.后端启动：
``` 
cd backend
go mod tidy
go run main.go
``` 

### 2.前端启动：
``` 
cd frontend
npm install
npm run dev
```

前端将在 http://localhost:5173 启动（代理到后端 API）

## 架构说明
分层架构：清晰分离了处理器、服务和模型层
RESTful API：提供标准的 HTTP API 接口
ORM 支持：使用 Storm 简化 BoltDB 操作
响应式前端：Vue3 组合式 API + Tailwind CSS 样式



