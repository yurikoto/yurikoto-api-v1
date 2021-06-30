# Yurikoto第一版API实现（基于Golang）

> 使用或自部署本代码请严格遵守AGPL-3.0协议。本仓库代码可供学习使用，但Yurikoto提供的壁纸、台词资源严禁商用

[![Open Source Helpers](https://www.codetriage.com/yurikoto/yurikoto-api-v1/badges/users.svg)](https://www.codetriage.com/yurikoto/yurikoto-api-v1) [![Go Report Card](https://goreportcard.com/badge/github.com/yurikoto/yurikoto-api-v1)](https://goreportcard.com/report/github.com/yurikoto/yurikoto-api-v1) [![Maintainability](https://api.codeclimate.com/v1/badges/1ef898f65c8c593baf49/maintainability)](https://codeclimate.com/github/yurikoto/yurikoto-api-v1/maintainability) [![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fyurikoto%2Fyurikoto-api-v1.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fyurikoto%2Fyurikoto-api-v1?ref=badge_shield)

## 简介

通过go-gin实现的Yurikoto第一版API，如果您有任何建议或改进想法，欢迎提交issue或pr。

[Yurikoto主页](https://yurikoto.com)

## 自部署指南

### 环境

以下为Yurikoto官方使用的环境
CentOS 8.2
Nginx 1.17
Go 1.15.7 （服务端可不需要）
MySQL 5.7
Redis 6.0.10

### 数据库

本项目使用gorm进行数据库操作，无需提前创建表等。如需要使用Yurikoto官方数据库，请提交issue并留下您的邮箱以获取数据库备份。

### 配置文件

根据文件内提示修改`config.template.ini`并重命名为`config.ini`

### 编译

```shell
go mod download
SET GOOS=linux
SET GOARCH=amd64
go build
```

执行上述命令后项目目录会出现可执行文件`yurikoto-api-go-v1`

### 运行

将可执行文件与`config.ini`与`favicon.ico`上传到服务器。在`lib/systemd/system`下创建`yurikoto-api-v1.service`，内容如下：

```ini
[Unit]
Description=Yurikoto API V1 Service
After=network.target

[Service]
Type=simple
Restart=on-failure
RestartSec=5s
ExecStart=/path/yurikoto-api-go-v1 -c /path/


[Install]
WantedBy=multi-user.target
```

并将`path`替换为可执行文件所在目录（`ExecStart`末尾斜杠需保留）。执行：

```shell
systemctl enable yurikoto-api-v1
systemctl start yurikoto-api-v1
systemctl status yurikoto-api-v1
```

若显示"active"字样，则说明部署成功。

### 反向代理
```nginx
location /
{
    proxy_pass http://127.0.0.1:3417;
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header REMOTE-HOST $remote_addr;
    
    add_header Access-Control-Allow-Origin *;
    add_header X-Cache $upstream_cache_status;
    
    #Set Nginx Cache
    
    	add_header Cache-Control no-cache;
    expires 12h;
}
```