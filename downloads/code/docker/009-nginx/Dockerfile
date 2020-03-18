# 指定基础镜像, 必须存在于Dockerfile 第一行
FROM  daocloud.io/nginx


# 作者信息
MAINTAINER wolferhua <wolferhua@qq.com>


# 添加文件到镜像
# 将 index.html 添加到 /usr/share/nginx/html
ADD ./index.html /usr/share/nginx/html
ADD ./www.conf /etc/nginx/conf.d/