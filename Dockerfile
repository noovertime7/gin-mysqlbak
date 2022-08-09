FROM centos
WORKDIR /app
RUN rm -rf /etc/yum.repos.d/*
ADD Centos-8.repo /etc/yum.repos.d/
RUN yum clean all && yum makecache && yum install -y mysql
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezon
ADD . .
ENTRYPOINT [ "/app/gin-mysqlbak" ]