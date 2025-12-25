# 修复依赖问题

如果遇到 "missing go.sum entry" 错误，请按以下步骤操作：

## 方法一：自动修复（推荐）

在 `backend` 目录下运行：

```bash
go mod download
go mod tidy
```

这会自动下载所有依赖并更新 `go.sum` 文件。

## 方法二：手动下载特定依赖

如果只缺少某个特定包：

```bash
go mod download github.com/gorilla/mux
```

## 方法三：清理并重新下载

如果上述方法都不行，可以清理缓存后重新下载：

```bash
# 清理模块缓存
go clean -modcache

# 重新下载依赖
go mod download
go mod tidy
```

## 验证

运行以下命令验证依赖是否正确：

```bash
go mod verify
```

## 如果仍有问题

1. 检查网络连接
2. 检查 Go 版本（需要 >= 1.21）
3. 检查 `go.mod` 文件是否正确
4. 尝试设置 Go 代理（如果在中国）：

```bash
go env -w GOPROXY=https://goproxy.cn,direct
go mod download
```

## 完成后

依赖修复后，可以运行测试服务器：

```bash
./quick_test.sh
```

或：

```bash
go run ./cmd/test/main.go
```

