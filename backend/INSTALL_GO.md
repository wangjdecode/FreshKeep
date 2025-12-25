# Go 安装指南

## macOS 安装方式

### 方式一：使用 Homebrew（推荐）

```bash
# 安装 Go
brew install go

# 验证安装
go version
```

### 方式二：官方安装包

1. 访问 https://go.dev/dl/
2. 下载 macOS 安装包（.pkg 文件）
3. 双击安装包进行安装
4. 安装完成后，重启终端

### 方式三：使用 Go 官方安装脚本

```bash
# 下载并安装最新版本的 Go
curl -L https://go.dev/VERSION?m=text | head -1 | xargs -I {} curl -L https://go.dev/dl/{}.darwin-amd64.pkg -o /tmp/go.pkg
sudo installer -pkg /tmp/go.pkg -target /
```

## 验证安装

安装完成后，在终端运行：

```bash
go version
```

应该看到类似输出：
```
go version go1.21.x darwin/arm64
```

## 配置环境变量

如果使用官方安装包，Go 通常会自动配置到 `/usr/local/go/bin`。

如果使用 Homebrew，Go 会安装在 `/opt/homebrew/bin/go`（Apple Silicon）或 `/usr/local/bin/go`（Intel）。

如果 `go` 命令仍然找不到，可以手动添加到 PATH：

```bash
# 对于 Apple Silicon Mac
echo 'export PATH=$PATH:/opt/homebrew/bin' >> ~/.zshrc

# 对于 Intel Mac
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.zshrc

# 重新加载配置
source ~/.zshrc
```

## 安装完成后

安装 Go 后，可以运行：

```bash
cd backend
./quick_test.sh
```

或者：

```bash
go run ./cmd/test/main.go
```

## 检查 Go 环境

运行以下命令检查 Go 环境：

```bash
go env
```

重要的环境变量：
- `GOROOT`: Go 的安装路径
- `GOPATH`: Go 工作空间路径
- `GOBIN`: Go 二进制文件安装路径

