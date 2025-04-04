// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/patient-fyd/jxust-softhub-api/internal/dao/internal"
)

// internalMembersDao is internal type for wrapping internal DAO implements.
type internalMembersDao = *internal.MembersDao

// membersDao is the data access object for table members.
// You can define custom methods on it to extend its functionality as you wish.
type membersDao struct {
	internalMembersDao
}

var (
	// Members is globally public accessible object for table members operations.
	Members = membersDao{
		internal.NewMembersDao(),
	}
)

// Fill with you ideas below.
