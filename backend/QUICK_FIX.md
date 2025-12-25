# 快速修复模块问题

## 问题

Go 尝试从 GitHub 获取模块，但仓库还不存在。

## 一键修复

运行以下命令：

```bash
cd backend

# 1. 设置环境变量（告诉 Go 这是私有模块）
./setup_env.sh

# 2. 清理并重新下载
go clean -modcache
go mod download
go mod tidy
```

## 或者手动执行

```bash
cd backend

# 设置私有模块
go env -w GOPRIVATE=github.com/freshkeep/backend
go env -w GONOPROXY=github.com/freshkeep/backend
go env -w GONOSUMDB=github.com/freshkeep/backend

# 清理缓存
go clean -modcache

# 重新下载依赖
go mod download

# 整理依赖
go mod tidy
```

## 验证

运行：

```bash
go mod verify
```

如果看到 "all modules verified" 就成功了。

## 然后测试

```bash
./quick_test.sh
```

或：

```bash
go run ./cmd/test/main.go
```

## 如果还是不行

可以尝试移除 `replace` 指令，只依赖环境变量：

1. 编辑 `go.mod`，删除或注释掉 `replace` 行
2. 确保运行了 `setup_env.sh`
3. 重新运行 `go mod download`

