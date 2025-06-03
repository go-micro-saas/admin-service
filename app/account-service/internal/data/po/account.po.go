package po

import (
	schemas "github.com/go-micro-saas/account-service/app/account-service/internal/data/schema/user"
	gormpkg "github.com/ikaiguang/go-srv-kit/data/gorm"
	"gorm.io/gorm"
)

type CreateAccountParam struct {
	UserModel     *User
	RegPhoneModel *UserRegPhone
	RegEmailModel *UserRegEmail
}

type QueryUserParam struct {
	UidList      []uint64
	ContactPhone string
	ContactEmail string

	PaginatorArgs *gormpkg.PaginatorArgs
}

func (s *QueryUserParam) WhereConditions(dbConn *gorm.DB) *gorm.DB {
	if len(s.UidList) == 1 {
		dbConn = dbConn.Where(schemas.UserSchema.TableName()+"."+schemas.FieldUserId+" = ?", s.UidList[0])
	} else if len(s.UidList) > 1 {
		dbConn = dbConn.Where(schemas.UserSchema.TableName()+"."+schemas.FieldUserId+" IN (?)", s.UidList)
	}
	if s.ContactPhone != "" {
		dbConn = dbConn.Where(schemas.UserSchema.TableName()+"."+schemas.FieldUserPhone+" = ?", s.ContactPhone)
	}
	if s.ContactEmail != "" {
		dbConn = dbConn.Where(schemas.UserSchema.TableName()+"."+schemas.FieldUserEmail+" = ?", s.ContactEmail)
	}
	return dbConn
}
