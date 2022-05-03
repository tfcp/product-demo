# product-demo
grpc project template...
快速搭建的grpc项目模板, 根据自身习惯持续更新...

## 项目介绍
    
    1.golangci-lint
        
        golang代码规范检查组件
        文档: https://golangci-lint.run/usage/install/
        
    2. kafka
    
    3. grpc
    
        rpc协议(需要注意协议版本v2,v3不同)
        1. 安装
        参考文档:(https://zhuanlan.zhihu.com/p/501542023)
        2. 使用(Makefile)
        make go
        
    4. grpcui
    
       基于反射的grpc调试工具
       安装:
        go get github.com/fullstorydev/grpcui
        go install github.com/fullstorydev/grpcui/cmd/grpcui
       mac:
        sudo cp `go env|grep 'GOPATH'|sed -e 's/GOPATH="//' -e 's/"//'`/bin/grpcui /usr/local/bin/
        chmod +x /usr/local/bin/grpcui 
        
       调试使用(仅限本地):
        1. make server # 设置本地调试模式
        2. grpcui -plaintext 127.0.0.1:8000
    
## 项目结构

    |--app         
        |--grpc             grpc协议文件目录
    |--internal
        |--middleware       中间件目录
        |--enum             常量
        |--model            模型
        |--service          服务
        |--client           外部请求(微服务)
        |--test             单元测试
    |--config
    |--bootstrap            启动目录
    |--scripts              脚本目录
    |--library
    |--logs
    |--Makefile             快捷命令
    |--main.go

## 常用命令
    1. make go
        加载新协议

    2. make clean
        清理所有协议文件

    3. make server
        本地程序调试
