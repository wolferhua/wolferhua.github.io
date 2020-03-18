#!/usr/bin/env bash 

# 执行编译
# scratch 镜像完全是空的，什么东西也不包含，所以生成main时候要按照下面的方式生成，使生成的main静态链接所有的库：
CGO_ENABLED=0 GOOS=linux go build -v -a -installsuffix cgo -o main .


# 执行构建
docker build -t goweb:scratch .