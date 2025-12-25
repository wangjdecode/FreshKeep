# 后端接口测试指南

## 快速测试（推荐）

如果你想快速测试接口是否可访问，可以使用简化的测试服务器：

```bash
cd backend
./quick_test.sh
```

或者直接运行：

```bash
go run ./cmd/test/main.go
```

服务器会在 `http://localhost:8000` 启动。

### 测试接口

在另一个终端中运行：

```bash
# 健康检查
curl http://localhost:8000/health

# 物品列表接口
curl http://localhost:8000/api/v1/items

# 分类接口
curl http://localhost:8000/api/v1/categories

# 统计接口
curl http://localhost:8000/api/v1/statistics/overview
```

## 完整测试（需要生成代码）

如果你想测试完整的 Kratos 服务器：

### 1. 安装工具

```bash
cd backend
make init
```

### 2. 生成代码

```bash
# 生成 Protobuf 代码
make api

# 生成 Wire 依赖注入代码
make wire
```

### 3. 运行服务器

```bash
make run
```

或者：

```bash
go run ./cmd/server/main.go ./cmd/server/wire_gen.go -conf configs
```

### 4. 测试接口

```bash
# 健康检查
curl http://localhost:8000/health

# API 接口（目前返回占位符）
curl http://localhost:8000/api/v1/items
```

## 预期响应

### 健康检查接口

```bash
$ curl http://localhost:8000/health
OK
```

### API 接口（测试服务器）

```bash
$ curl http://localhost:8000/api/v1/items
{
  "message": "Items API endpoint",
  "status": "working",
  "time": "2025-01-XX..."
}
```

### API 接口（完整服务器）

```bash
$ curl http://localhost:8000/api/v1/items
{
  "message": "API endpoint - Protobuf handlers will be registered here"
}
```

## 下一步

生成 Protobuf 代码后，需要：
1. 实现 HTTP 路由处理器，连接 Service 层
2. 实现完整的 CRUD 操作
3. 添加错误处理和验证
4. 添加数据库操作

