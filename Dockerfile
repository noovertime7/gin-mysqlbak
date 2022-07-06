FROM centos
WORKDIR /app
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezon
ADD . .
ENTRYPOINT [ "/app/gin-mysqlbak" ]