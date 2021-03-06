package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"filestore-server/meta"
	"filestore-server/util"
)

// FileUploadHandler : 处理文件上传
func FileUploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// 返回上传 HTML 界面
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			io.WriteString(w, "internal server error")
			return
		}
		io.WriteString(w, string(data))
	} else if r.Method == "POST" {
		// 接收文件流及存储到本地目录
		file, head, err := r.FormFile("file")
		if err != nil {
			fmt.Printf("Failed to get data, err:%s\n", err.Error())
			return
		}
		// 关闭文件句柄
		defer file.Close()

		// 文件云信息保存
		fileMeta := meta.FileMeta{
			FileName: head.Filename,
			Location: "/tmp/" + head.Filename,
			UploadAt: time.Now().Format("2006-01-02 15:04:05"),
		}

		// newFile, err := os.Create("/tmp/" + head.Filename)
		newFile, err := os.Create(fileMeta.Location)

		if err != nil {
			fmt.Printf("Failed to create file, err:%s\n", err.Error())
			return
		}
		// 关闭文件句柄
		defer newFile.Close()

		// _, err = io.Copy(newFile, file)
		fileMeta.FileSize, err = io.Copy(newFile, file)

		if err != nil {
			fmt.Printf("Failed to save data into file, err:%s\n", err.Error())
			return
		}

		// 游标重新回到文件头部
		newFile.Seek(0, 0)
		fileMeta.FileSha1 = util.FileSha1(newFile)

		// meta.UpdateFileMeta(fileMeta)
		_ = meta.UpdateFileMetaDB(fileMeta)

		// 上传成功重定向
		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}
}

// UploadSucHandler : 文件上传成功接口
func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Upload finished!")
}

// GetFileMetaHandler : 获取文件元信息
func GetFileMetaHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// 获取文件的filehash
	filehash := r.Form["filehash"][0]
	// 从程序内存中获取文件元信息
	// fMeta := meta.GetFileMeta(filehash)

	// 从数据库中获取文件元信息
	fMeta, err := meta.GetFileMetaDB(filehash)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(fMeta)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

// DownloadHandler : 文件下载接口
func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fsha1 := r.Form.Get("filehash")
	fm := meta.GetFileMeta(fsha1)
	// TODO：加载已存储到云端本地的文件内容，并返回到客户端
	f, err := os.Open(fm.Location)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// 关闭文件句柄
	defer f.Close()

	// 将内容读取出来,使用ReadAll方法全部加载到内存里，这里都是小文件，所以可以这么操作，如果是大文件，则需要使用流的方式实现
	data, err := ioutil.ReadAll(f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/octect-stream")
	// attachment表示文件将会提示下载到本地，而不是直接在浏览器中打开
	w.Header().Set("content-disposition", "attachment; filename=\""+fm.FileName+"\"")
	w.Write(data)
}

// FileMetaUpdateHandler : 更新元信息接口（重命名）
func FileMetaUpdateHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	opType := r.Form.Get("op")
	FileSha1 := r.Form.Get("filehash")
	newFileName := r.Form.Get("filename")

	// op 为0表示重命名操作
	if opType != "0" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	// 判断操作类型是否为POST
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	curFileMeta := meta.GetFileMeta(FileSha1)
	curFileMeta.FileName = newFileName
	meta.UpdateFileMeta(curFileMeta)

	// TODO: 更新文件表中的元信息记录
	data, err := json.Marshal(curFileMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// FileDeleteHandler : 删除文件及元信息
func FileDeleteHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fileSha1 := r.Form.Get("filehash")
	fMeta := meta.GetFileMeta(fileSha1)

	// 删除文件
	os.Remove(fMeta.Location)
	// 删除文件元信息
	meta.RemoveFileMeta(fileSha1)

	w.WriteHeader(http.StatusOK)
}
