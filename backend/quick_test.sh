#!/bin/bash

# 快速测试脚本 - 运行简化的测试服务器

set -e

echo "=== FreshKeep 后端快速测试 ==="
echo ""

cd "$(dirname "$0")"

# 检查 Go 是否安装
if ! command -v go &> /dev/null; then
    echo "错误: 未找到 Go 编译器"
    exit 1
fi

echo "✓ Go 版本: $(go version)"
echo ""

echo "检查依赖..."

# 设置私有模块，避免尝试从远程获取本地模块
if ! go env GOPRIVATE | grep -q "github.com/freshkeep/backend"; then
    echo "设置 Go 环境变量（避免从远程获取本地模块）..."
    go env -w GOPRIVATE=github.com/freshkeep/backend 2>/dev/null || true
    go env -w GONOPROXY=github.com/freshkeep/backend 2>/dev/null || true
    go env -w GONOSUMDB=github.com/freshkeep/backend 2>/dev/null || true
fi

if [ ! -f "go.sum" ] || [ ! -s "go.sum" ]; then
    echo "下载依赖（首次运行需要下载，可能需要一些时间）..."
    GOPRIVATE=github.com/freshkeep/backend go mod download || {
        echo ""
        echo "⚠ 依赖下载失败"
        echo "请手动运行以下命令："
        echo "  go env -w GOPRIVATE=github.com/freshkeep/backend"
        echo "  go mod download"
        echo "  go mod tidy"
        echo ""
        echo "详细说明请查看: FIX_MODULE.md"
        exit 1
    }
    go mod tidy || {
        echo "警告: go mod tidy 失败，但可能不影响运行"
    }
    echo "✓ 依赖下载完成"
else
    echo "✓ 依赖已存在"
fi
echo ""

echo "启动测试服务器..."
echo "服务器将在 http://localhost:8000 启动"
echo "按 Ctrl+C 停止"
echo ""

# 运行测试服务器
go run ./cmd/test/main.go

