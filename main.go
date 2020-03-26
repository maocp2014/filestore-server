package main

import (
	"filestore-server/handler"
	"fmt"
	"net/http"
)

func main() {
	// 定义文件上传接口路由
	http.HandleFunc("/file/upload", handler.FileUploadHandler)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Printf("Failed to start server, err:%s", err.Error())
	}
}
