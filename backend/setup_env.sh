#!/bin/bash

# 设置 Go 环境变量，避免从远程获取本地模块

echo "设置 Go 环境变量..."

# 设置私有模块
go env -w GOPRIVATE=github.com/freshkeep/backend

# 设置不通过代理获取这些模块
go env -w GONOPROXY=github.com/freshkeep/backend

# 设置不验证这些模块的校验和
go env -w GONOSUMDB=github.com/freshkeep/backend

echo "✓ 环境变量设置完成"
echo ""
echo "当前设置:"
echo "  GOPRIVATE: $(go env GOPRIVATE)"
echo "  GONOPROXY: $(go env GONOPROXY)"
echo "  GONOSUMDB: $(go env GONOSUMDB)"
echo ""

