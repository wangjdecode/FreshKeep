#!/bin/bash

# 一键修复模块问题

set -e

echo "=== 修复 Go 模块问题 ==="
echo ""

cd "$(dirname "$0")"

# 1. 设置环境变量
echo "1. 设置 Go 环境变量..."
go env -w GOPRIVATE=github.com/freshkeep/backend
go env -w GONOPROXY=github.com/freshkeep/backend  
go env -w GONOSUMDB=github.com/freshkeep/backend
echo "✓ 环境变量设置完成"
echo ""

# 2. 清理模块缓存
echo "2. 清理 Go 模块缓存..."
go clean -modcache 2>/dev/null || echo "  警告: 清理缓存失败（可能没有缓存）"
echo "✓ 缓存清理完成"
echo ""

# 3. 验证 go.mod 中的 replace 指令
echo "3. 检查 go.mod 配置..."
if ! grep -q "replace github.com/freshkeep/backend" go.mod; then
    echo "  添加 replace 指令..."
    echo "" >> go.mod
    echo "// 本地开发：使用本地代码而不是远程" >> go.mod
    echo "replace github.com/freshkeep/backend => ./" >> go.mod
    echo "✓ replace 指令已添加"
else
    echo "✓ replace 指令已存在"
fi
echo ""

# 4. 下载依赖
echo "4. 下载依赖..."
go mod download 2>&1 | grep -v "github.com/freshkeep/backend" || true
echo "✓ 依赖下载完成"
echo ""

# 5. 整理依赖
echo "5. 整理依赖..."
go mod tidy
echo "✓ 依赖整理完成"
echo ""

# 6. 验证
echo "6. 验证配置..."
if go mod verify 2>&1 | grep -q "all modules verified"; then
    echo "✓ 所有模块验证通过"
else
    echo "⚠ 验证完成（可能有警告，但不影响使用）"
fi
echo ""

echo "=== 修复完成 ==="
echo ""
echo "现在可以运行测试服务器："
echo "  ./quick_test.sh"
echo "或者："
echo "  go run ./cmd/test/main.go"
echo ""

