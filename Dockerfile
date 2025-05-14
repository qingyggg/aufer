FROM golang:1.23-alpine AS builder
WORKDIR /app
ENV GOPROXY="https://goproxy.cn,direct"
RUN apk add --no-cache bash
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    chmod +x build.sh && ./build.sh

# Create a minimal image for running the application
FROM alpine:latest

WORKDIR /app
MAINTAINER "mols"

# Install necessary packages
RUN apk --no-cache add ca-certificates tzdata bash

# Copy the binary from builder
COPY --from=builder /app/output/bin/hertz_service /app/hertz_service
COPY copy-env.sh /app/copy-env.sh
COPY ./static/ /app/static

# Set execution permissions
RUN chmod +x /app/hertz_service /app/copy-env.sh

# 创建环境变量文件目录
RUN mkdir -p /app/envs

EXPOSE 18014
ENTRYPOINT ["/app/copy-env.sh", "/app/hertz_service"]



















