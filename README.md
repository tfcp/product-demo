# product-demo
product and gin project template...
根据自身习惯持续更新...

## 项目介绍
1.go-jwt
    
    鉴权组件(不支持删除)

2.vue-element-admin
    
    统一后台前端解决方案
    文档: https://panjiachen.github.io/vue-element-admin-site/zh/guide

3.Element-Ui
    
    vue2.0 前端组件库
    文档: https://element.eleme.cn/#/zh-CN/component/installation
    
4.golangci-lint
    
    golang代码规范检查组件
    文档: https://golangci-lint.run/usage/install/
    
5.go.rice
    
    静态资源项目编译到二进制文件中(dist)
    (不用单独部署 小项目专用 大业务 微服务禁用该组件)

6.gin-swagger
    
    api文档: https://github.com/swaggo/gin-swagger
    demo: https://github.com/swaggo/gin-swagger/blob/master/example/basic/main.go
    
7. kafka

8. grpc
    
    微服务协议

9. endless

      [优雅重启](#https://www.bookstack.cn/read/gin-EDDYCJY-blog/golang-gin-2018-03-15-Gin%E5%AE%9E%E8%B7%B5-%E8%BF%9E%E8%BD%BD%E4%B8%83-Golang%E4%BC%98%E9%9B%85%E9%87%8D%E5%90%AFHTTP%E6%9C%8D%E5%8A%A1.md)

## 项目结构

    |--app         
        |--api              任务目录
        |--cron             任务目录
        |--process          进程目录
    |--internal
        |--middleware       常量
        |--enum             中间件目录
        |--model            模型
        |--service          服务
        |--client           外部请求(微服务)
    |--config
    |--bootstrap
    |--docs                 接口文档
    |--library
    |--logs
    |--router
    |--web
    |--Makefile             快捷命令
    |--main.go

## 常用命令
1. make build
编译项目

2. make run
运行web项目

3. make web
vue项目npm run dev 
