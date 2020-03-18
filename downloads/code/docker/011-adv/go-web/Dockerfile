FROM daocloud.io/golang:alpine  AS  goweb-build 
# 添加文件到镜像
COPY ./main.go /app/ 
# 切换目录，相当于cd xxx
WORKDIR /app/


RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./main.go



# 使用基础镜像，这个镜像是一个空的镜像，没有任何内容。且不能执行任何命令
FROM scratch

# 作者信息
MAINTAINER wolferhua <wolferhua@qq.com> 

# 切换目录，相当于cd xxx
WORKDIR /app/

# 添加文件到镜像
COPY --from=goweb-build /app/main /app/ 

# 执行命令
CMD ["/app/main"]

# docker build -t 192.168.65.95:5000/goweb:goweb .
# docker push 192.168.65.95:5000/goweb:goweb