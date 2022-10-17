## 登录平台
::: tip 提示
登录地址的端口取决于前端服务映射的端口
:::
前端后端部署完成后，登录系统，默认密码为 admin/admin@123 

**本次演示环境为集群环境，本地环境后面要去掉**
## 服务列表
agent部署成功后，会向server端发起注册，注册成功服务信息保存在服务管理页面

![服务列表](https://github.com/noovertime7/gin-mysqlbak/blob/main/img/service_list.jpg?raw=true)
## 添加应用
服务成功注册后，需要先添加应用才能添加任务，可选择添加mysql主机或elastic主机

![添加应用](https://github.com/noovertime7/gin-mysqlbak/blob/main/img/add_app.jpg?raw=true)

## 添加任务
应用添加成功后，点击服务列表中的任务，给该服务添加备份任务，如需钉钉通知或保存到oss，请打开对应开关并填写所需信息

![添加任务](https://github.com/noovertime7/gin-mysqlbak/blob/main/img/add_task.jpg?raw=true)

## 任务总览
任务总览界面包含集群内所有注册服务的任务列表，可以很方便的启动停止任务，同步周期取决与配置文件中的clusterSyncPeriod

![任务总览](https://github.com/noovertime7/gin-mysqlbak/blob/main/img/task_overview.jpg?raw=true)