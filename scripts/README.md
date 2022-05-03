## 脚本目录

### 文件列表
    |--scripts  
        |--start.sh     启动脚本

### 问题
    1. 收不到sigterm信号,导致平滑关闭无效
        因为dockerfile cmd走shell的话, 我们的容器主进程就是shell了 而我们的业务进程就变成子进程了 需要特殊处理才能收到sigterm信号
      否则永远都是默认的30秒被sigkill强制删除掉
