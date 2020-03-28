# 一、filestore-server分布式网盘

## 概况
[1.1 课程导学](./docs/1.1、课程导学.md)

[1.2 课程介绍](./docs/1.2、课程介绍.md)

## 精简版云盘

开发环境准备等; 接口逻辑的实现，包括上传文件，下载文件，查询以及更改文件元信息等功能; 结合Postman进行接口测试

[2.1 "云存储"系统原型之简单文件上传服务架构说明](./docs/2.1、云存储系统原型之简单文件上传服务架构说明.md)

[2.2 编码实战-实现上传接口](./docs/2.2、实现上传接口.md)

[2.3 编码实战-保存文件元信息](./docs/2.3、保存文件元信息.md)

[2.4 编码实战-单个文件信息查询接口](./docs/2.4、单个文件信息查询接口.md)

[2.5 编码实战-实现文件下载接口](./docs/2.5、实现文件下载接口.md)

[2.6 编码实战-实现文件修改和删除接口](./docs/2.6、实现文件修改和删除接口.md)

## 系统架构升级

"云存储"系统之基于MySQL实现的文件数据库【持久化云文件信息】

系统架构升级说明; MySQL基于读写分离的主从原理及实战部署; MySQL表字段设计及基于海量数据的水平分表; Go管理MySQL, 实现文件metaData的持久化

[3.1 MySQL基础知识](./docs/3.1、MySQL基础知识.md)

[3.2 MySQL主从数据同步演示](./docs/3.2、MySQL主从数据同步演示.md)

[3.3 文件表的设计及创建](./docs/3.3、文件表的设计及创建.md)

[3.4 编码实战-持久化元数据到文件表](./docs/3.4、持久化元数据到文件表.md)

[3.5 编码实战-从文件表中获取元数据](./docs/3.5、从文件表中获取元数据.md)

# 二、运行

## 方式一
```bash
# 编译运行,会生成一个二进制可运行包main
go build main.go   #后面指定文件的build会生成对应前缀的二进制包例如main
./main

go build           #后面不接文件build，会生成项目名的二进制包例如filestore-server
./filestore-server

# 访问
http://127.0.0.1:8090/file/upload
```

## 方式二
```bash
# 直接命名运行
go run main.go

# 访问
http://127.0.0.1:8090/file/upload
```

# 参考资料：

https://www.bilibili.com/video/av83350170?p=7

https://github.com/Lancger/file-storage-system

https://www.jianshu.com/p/2bf1ddf9fba1  Go实战仿百度云盘 实现企业级分布式云存储系统
