FROM golang:1.23-alpine AS builder

WORKDIR /app

# 设置 GOPROXY
ENV GOPROXY="https://goproxy.cn,direct"

# 安装 bash
RUN apk add --no-cache bash

# 复制 go.mod 和 go.sum 文件
COPY go.mod go.sum ./

# 使用缓存挂载下载依赖
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# 复制整个项目
COPY . .

# 切换到评论服务目录并执行构建
WORKDIR /app/cmd/comment
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    chmod +x build.sh && ./build.sh

# 创建运行环境
FROM alpine:latest

WORKDIR /app
MAINTAINER "mols"

# 安装运行时依赖
RUN apk --no-cache add ca-certificates tzdata bash

# 从构建阶段复制二进制文件
COPY --from=builder /app/cmd/comment/output/bin/aufer-comment-service /app/aufer-comment-service
COPY --from=builder /app/copy-env.sh /app/copy-env.sh

# 设置执行权限
RUN chmod +x /app/aufer-comment-service /app/copy-env.sh

# 创建环境变量文件目录
RUN mkdir -p /app/envs

EXPOSE 18012

# 使用绝对路径运行服务
ENTRYPOINT ["/app/copy-env.sh", "/app/aufer-comment-service"]










