# 一、filestore-server分布式网盘

[1.1 课程导学](./docs/1.1、课程导学.md)

[1.2 课程介绍](./docs/1.2、课程介绍.md)

[2.1 "云存储"系统原型之简单文件上传服务架构说明](./docs/2.1、云存储系统原型之简单文件上传服务架构说明.md)

[2.2 编码实战-实现上传接口](./docs/2.2、实现上传接口.md)

[2.3 编码实战-保存文件元信息](./docs/2.3、保存文件元信息.md)

[2.4 编码实战-单个文件信息查询接口](./docs/2.4、单个文件信息查询接口.md)

[2.5 编码实战-实现文件下载接口](./docs/2.5、实现文件下载接口.md)

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
