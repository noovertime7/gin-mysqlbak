FROM golang:alpine AS builder
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct
WORKDIR /build
COPY . .
RUN go mod tidy
RUN go build -ldflags="-s -w" -o gin-mysqlbak ./main.go

FROM centos
WORKDIR /app
ENV TZ Asia/Shanghai
COPY --from=builder /build/gin-mysqlbak /app/gin-mysqlbak
COPY --from=builder /build/conf/config.ini /app/conf/config.ini
COPY --from=builder /build/docker/mysqldump /usr/bin
RUN chmod 777 /usr/bin/mysqldump
EXPOSE 8880
CMD ["./gin-mysqlbak"]