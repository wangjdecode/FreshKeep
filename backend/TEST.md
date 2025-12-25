# 后端测试说明

## 前置要求

1. 安装 Go 1.21 或更高版本
2. 安装 protoc 编译器
3. 安装 Wire 工具

## 测试步骤

### 1. 生成 Protobuf 代码

首先需要生成配置和API的Protobuf代码：

```bash
cd backend

# 安装必要的工具（如果还没安装）
make init

# 生成Protobuf代码
make api
```

### 2. 生成 Wire 依赖注入代码

```bash
make wire
```

这会在 `cmd/server/` 目录下生成 `wire_gen.go` 文件。

### 3. 下载依赖

```bash
go mod download
go mod tidy
```

### 4. 运行服务器

```bash
# 方式1: 使用 make
make run

# 方式2: 直接运行
go run ./cmd/server/main.go ./cmd/server/wire_gen.go -conf configs
```

### 5. 测试接口

服务器启动后，可以通过以下方式测试：

#### 健康检查接口

```bash
curl http://localhost:8000/health
```

预期响应：
```
OK
```

#### API接口测试

```bash
curl http://localhost:8000/api/v1/items
```

预期响应：
```json
{"message": "API endpoint - Protobuf handlers will be registered here"}
```

## 常见问题

### 1. Wire 生成失败

如果 Wire 生成失败，检查：
- 所有依赖是否正确注入
- ProviderSet 是否正确配置
- 函数签名是否匹配

### 2. 数据库连接失败

确保：
- SQLite 数据库目录存在（`backend/data/`）
- 配置文件路径正确
- 数据库驱动已安装

### 3. Protobuf 代码生成失败

确保：
- protoc 已安装
- 所有 proto 文件语法正确
- 必要的 protoc 插件已安装

## 下一步

生成 Protobuf 代码后，需要：
1. 实现 HTTP 路由处理器
2. 连接 Service 层到 HTTP 处理器
3. 实现完整的 API 端点

