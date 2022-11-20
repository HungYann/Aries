[![Go](https://github.com/LIU-HONGYANG/Aries/actions/workflows/go.yml/badge.svg)](https://github.com/LIU-HONGYANG/Aries/actions/workflows/go.yml)

# Aries

## 简介

Aries是连接第三方接口和自动仓的代理服务器。具备转发第三方请求，并对请求结果进行处理的功能。

## 功能

![](https://tva1.sinaimg.cn/large/e6c9d24egy1h6fddq09tzj20sg0t275d.jpg)
业务处理能力如上时序图，其中代理服务器通过golang实现，持久化存储通过redis实现，配置服务，通过viper+yml方式实现，日志通过golang日志库配置记录
与此同时，在`转发请求`和`响应请求`阶段， 可以根据用户定义的各种参数多次尝试

## 结构

| 项目名称                 | 描述                                         |
| ------------------------ | -------------------------------------------- |
| build         | 项目容器化打包预留设计 |
| cmd          | 项目入口，项目从main开始编译和运行 |
| config   | 项目静态配置文件    |
| internal | 拆分的各类私有包，和cmd结构对应                             |

## 依赖

* go 1.17
* go mod

## 编译

```shell
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o Aries cmd/aries/main.go  #编译linux目标文件
```

## 发布

scp方式上传到线上虚拟机环境部署
