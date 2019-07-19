# Golang API 范例项目

## 目的

本项目采用了一系列Golang中比较流行的组件，可以以本项目为基础快速搭建Restful Web API

## 特色

本项目已经整合了许多开发API所必要的组件：

1. [Gin](https://github.com/gin-gonic/gin): 轻量级Web框架，自称路由速度是golang最快的 
2. [GORM](http://gorm.io/docs/index.html): ORM工具。本项目需要配合Mysql使用 
3. [Gin-Session](https://github.com/gin-contrib/sessions): Gin框架提供的Session操作工具
4. [Go-Redis](https://github.com/go-redis/redis): Golang Redis客户端
5. [Gin-Cors](https://github.com/gin-contrib/cors): Gin框架提供的跨域中间件
6. 自行实现了国际化i18n的一些基本功能
7. 本项目是使用基于cookie实现的session来保存登录状态的，如果需要可以自行修改为token验证

本项目已经预先实现了一些常用的代码方便参考和复用:

## 用户基础模型API

1. 实现了 `/api/v1/user/register` 用户注册接口
2. 实现了 `/api/v1/user/login` 用户登录接口
3. 实现了 `/api/v1/user/me` 用户资料接口(需要登录后获取 Session)
4. 实现了 `/api/v1/user/logout` 用户登出接口(需要登录后获取 Session)

## 视频基础模型API

1. 实现了 `/api/v1/videos` 视频投稿(post),视频列表(get)
2. 实现了 `/api/v1/video:id` 视频更新(post),视频详情(get),删除视频(delete)

本项目已经预先创建了一系列文件夹划分出下列模块:

1. `api` 文件夹就是MVC框架的 Controller，负责协调各部件完成任务
2. `model` 文件夹负责存储数据库模型和数据库操作相关的代码
3. `service` 负责处理比较复杂的业务，把业务代码模型化可以有效提高业务代码的质量（比如用户注册，充值，下单等）
4. `serializer` 储存通用的 JSON 模型，把 `model` 得到的数据库模型转换成 API 需要的 JSON 对象
5. `cache` 负责 Redis 缓存相关的代码
6. `auth` 权限控制文件夹
7. `util` 一些通用的小工具
8. `conf` 放一些静态存放的配置文件，其中 `locales` 内放置翻译相关的配置文件

## Godotenv

项目在启动的时候依赖以下环境变量，但是在也可以在项目根目录创建 `.env` 文件设置环境变量便于使用（建议开发环境使用）

```shell
MYSQL_DSN="db_user:db_password@/db_name?charset=utf8&parseTime=True&loc=Local" # Mysql连接地址
REDIS_ADDR="127.0.0.1:6379" # Redis端口和地址
REDIS_PW="" # Redis连接密码
REDIS_DB="" # Redis库从0到10
SESSION_SECRE="" # Seesion密钥，必须设置而且不要泄露
GIN_MODE="debug"
```

## Go Mod

本项目使用[Go Mod](https://github.com/golang/go/wiki/Modules)管理依赖。

```shell
go mod init pName
export GOPROXY=http://mirrors.aliyun.com/goproxy/
go run main.go // 自动安装
```

## 运行

```shell
go run main.go
```

项目运行后启动在3000端口（可以修改，参考gin文档)
