# SearchEngine3090ti

2022字节跳动青训营搜索引擎项目

<br>

### Usage
基本环境配置
- go 1.18
- Docker
- kitex
- protobuf(可选)

<br>

命令行下
> 1.设置基本依赖服务
>> docker-compose up
>>
>2.运行search RPC服务器服务
>>cd cmd/search
>>sh build.sh
>>sh output/bootstrap.sh
>>
>3.运行User RPC 服务器服务
>>cd cmd/user
>>sh build.sh
>>sh output/bootstrap.sh
>>
>4.运行 API 服务器服务
>>cd cmd/api
>>chmod +x run.sh
>>./run.sh
>>

<br>

### 内部开发指南
//TODO