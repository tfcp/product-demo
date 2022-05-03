#!/bin/bash

# 创建日志目录
mkdir -p $LOG_PATH

# 修改目录权限
chmod 777 -R $LOG_PATH

exec $APP_PATH/server $1 $2 $3
