# 指定基础镜像, 必须存在于Dockerfile 第一行
# 使用基础镜像，这个镜像是一个空的镜像，没有任何内容。且不能执行任何命令
FROM scratch

# 作者信息
MAINTAINER wolferhua <wolferhua@qq.com> 

# 切换目录，相当于cd xxx
WORKDIR /app/

# 添加文件到镜像
ADD /main /app/ 

# 执行命令
CMD ["/app/main"]