# SearchEngine3090ti


## Introduction
https://kpcbf4ul2l.feishu.cn/docx/doxcnh8E8ZdonIeLizsyJeQFmPA

## API Document
https://www.apifox.cn/apidoc/shared-f5f40433-908c-44e8-9661-e5ed84988a59

## Quick Start
Environment:
- go 1.18
- Docker
- kitex
- protobuf

1.设置基本依赖服务
> sudo docker-compose up

2.运行search RPC服务器服务
>cd cmd/search
>sh build.sh
>sudo sh output/bootstrap.sh

3.运行User RPC 服务器服务
>cd cmd/user
>sh build.sh
>sudo sh output/bootstrap.sh

4.运行Collection 服务器服务
>cd cmd/collection
>sh build.sh
>sudo sh output/bootstrap.sh

5.运行 API 服务器服务
>cd cmd/api
>chmod +x run.sh
>./run.sh
