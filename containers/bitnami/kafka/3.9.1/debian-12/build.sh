#!/bin/bash

# Kafka 3.9.1 镜像构建脚本

set -e

REGISTRY=${REGISTRY:-""}
IMAGE_NAME=${IMAGE_NAME:-"kafka"}
VERSION=${VERSION:-"3.9.1"}
PLATFORM=${PLATFORM:-"linux/amd64,linux/arm64"}

echo "构建 Kafka ${VERSION} 镜像..."

# 构建标准 x86 镜像
echo "构建标准 x86 镜像..."
docker build -t ${REGISTRY}${IMAGE_NAME}:${VERSION} -f Dockerfile .

# 构建 ARM64 镜像 (如果需要)
if [[ "$PLATFORM" == *"arm64"* ]]; then
    echo "构建 ARM64 镜像..."
    docker build -t ${REGISTRY}${IMAGE_NAME}:${VERSION}-arm64 -f Dockerfile.arm64 .
fi

echo "镜像构建完成!"
echo "标准镜像: ${REGISTRY}${IMAGE_NAME}:${VERSION}"

if [[ "$PLATFORM" == *"arm64"* ]]; then
    echo "ARM64镜像: ${REGISTRY}${IMAGE_NAME}:${VERSION}-arm64"
fi

echo ""
echo "使用方法:"
echo "  docker run -d --name kafka ${REGISTRY}${IMAGE_NAME}:${VERSION}"
echo ""
echo "或者使用 docker-compose:"
echo "  docker-compose up -d" 