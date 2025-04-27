# 第一阶段：构建应用
FROM golang:1.24-alpine AS build

# 安装依赖
RUN apk add --no-cache git ca-certificates tzdata

# 设置工作目录
WORKDIR /app

# 复制Go模块定义
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 设置时区
ENV TZ=Asia/Shanghai

# 编译
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o arknights-cricket-fighting .

# 第二阶段：创建最终镜像
FROM alpine:latest

# 安装依赖
RUN apk --no-cache add ca-certificates tzdata

# 从构建阶段复制编译好的应用
COPY --from=build /app/arknights-cricket-fighting /app/arknights-cricket-fighting

# 复制静态文件和视图模板
COPY --from=build /app/static /app/static
COPY --from=build /app/app/views /app/app/views

# 设置工作目录
WORKDIR /app

# 设置环境变量
ENV TZ=Asia/Shanghai
ENV PORT=18080

# 暴露端口
EXPOSE 18080

# 启动命令
CMD ["/app/arknights-cricket-fighting"] 