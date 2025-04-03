package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// UploadReq 上传文件请求
type UploadReq struct {
	g.Meta      `path:"/api/file/upload" method:"post" mime:"multipart/form-data" tags:"FileService" summary:"上传文件"`
	File        *ghttp.UploadFile `p:"file" v:"required#请选择上传文件"` // 上传的文件对象
	RelatedType string            `p:"related_type" dc:"关联类型，如news、activity、resource等"`
	RelatedId   uint              `p:"related_id" dc:"关联ID，指向具体的内容ID"`
}

// UploadRes 上传文件响应
type UploadRes struct {
	FileId   uint   `json:"fileId"`   // 文件ID
	FileName string `json:"fileName"` // 文件名
	FileSize int64  `json:"fileSize"` // 文件大小
	FileType string `json:"fileType"` // 文件MIME类型
	FilePath string `json:"filePath"` // 文件存储路径/URL
}

// ListReq 文件列表请求
type ListReq struct {
	g.Meta      `path:"/api/file/list" method:"get" tags:"FileService" summary:"获取文件列表"`
	RelatedType string `p:"related_type" dc:"关联类型，如news、activity、resource等"`
	RelatedId   uint   `p:"related_id" dc:"关联ID，指向具体的内容ID"`
	// 分页参数
	PageNum  int `p:"page_num" v:"min:0#分页号码错误" dc:"分页号码, 默认1"`
	PageSize int `p:"page_size" v:"max:500#每页数量最大500条" dc:"分页数量, 最大500"`
}

// ListRes 文件列表响应
type ListRes struct {
	List     []FileInfo `json:"list"`      // 文件列表
	Total    int        `json:"total"`     // 总数
	PageNum  int        `json:"page_num"`  // 分页号码
	PageSize int        `json:"page_size"` // 分页数量
}

// FileInfo 文件信息
type FileInfo struct {
	FileId       uint   `json:"fileId"`       // 文件ID
	FileName     string `json:"fileName"`     // 文件名
	FileSize     int64  `json:"fileSize"`     // 文件大小
	FileType     string `json:"fileType"`     // 文件MIME类型
	FilePath     string `json:"filePath"`     // 文件存储路径/URL
	UploaderId   uint   `json:"uploaderId"`   // 上传者ID
	UploaderName string `json:"uploaderName"` // 上传者姓名
	UploadTime   string `json:"uploadTime"`   // 上传时间
	RelatedType  string `json:"relatedType"`  // 关联类型
	RelatedId    uint   `json:"relatedId"`    // 关联ID
}

// DeleteReq 删除文件请求
type DeleteReq struct {
	g.Meta `path:"/api/file/delete" method:"delete" tags:"FileService" summary:"删除文件"`
	FileId uint `p:"fileId" v:"required#文件ID不能为空"`
}

// DeleteRes 删除文件响应
type DeleteRes struct{}

// DownloadReq 下载文件请求
type DownloadReq struct {
	g.Meta `path:"/api/file/download/{fileId}" method:"get" tags:"FileService" summary:"下载文件"`
	FileId uint `p:"fileId" v:"required#文件ID不能为空" in:"path"`
}

// DownloadRes 下载文件响应 - 此处不需要定义，框架会自动处理文件下载响应
type DownloadRes struct{}
