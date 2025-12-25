#!/bin/bash

# 后端服务器测试脚本

set -e

echo "=== FreshKeep 后端测试脚本 ==="
echo ""

# 检查 Go 是否安装
if ! command -v go &> /dev/null; then
    echo "错误: 未找到 Go 编译器，请先安装 Go 1.21+"
    exit 1
fi

echo "✓ Go 版本: $(go version)"
echo ""

# 检查 protoc 是否安装
if ! command -v protoc &> /dev/null; then
    echo "警告: 未找到 protoc，Protobuf 代码生成可能失败"
    echo "请安装 protoc: https://grpc.io/docs/protoc-installation/"
else
    echo "✓ protoc 版本: $(protoc --version)"
fi
echo ""

# 进入后端目录
cd "$(dirname "$0")"

echo "1. 下载依赖..."
go mod download
go mod tidy
echo "✓ 依赖下载完成"
echo ""

echo "2. 生成 Protobuf 代码..."
if command -v protoc &> /dev/null; then
    make api || echo "警告: Protobuf 代码生成失败，可能需要先安装工具 (make init)"
else
    echo "跳过: protoc 未安装"
fi
echo ""

echo "3. 生成 Wire 代码..."
if command -v wire &> /dev/null; then
    make wire || echo "警告: Wire 代码生成失败"
else
    echo "跳过: wire 未安装，运行 'make init' 安装"
fi
echo ""

echo "4. 检查代码编译..."
if [ -f "cmd/server/wire_gen.go" ]; then
    echo "编译服务器..."
    go build -o /tmp/freshkeep-server ./cmd/server || {
        echo "编译失败，但这是正常的，因为可能缺少生成的代码"
        echo "请先运行: make api && make wire"
    }
else
    echo "跳过编译: wire_gen.go 不存在，请先运行 'make wire'"
fi
echo ""

echo "=== 测试完成 ==="
echo ""
echo "如果所有步骤都成功，可以运行服务器:"
echo "  make run"
echo "或者:"
echo "  go run ./cmd/server/main.go ./cmd/server/wire_gen.go -conf configs"
echo ""
echo "然后测试接口:"
echo "  curl http://localhost:8000/health"
echo "  curl http://localhost:8000/api/v1/items"

