#!/bin/bash

# 创建目录
mkdir -p /app/envs

# 复制环境变量文件
cp /envs/dev.env /app/envs/dev.env
cp /envs/prod.env /app/envs/prod.env

echo "环境变量文件已复制"

# 执行原始命令
exec "$@" 