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
	//  端口监听
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Printf("Failed to start server, err:%s", err.Error())
	}
}
