package main

import (
	"filestore-server/handler"
	"fmt"
	"net/http"
)

func main() {
	// 定义文件上传接口路由
	http.HandleFunc("/file/upload", handler.FileUploadHandler)
	// 定义文件上传成功接口路由
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
	// 定义查询文件云信息接口路由
	http.HandleFunc("/file/meta", handler.GetFileMetaHandler)
	// 定义文件下载接口路由
	http.HandleFunc("/file/download", handler.DownloadHandler)
	// 定义文件重命名接口路由
	http.HandleFunc("/file/update", handler.FileMetaUpdateHandler)
	// 定义文件重命名接口路由
	http.HandleFunc("/file/delete", handler.FileDeleteHandler)

	// 用户注册
	http.HandleFunc("/user/signup", handler.SignUpHandler)
	// 用户登录
	http.HandleFunc("/user/signin", handler.SignInHandler)

	//  端口监听
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Printf("Failed to start server, err:%s", err.Error())
	}
}
