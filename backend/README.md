# FreshKeep Backend

保鲜记后端服务，基于Kratos框架开发。

## 前置要求

- Go 1.21 或更高版本
- protoc 编译器（用于生成 Protobuf 代码）
- Wire 工具（用于依赖注入）

### 快速检查环境

```bash
./check_go.sh
```

如果 Go 未安装，请查看 [INSTALL_GO.md](INSTALL_GO.md) 了解安装方法。

## 技术栈

- Go 1.21+
- Kratos v2.7.2
- Protobuf
- GORM
- PostgreSQL/SQLite

## 项目结构

```
backend/
├── api/                    # API定义（Protobuf）
│   └── freshkeep/
│       └── v1/
├── cmd/                    # 应用入口
│   ├── server/            # 主服务器
│   └── test/              # 测试服务器（无需生成代码）
├── internal/               # 内部代码
│   ├── biz/               # 业务逻辑层
│   ├── service/           # 服务层
│   ├── data/             # 数据访问层
│   └── conf/             # 配置
├── configs/               # 配置文件
└── pkg/                   # 公共包
    └── database/
```

## 快速开始

### 1. 快速测试（无需生成代码）

如果你想快速测试接口是否可访问：

```bash
# 方式一：使用脚本
./quick_test.sh

# 方式二：直接运行
go run ./cmd/test/main.go
```

服务器会在 `http://localhost:8000` 启动。

测试接口：
```bash
curl http://localhost:8000/health
curl http://localhost:8000/api/v1/items
```

### 2. 完整开发流程

#### 安装工具

```bash
make init
```

这会安装：
- protoc-gen-go
- protoc-gen-go-grpc
- kratos CLI
- protoc-gen-go-http
- wire

#### 生成代码

```bash
# 生成Protobuf代码
make api

# 生成Wire依赖注入代码
make wire
```

#### 运行服务器

```bash
# 方式一：使用 make
make run

# 方式二：直接运行
go run ./cmd/server/main.go ./cmd/server/wire_gen.go -conf configs
```

## 开发

### 生成代码

```bash
# 生成Protobuf代码
make api

# 生成Wire依赖注入代码
make wire
```

### 运行

```bash
go run cmd/server/main.go cmd/server/wire_gen.go -conf configs
```

### 测试

```bash
# 运行测试
make test

# 或
go test -v ./...
```

## API文档

API使用Protobuf定义，支持HTTP和gRPC两种协议。

### 已定义的API

- Item Service: 物品管理
- Category Service: 分类管理
- Statistics Service: 统计信息
- Barcode Service: 条形码扫描
- Image Service: 图片识别

详细API定义请查看 `api/freshkeep/v1/` 目录下的 `.proto` 文件。

## 测试

查看 [README_TEST.md](README_TEST.md) 了解详细的测试说明。

## 常见问题

### Go 未安装

运行 `./check_go.sh` 检查，如果未安装，查看 [INSTALL_GO.md](INSTALL_GO.md)。

### Wire 生成失败

确保所有依赖都正确注入，检查 ProviderSet 配置。

### 数据库连接失败

- 确保 SQLite 数据库目录存在（`backend/data/`）
- 检查配置文件路径
- 确认数据库驱动已安装
