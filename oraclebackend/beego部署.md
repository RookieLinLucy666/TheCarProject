# beego后端部署启动

## 环境配置
首先确认配置好go环境，设置好GOROOT与GOPATH，并安装好git环境
```bigquery
go get github.com/astaxie/beego
go get github.com/beego/bee #安装bee工具
```
命令行中输入bee指令，得到结果即为配置好环境

## 数据库配置
```bigquery
#登陆数据库
mysql -uroot -p
#创建库并使用
create database the_car_project;
use the_car_project;
source /yourpath/TheCarProject/oraclebackend/the_car_project.sql;
```
在oraclebackend/main.go中找到代码段
```bigquery
orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/the_car_project?charset=utf8")
```
将第三个参数设置为自身的数据库信息

## 启动项目
```bigquery
#移动到指定目录下
cd /yourpath/TheCarProject/oraclebackend
#使用bee工具开启项目
bee run
```