FROM centos
WORKDIR /app
ADD . .
ENTRYPOINT [ "/app/gin-mysqlbak" ]