// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/patient-fyd/jxust-softhub-api/internal/dao/internal"
)

// internalContentTagsDao is internal type for wrapping internal DAO implements.
type internalContentTagsDao = *internal.ContentTagsDao

// contentTagsDao is the data access object for table content_tags.
// You can define custom methods on it to extend its functionality as you wish.
type contentTagsDao struct {
	internalContentTagsDao
}

var (
	// ContentTags is globally public accessible object for table content_tags operations.
	ContentTags = contentTagsDao{
		internal.NewContentTagsDao(),
	}
)

// Fill with you ideas below.
