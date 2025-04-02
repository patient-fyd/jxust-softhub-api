package service

import (
	"context"

	"github.com/patient-fyd/jxust-softhub-api/internal/model"
)

// IFile 文件管理服务接口
type IFile interface {
	// Upload 上传文件
	Upload(ctx context.Context, in model.FileUploadInput) (*model.FileUploadOutput, error)

	// List 获取文件列表
	List(ctx context.Context, in model.FileListInput) (*model.FileListOutput, error)

	// Delete 删除文件
	Delete(ctx context.Context, in model.FileDeleteInput) (*model.FileDeleteOutput, error)

	// Download 下载文件
	Download(ctx context.Context, in model.FileDownloadInput) (string, error)
}

var (
	localFile IFile
)

// File 获取文件管理服务实例
func File() IFile {
	if localFile == nil {
		panic("implement not found for interface IFile, forgot register?")
	}
	return localFile
}

// RegisterFile 注册文件管理服务实现
func RegisterFile(i IFile) {
	localFile = i
}
