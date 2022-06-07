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
>> docker-compose up / sudo docker-compose up
>>
>2.运行search RPC服务器服务
>>cd cmd/search
>>sh build.sh
>>sh output/bootstrap.sh / sudo sh output/bootstrap.sh
>>
>3.运行User RPC 服务器服务
>>cd cmd/user
>>sh build.sh
>>sh output/bootstrap.sh / sudo sh output/bootstrap.sh
>>
>4.运行 API 服务器服务
>>cd cmd/api
>>chmod +x run.sh
>>./run.sh
>>

<br>

### 内部开发指南
####收藏夹接口
|  描述             |      URL           |     方法    |    传参      |      接收 |
| -----------       | -----------        | ----------- | ----------- | ----------- |
| 获取用户收藏夹列表 | /collection/       |        GET      |          |      {...,"name":[colltName1,colltName2...],"colltId":[colltId1,colltId2...]}      |
| 获取收藏夹详情   | /collection/:collt        | GET | Query:?collt=id | {...,"name":colltName,"entry":[entry1,entry2...]} |
| 创建收藏夹 | /collection/create | POST | POSTFORM:key=name,value=? | {...} |
| 删除收藏夹 | /collection/delete:collt | GET | Query:?collt=id | {...} |
| 添加搜索记录于指定收藏夹 | /collection/:collt/add | POST | Query:?collt=id ; POSTFORM:key=newentry,value=? | {...} |
| 删除指定收藏夹中指定搜索记录 | /collection/:collt/delete| GET | Query:?collt=id POSTFROM:key=entry,value=? | {...} |
| 设置指定收藏夹名字 | /collection/:collt/setname | POST |Query:?collt=id ; POSTFROM:key=newname,value=? | {...} |
//TODO