// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FilesDao is the data access object for table files.
type FilesDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns FilesColumns // columns contains all the column names of Table for convenient usage.
}

// FilesColumns defines and stores column names for table files.
type FilesColumns struct {
	FileId      string // 文件ID，主键，自增
	FileName    string // 文件名称
	FileSize    string // 文件大小（字节）
	FileType    string // 文件类型/MIME类型
	FilePath    string // 文件存储路径
	UploaderId  string // 上传者ID，关联users表
	UploadTime  string // 上传时间
	Md5Hash     string // 文件MD5哈希值，用于去重
	RelatedType string // 关联类型，如news、activity、resource等
	RelatedId   string // 关联ID，指向具体的内容ID
}

// filesColumns holds the columns for table files.
var filesColumns = FilesColumns{
	FileId:      "fileId",
	FileName:    "fileName",
	FileSize:    "fileSize",
	FileType:    "fileType",
	FilePath:    "filePath",
	UploaderId:  "uploaderId",
	UploadTime:  "uploadTime",
	Md5Hash:     "md5Hash",
	RelatedType: "relatedType",
	RelatedId:   "relatedId",
}

// NewFilesDao creates and returns a new DAO object for table data access.
func NewFilesDao() *FilesDao {
	return &FilesDao{
		group:   "default",
		table:   "files",
		columns: filesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *FilesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *FilesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *FilesDao) Columns() FilesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *FilesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *FilesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *FilesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
