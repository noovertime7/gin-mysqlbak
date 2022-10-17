<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [gin-mysqlbak](#gin-mysqlbak)
  - [一、现在开始](#%E4%B8%80%E7%8E%B0%E5%9C%A8%E5%BC%80%E5%A7%8B)
    - [1.1、二进制编译运行](#11%E4%BA%8C%E8%BF%9B%E5%88%B6%E7%BC%96%E8%AF%91%E8%BF%90%E8%A1%8C)
    - [1.2 Kubernetes环境部署](#12-kubernetes%E7%8E%AF%E5%A2%83%E9%83%A8%E7%BD%B2)
  - [二、架构](#%E4%BA%8C%E6%9E%B6%E6%9E%84)
    - [单机版本](#%E5%8D%95%E6%9C%BA%E7%89%88%E6%9C%AC)
    - [集群版本](#%E9%9B%86%E7%BE%A4%E7%89%88%E6%9C%AC)
  - [三、操作及页面演示](#%E4%B8%89%E6%93%8D%E4%BD%9C%E5%8F%8A%E9%A1%B5%E9%9D%A2%E6%BC%94%E7%A4%BA)
    - [单机版本操作演示](#%E5%8D%95%E6%9C%BA%E7%89%88%E6%9C%AC%E6%93%8D%E4%BD%9C%E6%BC%94%E7%A4%BA)
    - [集群版本备份功能演示](#%E9%9B%86%E7%BE%A4%E7%89%88%E6%9C%AC%E5%A4%87%E4%BB%BD%E5%8A%9F%E8%83%BD%E6%BC%94%E7%A4%BA)
  - [四、其他](#%E5%9B%9B%E5%85%B6%E4%BB%96)
    - [log / redis / mysql / http.client 常用方法](#log--redis--mysql--httpclient-%E5%B8%B8%E7%94%A8%E6%96%B9%E6%B3%95)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# gin-mysqlbak
<p> 
<img src=https://img.shields.io/github/languages/top/noovertime7/gin-mysqlbak  alt="code-size" />
<img src="https://img.shields.io/github/languages/code-size/noovertime7/gin-mysqlbak" alt="code-size" />
<img src="https://img.shields.io/github/last-commit/noovertime7/gin-mysqlbak" alt="code-size"/>
</p>

gin-mysqlbak:一款简单高效、支持多集群统一备份的Mysql数据库备份平台！

1. 请求链路日志打印，涵盖mysql/redis/request，集群版本支持jaeger链路追踪
2. 支持备份文件直接下载到本地，可一键还原至原数据库。
3. 支持对接S3协议对象存储存储备份文件，现已支持minio、阿里oss
4. 支持主机健康检查，主机离线通过钉钉发送告警
5. 支持钉钉推送备份状态，成功失败发送钉钉消息
6. 通过部署agent完成异地多节点备份，备份任务统一管理，备份数据集中存储
7. 支持ElasticSearch快照管理，快照信息查看


项目地址：https://github.com/noovertime7/gin-mysqlbak

前端地址: https://github.com/noovertime7/gin-mysql-web

agent地址: https://github.com/noovertime7/gin-mysqlbak-agent
## 一、现在开始

**注意: 安装开始前，请先创建gin-mysqlbak数据库，刷入sql文件初始化数据库，sql文件在项目sql文件夹下**

### 1.1、二进制编译运行

- 开始前请确保机器已安装go编译环境

- 安装软件依赖
go mod使用请查阅：

https://blog.csdn.net/e421083458/article/details/89762113
```shell
git clone https://github.com/noovertime7/gin-mysqlbak.git
cd gin-mysqlbak
go mod tidy
```
- 确保正确配置了 conf/mysql_map.toml、conf/config.ini

- 运行入口main.go (默认监听8880端口)

```shell
go run main.go
```
### 1.2 Docker部署

#### 1.2.1 后端服务部署

首先请在root目录下准备好三个配置文件，base.toml不用改 ,配置文件实例在项目conf文件夹下

```docker
docker run -itd --name gin-mysql-server \
-p 8880:8880 \
-v /root/config.ini:/app/conf/config.ini \
-v /root/base.toml:/app/conf/dev/base.toml \ 
-v /root/mysql_map.toml:/app/conf/dev/mysql_map.toml \
harbor-tj.ts-it.cn:63333/mysqlbak/gin-mysqlbak-server:2.0.3-SP3
```

#### 1.2.2 前端服务部署

首先请在root目录下准备好前端配置文件

```docker
docker run -itd --name gin-mysql-web -p 8881:80 -v /root/default.conf:/etc/nginx/conf.d/default.conf  harbor-tj.ts-it.cn:63333/mysqlbak/gin-mysqlbak-web:2.0.2-SP3
```

#### 1.2.3 agent部署

首先请在root目录下准备好agent配置文件，agent配置文件在agent仓库下的/domain/config/config.ini

```docker
docker run -itd --name gin-mysql-agent  \
-- net=host --restart=always \
-v /root/config.ini:/app/domain/config/config.ini \
-v /root/bakfile:/app/bakfile \
harbor-tj.ts-it.cn:63333/mysqlbak/gin-mysqlbak-agent:2.0.4
```

### 1.3 Kubernetes环境部署

- 开始前请准备状态良好的k8s集群,**请根据实际环境修改conf文件夹下的配置文件**
- 创建mysqlbak命名空间部署后端服务

```
kubectl create ns mysqlbak  ## 创建命名空间
cd gin-mysqlbak/kubernetes && kubectl apply -f ./conf  ## 创建configmap配置文件
kubectl apply -f gin-mysqlbak-server-deploy.yaml  ## 创建后端服务
kubectl apply -f gin-mysqlbak-server-web.yaml  ## 创建前端服务
```

## 二、架构

### 2.1、单机版本

实现原理：

单机版本比较简单，使用server端备份能力进行备份，使用cron表达式创建备份任务，两种备份方式，xorm备份失败后，使用mysqldump进行备份，确保备份任务能百分百成功

 ![aloneserver.jpg](https://github.com/noovertime7/gin-mysqlbak/blob/main/img/aloneserver.jpg?raw=true) 

### 2.2、集群版本

实现原理：

原先的server端作为微服务网关，各个agent通过服务端接口进行服务的注册发现，上报agent信息，server端通过不通的服务名调用agent，agent完成数据备份与存储

 ![cluster.jpg](https://github.com/noovertime7/gin-mysqlbak/blob/main/img/cluster.jpg?raw=true) 



## 三、操作及页面演示

### 3.1、单机版本操作演示

- 首页大盘展示

 ![dashboard.gif](https://github.com/noovertime7/gin-mysqlbak/blob/main/img/dashboard.gif?raw=true)    

- 备份主机操作演示

![hostlist.gif](https://github.com/noovertime7/gin-mysqlbak/blob/main/img/hostlist.gif?raw=true) 

- 备份主机添加备份任务操作演示

![hosttask.gif](https://github.com/noovertime7/gin-mysqlbak/blob/main/img/hosttask.gif?raw=true) 

- 任务列表操作展示

 ![task.gif](https://github.com/noovertime7/gin-mysqlbak/blob/main/img/task.gif?raw=true) 

- 备份历史记录操作展示

 ![bakhistory.gif](https://github.com/noovertime7/gin-mysqlbak/blob/main/img/bakhistory.gif?raw=true) 

### 3.2、集群版本备份功能演示

- 集群列表

 ![cluster_servicelist.gif](https://github.com/noovertime7/gin-mysqlbak/blob/main/img/cluster_servicelist.gif?raw=true) 

- 集群主机列表

 ![cluster_host.gif](https://github.com/noovertime7/gin-mysqlbak/blob/main/img/cluster_host.gif?raw=true) 

- 集群任务列表

 ![cluster_tasklist.gif](https://github.com/noovertime7/gin-mysqlbak/blob/main/img/cluster_tasklist.gif?raw=true) 

- 集群历史记录列表

 ![cluster_history.gif](https://github.com/noovertime7/gin-mysqlbak/blob/main/img/cluster_history.gif?raw=true) 

## 四、FAQ

1.为什么服务端会有三个配置文件？

前期使用脚手架创建，后面发现对脚手架太过依赖，正在逐步去除脚手架相关配置，主要配置文件为config.ini

2.目前的工作方向？

目前正在使用ant design vue重构前端代码，主要增加权限管理、安全相关功能，也是为了前端美化页面

## 五、其他
这个项目的前端后端都是我一个人完成，所以会有很多不足，希望能提issue，或者提PR一起来完善这个项目
