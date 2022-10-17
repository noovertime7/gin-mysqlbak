集群版本需要部署Agent客户端，部署配置请参考 [agent部署](https://github.com/noovertime7/gin-mysqlbak-agent)

前端服务部署请参考 [前端部署](https://github.com/noovertime7/gin-mysqlbak/tree/main/front)
## 二进制部署
::: warning  注意
二进制安装开始前，请确认系统中已安装go环境与git环境 
:::

### 1、 克隆仓库
```shell
git clone https://github.com/noovertime7/gin-mysqlbak.git
```
### 2、修改配置文件
编译开始前，请先修改配置文件，需要修改的位置已进行标注

创建conf目录，在conf目录下创建config.ini，文件内容如下
```ini {38-42}
[base]
debug_mode="debug"
time_location="Asia/Chongqing"
download_url="http://127.0.0.1:8880"  ## 已废弃，可不修改
[http]
addr = ":8880"                       # 监听地址, default ":8700"
read_timeout = 10                   # 读取超时时长
write_timeout = 10                  # 写入超时时长
max_header_bytes = 20               # 最大的header大小，二进制位长度
[HostLostAlarms]
enable = false          ## 是否启用主机离线告警，启用后会持续监控主机在线，离线后通过下面的钉钉ac发送告警
accessToken = "xxxxx"   ## 钉钉accessToken
secret = "xxxx"         ## 钉钉secret
[dingProxyAgent]       
enable = false            ## 是否启用钉钉发送代理，适用于无外网环境下钉钉消息的发送
addr = 127.0.0.1:39999    
title = "测试"
content = "测试环境"
[log]
log_level = "trace"         #日志打印最低级别
[log_file_writer]
on = true
log_path = "./logs/gin-mysqlbak.inf.log"
rotate_log_path = "./logs/gin-mysqlbak.inf.log.%Y%M%D%H"
wf_log_path = "./logs/gin-mysqlbak.wf.log"
rotate_wf_log_path = "./logs/gin-mysqlbak.wf.log.%Y%M%D%H"
[log_console_writer]      
on = false   #工作台输出
color = false
[swagger]
title="gin-mysqlbak swagger API"
desc="This is a sample server celler server."
host="127.0.0.1:8880"
base_path=""
[jaeger] 
enable = false  ## 是否开启链路追踪,如开启后需要部署jaeger
addr = 127.0.0.1:6831
[mysql]
host = 127.0.0.1  ## 数据库地址连接信息
port = 3306
user = root
password = {{your_password}}
dbname = gin-mysqlbak
[app] ## 应用相关配置，用于主机管理头像，可不修改
mysqlAvatar: http://qiniu.yunxue521.top/mysql.jpeg  
elasticAvatar: http://qiniu.yunxue521.top/elastic.jpg
[cluster]
## 集群任务同步周期(分钟)
clusterSyncPeriod = "*/30 * * * *"  ## 半小时同步一次集群内所有任务到服务端统一管理
```
### 3、初始化数据库
将工程内sql文件夹下的init_3.0.0.sql拷贝至mysql主机根目录下

当前版本仍需要手动刷入sql文件，自动初始化数据库将在3.0.1支持
```sql
## 创建数据库
CREATE DATABASE `gin-mysqlbak`;
use gin-mysqlbak;
## 刷入初始化sql文件
source /init_3.0.0.sql
```
### 4、编译二进制文件并运行
手动编译或选择直接下载[release](https://github.com/noovertime7/gin-mysqlbak/releases/tag/v3.0.0/)中的文件
```shell
go build -o gin-mysqlbak-server main.go
```
### 5、运行
```shell
./gin-mysqlbak-server
```
## 容器部署
部署开始前,请先创建config.ini配置文件夹，并根据实际修改相关配置,通过--volume参数可以挂载到容器内部
```shell
docker run -itd --name gin-mysql-server \
-p 8880:8880 \
--restart=always \
-v /root/config.ini:/app/conf/config.ini \
chenteng/gin-mysqlbak-server:3.0.0-SP1
```
## Kubernetes部署
::: tip 提示
安装开始前，请确认拥有一个健康可用的Kubernetes集群，支持K3S
:::
### 1、创建命名空间
命名空间不存在请先创建命名空间，修改命名空间请修改yaml中namespace选项
```shell
kubectl create ns mysqlbak
```
### 2、创建configmap挂载配置文件
新建mysqlbak-ini.yaml文件内容如下
```yaml
---
apiVersion: v1
data:
  config.ini: |
    [base]
    debug_mode="debug"
    time_location="Asia/Chongqing"
    download_url="http://127.0.0.1:19009"   ## 用于下载备份文件使用，需要与server端svc开放的地址端口相同
    [http]
    addr = ":8880"                       # 监听地址, default ":8700"
    read_timeout = 10                   # 读取超时时长
    write_timeout = 10                  # 写入超时时长
    max_header_bytes = 20               # 最大的header大小，二进制位长度
    [HostLostAlarms]
    enable = false  ## 是否启动主机检测 
    accessToken = "xxx"  ## 主机检测钉钉报警配置
    secret = "xxx"
    [dingProxyAgent]
    enable = false    ## 用于无外网环境下发送钉钉消息
    addr = 127.0.0.1:39999
    title = "公司测试环境备份平台"
    content = "公司测试环境备份平台"
    [swagger]
    title="gin-mysqlbak swagger API"
    desc="This is a sample server celler server."
    host="127.0.0.1:8880"
    base_path=""
    [jaeger]
    enable = false
    addr = 127.0.0.1:6831
    [mysql]
    host = 127.0.0.1
    port = 3306
    user = root
    password = chenteng
    dbname = gin-mysqlbak
    [app] ## 应用相关配置
    mysqlAvatar: http://qiniu.yunxue521.top/mysql.jpeg
    elasticAvatar: http://qiniu.yunxue521.top/elastic.jpg
    [cluster]
    ## 集群任务同步周期(分钟)
    clusterSyncPeriod = "*/30 * * * *"
kind: ConfigMap
metadata:
  annotations: {}
  name: mysqlbak-ini-conf
  namespace: mysqlbak
```
### 3、创建deployment控制器 & service
新建gin-mysqlbak-server-deploy.yaml，文件内容去下
```yaml
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: gin-mysqlbak-server
  namespace: mysqlbak
spec:
  progressDeadlineSeconds: 600
  replicas: 1
  revisionHistoryLimit: 10
  selector:
    matchLabels:
      app: gin-mysqlbak-server
  strategy:
    rollingUpdate:
      maxSurge: 25%
      maxUnavailable: 25%
    type: RollingUpdate
  template:
    metadata:
      creationTimestamp: null
      labels:
        app: gin-mysqlbak-server
    spec:
      containers:
        - image: 'chenteng/gin-mysqlbak-server:3.0.0'
          imagePullPolicy: IfNotPresent
          name: gin-mysqlbak-server
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
            - mountPath: /app/conf/config.ini
              name: iniconf
              subPath: config.ini
            - mountPath: /app/bakfile
              name: mysqlbak-data
              subPath: mysql-data
      dnsPolicy: ClusterFirst
      restartPolicy: Always
      schedulerName: default-scheduler
      terminationGracePeriodSeconds: 30
      volumes:
        - hostPath:
            path: /data/gin-mysqlbak
            type: DirectoryOrCreate
          name: mysqlbak-data
        - configMap:
            defaultMode: 420
            name: mysqlbak-ini-conf
          name: iniconf
---
apiVersion: v1
kind: Service
metadata:
  labels:
    app: gin-mysqlbak-server
  name: gin-mysqlbak-server
  namespace: mysqlbak
spec:
  externalTrafficPolicy: Cluster
  ipFamilyPolicy: SingleStack
  ports:
    - nodePort: 19009   ## server端地址
      port: 8880
      protocol: TCP
      targetPort: 8880
  selector:
    app: gin-mysqlbak-server
  sessionAffinity: None
  type: NodePort
```
#### 4、部署到K8S集群
```shell
kubectl apply -f mysqlbak-ini.yaml && kubectl apply -f gin-mysqlbak-server-deploy.yaml

kubectl get pod -n mysqlbak  ## pod状态均为Running则部署成功
```