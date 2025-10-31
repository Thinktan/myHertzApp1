# ========= 第一阶段：构建阶段 =========
#FROM golang:1.24-alpine AS builder
FROM --platform=linux/amd64 golang:1.24-alpine AS builder

# 安装必要组件
RUN #apk add --no-cache git
RUN apk add --no-cache git build-base

# 设置工作目录
WORKDIR /app

# 复制本地代码
COPY . .

# 初始化 mod（如未初始化）
RUN go mod tidy

# 构建二进制文件
RUN #go build -o main .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main .

# ========= 第二阶段：运行阶段 =========
#FROM alpine:latest
FROM --platform=linux/amd64 alpine:latest

WORKDIR /root/

# 拷贝构建好的二进制文件
COPY --from=builder /app/main .

# 暴露服务端口
EXPOSE 9000

# 启动命令
CMD ["./main"]