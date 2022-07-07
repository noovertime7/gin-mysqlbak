<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [gin-mysqlbak](#gin-mysqlbak)
  - [一、现在开始](#%E4%B8%80%E7%8E%B0%E5%9C%A8%E5%BC%80%E5%A7%8B)
    - [1.1、二进制编译运行](#11%E4%BA%8C%E8%BF%9B%E5%88%B6%E7%BC%96%E8%AF%91%E8%BF%90%E8%A1%8C)
    - [1.2 Kubernetes环境部署](#12-kubernetes%E7%8E%AF%E5%A2%83%E9%83%A8%E7%BD%B2)
  - [二、文件分层](#%E4%BA%8C%E6%96%87%E4%BB%B6%E5%88%86%E5%B1%82)
  - [三、操作及页面演示](#%E4%B8%89%E6%93%8D%E4%BD%9C%E5%8F%8A%E9%A1%B5%E9%9D%A2%E6%BC%94%E7%A4%BA)
  - [四、其他](#%E5%9B%9B%E5%85%B6%E4%BB%96)
    - [log / redis / mysql / http.client 常用方法](#log--redis--mysql--httpclient-%E5%B8%B8%E7%94%A8%E6%96%B9%E6%B3%95)
    - [swagger文档生成](#swagger%E6%96%87%E6%A1%A3%E7%94%9F%E6%88%90)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# gin-mysqlbak

gin-mysqlbak:一款简单高效的Mysql数据库备份平台！

1. 请求链路日志打印，涵盖mysql/redis/request
2. 支持备份文件直接下载到本地，一键还原至原数据库。
3. 支持对接OSS对象存储存储备份文件
4. 支持主机健康检查，主机离线通过钉钉发送告警
5. 支持钉钉推送备份状态
6. 支持swagger文档生成

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
```

## 二、文件分层

```
├── bakfile   ## 备份文件存放位置
├── cmd       ## 二进制启动文件
│   └── gin-mysqlbak
├── conf     ## 配置文件存放位置
│   └── dev
│       ├── base.toml
│       └── mysql_map.toml
├── controller  ## controller层
│   ├── admin.go
│   ├── admin_login.go
│   ├── bak.go
│   ├── dashboard.go
│   ├── host.go
│   ├── public.go
│   └── task.go
├── core   ## 核心备份功能
│   └── bak.go
├── dao   ## 数据库层
│   ├── admin_login.go
│   ├── bakhistory.go
│   ├── ding.go
│   ├── host.go
│   ├── oss.go
│   ├── task.go
│   └── task_info.go
├── Dockerfile
├── docs  ## 文档
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── dto    ## 模型层
│   ├── admin.go
│   ├── admin_login.go
│   ├── bak.go
│   ├── dashboard.go
│   ├── host.go
│   └── task.go
├── go.mod
├── go.sum
├── img
│   ├── bakhistory.gif
│   ├── dashboard.gif
│   ├── hostlist.gif
│   ├── hosttask.gif
│   └── task.gif
├── kubernetes
│   ├── conf
│   │   ├── mysqlbak-base-conf.yaml
│   │   ├── mysqlbak-mysql-conf.yaml
│   │   └── mysqlbak-web-nginx-default.yaml
│   ├── gin-mysqlbak-server-deploy.yaml
│   └── gin-mysqlbak-web-deploy.yaml
├── logs
│   ├── gin-mysqlbak.inf.log
│   └── gin-mysqlbak.wf.log
├── main.go
├── middleware
│   ├── ip_auth.go
│   ├── recovery.go
│   ├── request_log.go
│   ├── response.go
│   ├── session_auth.go
│   └── translation.go
├── public
│   ├── alioss
│   │   └── alioss.go
│   ├── const.go
│   ├── ding
│   │   └── ding.go
│   ├── log.go
│   ├── params.go
│   └── util.go
├── README.md
├── router
│   ├── httpserver.go
│   └── route.go
├── services
│   └── stopAllTask.go
├── sql
│   └── gin-mysqlbak.sql
└── test
```
层次划分
控制层 --> 逻辑处理层 --> DB数据层

## 三、操作及页面演示

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



## 四、其他

### log / redis / mysql / http.client 常用方法

参考文档：https://github.com/e421083458/golang_common

### swagger文档生成

https://github.com/swaggo/swag/releases

- 下载对应操作系统的执行文件到$GOPATH/bin下面

如下：
```
➜  gin_scaffold git:(master) ✗ ll -r $GOPATH/bin
total 434168
-rwxr-xr-x  1 niuyufu  staff    13M  4  3 17:38 swag
```

- 设置接口文档参考： `controller/demo.go` 的 Bind方法的注释设置

```
// ListPage godoc
// @Summary 测试数据绑定
// @Description 测试数据绑定
// @Tags 用户
// @ID /demo/bind
// @Accept  json
// @Produce  json
// @Param polygon body dto.DemoInput true "body"
// @Success 200 {object} middleware.Response{data=dto.DemoInput} "success"
// @Router /demo/bind [post]
```

- 生成接口文档：`swag init`
- 然后启动服务器：`go run main.go`，浏览地址: http://127.0.0.1:8880/swagger/index.html
