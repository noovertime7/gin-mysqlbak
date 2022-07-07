<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [gin_scaffold](#gin_scaffold)
    - [现在开始](#%E7%8E%B0%E5%9C%A8%E5%BC%80%E5%A7%8B)
    - [文件分层](#%E6%96%87%E4%BB%B6%E5%88%86%E5%B1%82)
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

- 开始前请准备状态良好的k8s集群
- 创建mysqlbak命名空间部署后端服务

```
kubectl create ns mysqlbak
```

- 创建configmap文件用于替换容器中的配置文件,**文件内容请根据实际修改**

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: mysqlbak-base-conf
  namespace: mysqlbak
data:
  base.toml: |
    [base]
        debug_mode="debug"
        time_location="Asia/Chongqing"
        cluster_url="http://yunxue521.top:32081"  # 外网访问地址
    
    [http]
        addr = ":8880"                       # 监听地址, default ":8700"
        read_timeout = 10                   # 读取超时时长
        write_timeout = 10                  # 写入超时时长
        max_header_bytes = 20               # 最大的header大小，二进制位长度
    
    [log]
        log_level = "trace"         #日志打印最低级别
        [log.file_writer]           #文件写入配置
            on = true
            log_path = "./logs/gin-mysqlbak.inf.log"
            rotate_log_path = "./logs/gin-mysqlbak.inf.log.%Y%M%D%H"
            wf_log_path = "./logs/gin-mysqlbak.wf.log"
            rotate_wf_log_path = "./logs/gin-mysqlbak.wf.log.%Y%M%D%H"
        [log.console_writer]        #工作台输出
            on = false
            color = false
    
    [swagger]
        title="gin_scaffold swagger API"
        desc="This is a sample server celler server."
        host="127.0.0.1:8880"
        base_path=""
    
    [redis]
        host="127.0.0.1:6379"
        password="123456"
        max_active = 100
        max_idle = 100
        down_grade = false
    [dingMonitor]   # 用于钉钉推送主机报警
        accessToken = "77f579efbefeefc316b55d3caea1ba1963db2f1319aa7520cbfd9626de07fffc"
        secret = "SEC72586e3f7ff6db4b2ad24eac905f308a9ddb0b1b9809af31e5623a14abb42fff"
```

- 创建数据库配置文件

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: mysqlbak-mysql-conf
  namespace: mysqlbak
data:
  mysql_map.toml: |
    [list]
        [list.default]
            driver_name = "mysql"
            data_source_name = "root:123456@tcp(127.0.0.1:3306)/gin-mysqlbak?charset=utf8&parseTime=true&loc=Asia%2FChongqing"
            max_open_conn = 20
            max_idle_conn = 10
            max_conn_life_time = 100
```



### 文件分层

```
├── README.md
├── conf            配置文件夹
│   └── dev
│       ├── base.toml
│       ├── mysql_map.toml
│       └── redis_map.toml
├── controller      控制器
│   └── demo.go
├── dao             DB数据层
│   └── demo.go
├── docs            swagger文件层
├── dto             输入输出结构层
│   └── demo.go
├── go.mod
├── go.sum
├── main.go         入口文件
├── middleware      中间件层
│   ├── panic.go
│   ├── response.go
│   ├── token_auth.go
│   └── translation.go
├── public          公共文件
│   ├── log.go
│   ├── mysql.go
│   └── validate.go
└── router          路由层
│   ├── httpserver.go
│   └── route.go
└── services        逻辑处理层
```
层次划分
控制层 --> 逻辑处理层 --> DB数据层

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
