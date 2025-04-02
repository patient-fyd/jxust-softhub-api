package model

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
)

// FileUploadInput 上传文件输入
type FileUploadInput struct {
	File        *ghttp.UploadFile // 上传的文件对象
	UploaderId  uint              // 上传者ID
	RelatedType string            // 关联类型，如news、activity、resource等
	RelatedId   uint              // 关联ID，指向具体的内容ID
}

// FileUploadOutput 上传文件输出
type FileUploadOutput struct {
	FileId   uint   // 文件ID
	FileName string // 文件名
	FileSize int64  // 文件大小
	FileType string // 文件MIME类型
	FilePath string // 文件存储路径/URL
}

// FileListInput 文件列表查询输入
type FileListInput struct {
	RelatedType string // 关联类型，如news、activity、resource等
	RelatedId   uint   // 关联ID，指向具体的内容ID
	PageNum     int    // 分页号码
	PageSize    int    // 分页数量
}

// FileListOutput 文件列表查询输出
type FileListOutput struct {
	List     []*FileInfo // 文件列表
	Total    int         // 总数
	PageNum  int         // 分页号码
	PageSize int         // 分页数量
}

// FileInfo 文件信息
type FileInfo struct {
	FileId       uint        // 文件ID
	FileName     string      // 文件名
	FileSize     int64       // 文件大小
	FileType     string      // 文件MIME类型
	FilePath     string      // 文件存储路径/URL
	UploaderId   uint        // 上传者ID
	UploaderName string      // 上传者姓名
	UploadTime   *gtime.Time // 上传时间
	RelatedType  string      // 关联类型
	RelatedId    uint        // 关联ID
	Md5Hash      string      // 文件MD5哈希值
}

// FileDeleteInput 删除文件输入
type FileDeleteInput struct {
	FileId uint // 文件ID
}

// FileDeleteOutput 删除文件输出
type FileDeleteOutput struct{}

// FileDownloadInput 下载文件输入
type FileDownloadInput struct {
	FileId uint // 文件ID
}
