package meta

// FileMeta : 文件元信息结构
type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

// 定义一个map存放文件云信息
var fileMetas map[string]FileMeta

// 用于初始化文件元信息
func init() {
	fileMetas = make(map[string]FileMeta)
}

// UpdateFileMeta : 新增/更新文件元信息
func UpdateFileMeta(fmeta FileMeta) {
	fileMetas[fmeta.FileSha1] = fmeta
}

// GetFileMeta : 通过sha1值获取文件的元信息对象
func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}
