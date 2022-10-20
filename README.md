<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [gin-mysqlbak](#gin-mysqlbak)
  - [一、现在开始](#%E4%B8%80%E7%8E%B0%E5%9C%A8%E5%BC%80%E5%A7%8B)
  - [二、功能演示](#%E4%BA%8C%E5%8A%9F%E8%83%BD%E6%BC%94%E7%A4%BA)
  - [三、FAQ](#%E5%9B%9Bfaq)
  - [四、其他](#%E4%BA%94%E5%85%B6%E4%BB%96)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# gin-mysqlbak
<p> 
<img src=https://img.shields.io/github/languages/top/noovertime7/gin-mysqlbak  alt="code-size" />
<img src="https://img.shields.io/github/languages/code-size/noovertime7/gin-mysqlbak" alt="code-size" />
<img src="https://img.shields.io/github/last-commit/noovertime7/gin-mysqlbak" alt="code-size"/>
</p>

gin-mysqlbak:一款简单高效、支持多集群统一备份的数据库备份平台，现已支持mysql与ElasticSearch备份。

1. 请求链路日志打印，涵盖mysql/redis/request，集群版本支持jaeger链路追踪
2. 支持备份文件直接下载到本地，可一键还原至原数据库。
3. 支持对接S3协议对象存储存储备份文件，现已支持minio、阿里oss
4. 支持主机健康检查，主机离线通过钉钉发送告警
5. 支持钉钉推送备份状态，成功失败发送钉钉消息
6. 通过部署agent完成异地多节点备份，server作为微服务网关，备份任务统一管理，备份数据集中存储
7. 支持ElasticSearch快照管理，快照信息查看
8. 数据加密存储，数据库敏感信息密文保存，保证数据安全


项目地址：https://github.com/noovertime7/gin-mysqlbak

前端地址: https://github.com/noovertime7/gin-mysqlbak/tree/main/front

agent地址: https://github.com/noovertime7/gin-mysqlbak-agent
## 一、现在开始
部署使用请查阅 https://noovertime7.github.io/mysqlbak-website/pages/help/start.html
## 二、功能演示
 首页大屏

![dashboard](http://qiniu.yunxue521.top/mysqlbak/dashboard.jpg)
 服务管理

agent部署成功后，会向server端发起注册，注册成功服务信息保存在服务管理页面

![服务管理](http://qiniu.yunxue521.top/mysqlbak/service_list.jpg)
 应用管理
当前服务下所有应用，包括mysql与elastic应用

![应用管理](http://qiniu.yunxue521.top/mysqlbak/app.jpg)
 任务管理
当前服务下所有任务，包括mysql与elastic任务

![任务管理](http://qiniu.yunxue521.top/mysqlbak/task.jpg)
 任务总览
任务总览界面包含集群内所有注册服务的任务列表，可以很方便的启动停止任务，同步周期取决与配置文件中的clusterSyncPeriod

![任务总览](http://qiniu.yunxue521.top/mysqlbak/task_overview.jpg)
 历史记录
包括mysql历史记录与elastic快照记录

![历史记录](http://qiniu.yunxue521.top/mysqlbak/history.jpg)

[更多功能演示](https://noovertime7.github.io/mysqlbak-website/pages/show/show.html)

## 三、FAQ

1.目前的工作方向？

前端重构完成后，发布3.0.0版本，未来会持续优化系统

## 四、其他
这个项目的前端后端包括agent都是我一个人完成，所以会有很多不足，希望能提issue，或者提PR一起来完善这个项目
