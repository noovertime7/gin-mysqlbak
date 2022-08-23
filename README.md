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

gin-mysqlbak:一款简单高效、支持多集群统一备份的Mysql数据库备份平台！

1. 请求链路日志打印，涵盖mysql/redis/request，集群版本支持jaeger链路追踪
2. 支持备份文件直接下载到本地，可一键还原至原数据库。
3. 支持对接S3协议对象存储存储备份文件，现已支持minio、阿里oss
4. 支持主机健康检查，主机离线通过钉钉发送告警
5. 支持钉钉推送备份状态，成功失败早知道
6. 通过部署agent完成异地多节点备份，备份任务统一管理，备份数据集中存储

项目地址：https://github.com/noovertime7/gin-mysqlbak
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
- 确保正确配置了 conf/mysql_map.toml、conf/base.toml

- 运行入口main.go (默认监听8880端口)

```shell
go run main.go
```
### 1.2 Kubernetes环境部署

- 开始前请准备状态良好的k8s集群,**请根据实际环境修改conf文件夹下的配置文件**
- 创建mysqlbak命名空间部署后端服务

```
kubectl create ns mysqlbak  ## 创建命名空间
cd gin-mysqlbak/kubernetes && kubectl apply -f ./conf  ## 创建configmap配置文件
kubectl apply -f gin-mysqlbak-server-deploy.yaml  ## 创建后端服务
kubectl apply -f gin-mysqlbak-server-web.yaml  ## 创建前端服务
```

## 二、架构

### 单机版本

实现原理：

单机版本比较简单，使用server端备份能力进行备份，使用cron表达式创建备份任务，两种备份方式，xorm备份失败后，使用mysqldump进行备份，确保备份任务能百分百成功

 ![aloneserver.jpg](https://github.com/noovertime7/gin-mysqlbak/blob/main/img/aloneserver.jpg?raw=true) 

### 集群版本

实现原理：

原先的server端作为微服务网关，各个agent通过服务端接口进行服务的注册发现，上报agent信息，server端通过不通的服务名调用agent，agent完成数据备份与存储

 ![cluster.jpg](https://github.com/noovertime7/gin-mysqlbak/blob/main/img/cluster.jpg?raw=true) 



## 三、操作及页面演示

### 单机版本操作演示

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

### 集群版本备份功能演示





## 四、其他

### log / redis / mysql / http.client 常用方法

参考文档：https://github.com/e421083458/golang_common
