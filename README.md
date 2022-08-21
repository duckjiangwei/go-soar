# hk591_go 介绍
这是一个基于 Gin 框架改造的 web 框架

# 部署
1. 初次部署或检测到 go.sum 文件有更新，下载相关包1
```
go mod tidy
```
2. 环境变量文件
```
cp .env.production .env
```
3. 生成可执行文件
```
go build main.go
```
4. 执行
```
./main
```