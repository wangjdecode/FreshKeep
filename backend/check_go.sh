#!/bin/bash

# Go 环境检查脚本

echo "=== Go 环境检查 ==="
echo ""

# 检查 Go 是否安装
if command -v go &> /dev/null; then
    echo "✓ Go 已安装"
    echo "  版本: $(go version)"
    echo ""
    
    # 检查 Go 环境
    echo "Go 环境信息:"
    echo "  GOROOT: $(go env GOROOT)"
    echo "  GOPATH: $(go env GOPATH)"
    echo "  GOBIN:  $(go env GOBIN)"
    echo ""
    
    # 检查 Go 版本
    GO_VERSION=$(go version | awk '{print $3}')
    echo "✓ Go 版本: $GO_VERSION"
    
    # 检查版本是否 >= 1.21
    MAJOR=$(echo $GO_VERSION | cut -d. -f1 | tr -d 'go')
    MINOR=$(echo $GO_VERSION | cut -d. -f2)
    
    if [ "$MAJOR" -gt 1 ] || ([ "$MAJOR" -eq 1 ] && [ "$MINOR" -ge 21 ]); then
        echo "✓ Go 版本符合要求 (>= 1.21)"
    else
        echo "⚠ 警告: Go 版本低于 1.21，建议升级"
    fi
    echo ""
    
    echo "=== 检查完成 ==="
    echo ""
    echo "可以运行测试服务器:"
    echo "  ./quick_test.sh"
    echo "或者:"
    echo "  go run ./cmd/test/main.go"
    
else
    echo "✗ Go 未安装"
    echo ""
    echo "请先安装 Go:"
    echo ""
    echo "方式一（推荐）: 使用 Homebrew"
    echo "  brew install go"
    echo ""
    echo "方式二: 下载官方安装包"
    echo "  访问: https://go.dev/dl/"
    echo ""
    echo "详细说明请查看: INSTALL_GO.md"
    exit 1
fi

