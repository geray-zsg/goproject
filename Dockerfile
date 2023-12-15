# 使用官方 BusyBox 镜像作为基础镜像
FROM busybox:1.28.4

# 添加一个具有 UID 1001 的普通用户
RUN adduser -D -u 1001 myuser

# 切换到该用户
USER myuser

# 设置工作目录
WORKDIR /data/
