# 指定基础镜像, 必须存在于Dockerfile 第一行
FROM daocloud.io/golang


# 作者信息
MAINTAINER wolferhua <wolferhua@qq.com> 

# 添加文件到镜像
ADD . /usr/src/app

# 切换目录，相当于cd xxx
WORKDIR /usr/src/app

# 执行编译g
RUN go build -o app main.go 

# 切换目录
WORKDIR /app

# 将编译后的移动到 /app 目录
RUN mv /usr/src/app/app .

# 开放端口，在运行时，并不会正在的开放端口。还是需要-p
# 这里只是标明需要使用的端口
EXPOSE 8080

# 镜像启动时执行的命令

CMD ["./app"]

# 构建命令
#docker build -t goweb .