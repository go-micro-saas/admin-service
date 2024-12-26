// Package data
// Code generated by ikaiguang. <https://github.com/ikaiguang>
package data

import (
	"bytes"
	context "context"
	"database/sql"
	"github.com/go-kratos/kratos/v2/log"
	errorv1 "github.com/go-micro-saas/admin-service/api/admin-service/v1/errors"
	"github.com/go-micro-saas/admin-service/app/admin-service/internal/data/po"
	datarepos "github.com/go-micro-saas/admin-service/app/admin-service/internal/data/repo"
	schemas "github.com/go-micro-saas/admin-service/app/admin-service/internal/data/schema/user_reg_email"
	gormpkg "github.com/ikaiguang/go-srv-kit/data/gorm"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	gorm "gorm.io/gorm"
	"strings"
)

// userRegEmailDataRepo repo
type userRegEmailDataRepo struct {
	log                *log.Helper
	dbConn             *gorm.DB             // *gorm.DB
	UserRegEmailSchema schemas.UserRegEmail // UserRegEmail
}

// NewUserRegEmailDataRepo new data repo
func NewUserRegEmailDataRepo(logger log.Logger, dbConn *gorm.DB) datarepos.UserRegEmailDataRepo {
	logHelper := log.NewHelper(log.With(logger, "module", "admin-service/data/user_reg_email"))
	return &userRegEmailDataRepo{
		log:    logHelper,
		dbConn: dbConn,
	}
}

func (s *userRegEmailDataRepo) NewTransaction(ctx context.Context, opts ...*sql.TxOptions) gormpkg.TransactionInterface {
	return gormpkg.NewTransaction(ctx, s.dbConn, opts...)
}

// =============== 创建 ===============

// create insert one
func (s *userRegEmailDataRepo) create(ctx context.Context, dbConn *gorm.DB, dataModel *po.UserRegEmail) (err error) {
	err = dbConn.WithContext(ctx).
		Table(s.UserRegEmailSchema.TableName()).
		Create(dataModel).Error
	if err != nil {
		e := errorpkg.ErrorInternalServer("")
		return errorpkg.Wrap(e, err)
	}
	return
}

// Create insert one
func (s *userRegEmailDataRepo) Create(ctx context.Context, dataModel *po.UserRegEmail) error {
	return s.create(ctx, s.dbConn, dataModel)
}

// CreateWithDBConn create
func (s *userRegEmailDataRepo) CreateWithDBConn(ctx context.Context, dbConn *gorm.DB, dataModel *po.UserRegEmail) error {
	return s.create(ctx, dbConn, dataModel)
}

func (s *userRegEmailDataRepo) CreateWithTransaction(ctx context.Context, tx gormpkg.TransactionInterface, dataModel *po.UserRegEmail) (err error) {
	// 在外部设置即可
	fc := func(ctx context.Context, tx *gorm.DB) error {
		return tx.WithContext(ctx).
			Table(s.UserRegEmailSchema.TableName()).
			Create(dataModel).Error
	}
	err = tx.Do(ctx, fc)
	if err != nil {
		if gormpkg.IsErrDuplicatedKey(err) {
			e := errorv1.DefaultErrorS103UserExist()
			return errorpkg.Wrap(e, err)
		} else {
			e := errorpkg.ErrorInternalServer("")
			return errorpkg.Wrap(e, err)
		}
	}
	return
}

// existCreate exist create
func (s *userRegEmailDataRepo) existCreate(ctx context.Context, dbConn *gorm.DB, dataModel *po.UserRegEmail) (anotherModel *po.UserRegEmail, isNotFound bool, err error) {
	anotherModel = new(po.UserRegEmail)
	err = dbConn.WithContext(ctx).
		Table(s.UserRegEmailSchema.TableName()).
		Where(schemas.FieldId+" = ?", dataModel.Id).
		First(anotherModel).Error
	if err != nil {
		if gormpkg.IsErrRecordNotFound(err) {
			isNotFound = true
			err = nil
		} else {
			e := errorpkg.ErrorInternalServer("")
			err = errorpkg.Wrap(e, err)
		}
		return
	}
	return
}

// ExistCreate exist create
func (s *userRegEmailDataRepo) ExistCreate(ctx context.Context, dataModel *po.UserRegEmail) (anotherModel *po.UserRegEmail, isNotFound bool, err error) {
	return s.existCreate(ctx, s.dbConn, dataModel)
}

// ExistCreateWithDBConn exist create
func (s *userRegEmailDataRepo) ExistCreateWithDBConn(ctx context.Context, dbConn *gorm.DB, dataModel *po.UserRegEmail) (anotherModel *po.UserRegEmail, isNotFound bool, err error) {
	return s.existCreate(ctx, dbConn, dataModel)
}

// createInBatches create many
func (s *userRegEmailDataRepo) createInBatches(ctx context.Context, dbConn *gorm.DB, dataModels []*po.UserRegEmail, batchSize int) (err error) {
	err = dbConn.WithContext(ctx).
		Table(s.UserRegEmailSchema.TableName()).
		CreateInBatches(dataModels, batchSize).Error
	if err != nil {
		e := errorpkg.ErrorInternalServer("")
		return errorpkg.Wrap(e, err)
	}
	return
}

// CreateInBatches create many
func (s *userRegEmailDataRepo) CreateInBatches(ctx context.Context, dataModels []*po.UserRegEmail, batchSize int) error {
	return s.createInBatches(ctx, s.dbConn, dataModels, batchSize)
}

// CreateInBatchesWithDBConn create many
func (s *userRegEmailDataRepo) CreateInBatchesWithDBConn(ctx context.Context, dbConn *gorm.DB, dataModels []*po.UserRegEmail, batchSize int) error {
	return s.createInBatches(ctx, dbConn, dataModels, batchSize)
}

// =============== 更新 ===============

// update update
func (s *userRegEmailDataRepo) update(ctx context.Context, dbConn *gorm.DB, dataModel *po.UserRegEmail) (err error) {
	err = dbConn.WithContext(ctx).
		Table(s.UserRegEmailSchema.TableName()).
		// Where(schemas.FieldId+" = ?", dataModel.Id).
		Save(dataModel).Error
	if err != nil {
		e := errorpkg.ErrorInternalServer("")
		return errorpkg.Wrap(e, err)
	}
	return
}

// Update update
func (s *userRegEmailDataRepo) Update(ctx context.Context, dataModel *po.UserRegEmail) error {
	return s.update(ctx, s.dbConn, dataModel)
}

// UpdateWithDBConn update
func (s *userRegEmailDataRepo) UpdateWithDBConn(ctx context.Context, dbConn *gorm.DB, dataModel *po.UserRegEmail) error {
	return s.update(ctx, dbConn, dataModel)
}

// existUpdate exist update
func (s *userRegEmailDataRepo) existUpdate(ctx context.Context, dbConn *gorm.DB, dataModel *po.UserRegEmail) (anotherModel *po.UserRegEmail, isNotFound bool, err error) {
	anotherModel = new(po.UserRegEmail)
	err = dbConn.WithContext(ctx).
		Table(s.UserRegEmailSchema.TableName()).
		Where(schemas.FieldId+" = ?", dataModel.Id).
		Where(schemas.FieldId+" != ?", dataModel.Id).
		First(anotherModel).Error
	if err != nil {
		if gormpkg.IsErrRecordNotFound(err) {
			isNotFound = true
			err = nil
		} else {
			e := errorpkg.ErrorInternalServer("")
			err = errorpkg.Wrap(e, err)
		}
		return
	}
	return
}

// ExistUpdate exist update
func (s *userRegEmailDataRepo) ExistUpdate(ctx context.Context, dataModel *po.UserRegEmail) (anotherModel *po.UserRegEmail, isNotFound bool, err error) {
	return s.existUpdate(ctx, s.dbConn, dataModel)
}

// ExistUpdateWithDBConn exist update
func (s *userRegEmailDataRepo) ExistUpdateWithDBConn(ctx context.Context, dbConn *gorm.DB, dataModel *po.UserRegEmail) (anotherModel *po.UserRegEmail, isNotFound bool, err error) {
	return s.existUpdate(ctx, dbConn, dataModel)
}

// =============== query one : 查一个 ===============

// queryOneById query one by id
func (s *userRegEmailDataRepo) queryOneById(ctx context.Context, dbConn *gorm.DB, id interface{}) (dataModel *po.UserRegEmail, isNotFound bool, err error) {
	dataModel = new(po.UserRegEmail)
	err = dbConn.WithContext(ctx).
		Table(s.UserRegEmailSchema.TableName()).
		Where(schemas.FieldId+" = ?", id).
		First(dataModel).Error
	if err != nil {
		if gormpkg.IsErrRecordNotFound(err) {
			err = nil
			isNotFound = true
		} else {
			e := errorpkg.ErrorInternalServer("")
			err = errorpkg.Wrap(e, err)
		}
		return
	}
	return
}

// QueryOneById query one by id
func (s *userRegEmailDataRepo) QueryOneById(ctx context.Context, id interface{}) (dataModel *po.UserRegEmail, isNotFound bool, err error) {
	return s.queryOneById(ctx, s.dbConn, id)
}

// QueryOneByIdWithDBConn query one by id
func (s *userRegEmailDataRepo) QueryOneByIdWithDBConn(ctx context.Context, dbConn *gorm.DB, id interface{}) (dataModel *po.UserRegEmail, isNotFound bool, err error) {
	return s.queryOneById(ctx, dbConn, id)
}

// QueryOneByUserEmail query one by id
func (s *userRegEmailDataRepo) QueryOneByUserEmail(ctx context.Context, userEmail string) (dataModel *po.UserRegEmail, isNotFound bool, err error) {
	dataModel = new(po.UserRegEmail)
	err = s.dbConn.WithContext(ctx).
		Table(s.UserRegEmailSchema.TableName()).
		Where(schemas.FieldUserEmail+" = ?", userEmail).
		First(dataModel).Error
	if err != nil {
		if gormpkg.IsErrRecordNotFound(err) {
			err = nil
			isNotFound = true
		} else {
			e := errorpkg.ErrorInternalServer("")
			err = errorpkg.Wrap(e, err)
		}
		return
	}
	return
}

// queryOneByConditions query one by conditions
func (s *userRegEmailDataRepo) queryOneByConditions(ctx context.Context, dbConn *gorm.DB, conditions map[string]interface{}) (dataModel *po.UserRegEmail, isNotFound bool, err error) {
	dataModel = new(po.UserRegEmail)
	dbConn = dbConn.WithContext(ctx).Table(s.UserRegEmailSchema.TableName())
	err = s.WhereConditions(dbConn, conditions).
		First(dataModel).Error
	if err != nil {
		if gormpkg.IsErrRecordNotFound(err) {
			err = nil
			isNotFound = true
		} else {
			e := errorpkg.ErrorInternalServer("")
			err = errorpkg.Wrap(e, err)
		}
		return
	}
	return
}

// QueryOneByConditions query one by conditions
func (s *userRegEmailDataRepo) QueryOneByConditions(ctx context.Context, conditions map[string]interface{}) (dataModel *po.UserRegEmail, isNotFound bool, err error) {
	return s.queryOneByConditions(ctx, s.dbConn, conditions)
}

// QueryOneByConditionsWithDBConn query one by conditions
func (s *userRegEmailDataRepo) QueryOneByConditionsWithDBConn(ctx context.Context, dbConn *gorm.DB, conditions map[string]interface{}) (dataModel *po.UserRegEmail, isNotFound bool, err error) {
	return s.queryOneByConditions(ctx, dbConn, conditions)
}

// =============== query all : 查全部 ===============

// queryAllByConditions query all by conditions
func (s *userRegEmailDataRepo) queryAllByConditions(ctx context.Context, dbConn *gorm.DB, conditions map[string]interface{}) (dataModels []*po.UserRegEmail, err error) {
	dbConn = dbConn.WithContext(ctx).Table(s.UserRegEmailSchema.TableName())
	err = s.WhereConditions(dbConn, conditions).
		Find(&dataModels).Error
	if err != nil {
		e := errorpkg.ErrorInternalServer("")
		err = errorpkg.Wrap(e, err)
		return dataModels, err
	}
	return
}

// QueryAllByConditions query all by conditions
func (s *userRegEmailDataRepo) QueryAllByConditions(ctx context.Context, conditions map[string]interface{}) ([]*po.UserRegEmail, error) {
	return s.queryAllByConditions(ctx, s.dbConn, conditions)
}

// QueryAllByConditionsWithDBConn query all by conditions
func (s *userRegEmailDataRepo) QueryAllByConditionsWithDBConn(ctx context.Context, dbConn *gorm.DB, conditions map[string]interface{}) ([]*po.UserRegEmail, error) {
	return s.queryAllByConditions(ctx, dbConn, conditions)
}

// =============== list : 列表 ===============

// list 列表
func (s *userRegEmailDataRepo) list(ctx context.Context, dbConn *gorm.DB, conditions map[string]interface{}, paginatorArgs *gormpkg.PaginatorArgs) (dataModels []*po.UserRegEmail, recordCount int64, err error) {
	// query where
	dbConn = dbConn.WithContext(ctx).Table(s.UserRegEmailSchema.TableName())
	dbConn = s.WhereConditions(dbConn, conditions)
	dbConn = gormpkg.AssembleWheres(dbConn, paginatorArgs.PageWheres)

	err = dbConn.Count(&recordCount).Error
	if err != nil {
		e := errorpkg.ErrorInternalServer("")
		err = errorpkg.Wrap(e, err)
		return
	} else if recordCount == 0 {
		return // empty
	}

	// pagination
	dbConn = gormpkg.AssembleOrders(dbConn, paginatorArgs.PageOrders)
	err = gormpkg.Paginator(dbConn, paginatorArgs.PageOption).
		Find(&dataModels).Error
	if err != nil {
		e := errorpkg.ErrorInternalServer("")
		err = errorpkg.Wrap(e, err)
		return
	}
	return
}

// List 列表
func (s *userRegEmailDataRepo) List(ctx context.Context, conditions map[string]interface{}, paginatorArgs *gormpkg.PaginatorArgs) ([]*po.UserRegEmail, int64, error) {
	return s.list(ctx, s.dbConn, conditions, paginatorArgs)
}

// ListWithDBConn 列表
func (s *userRegEmailDataRepo) ListWithDBConn(ctx context.Context, dbConn *gorm.DB, conditions map[string]interface{}, paginatorArgs *gormpkg.PaginatorArgs) ([]*po.UserRegEmail, int64, error) {
	return s.list(ctx, dbConn, conditions, paginatorArgs)
}

// =============== delete : 删除 ===============

// delete delete one
func (s *userRegEmailDataRepo) delete(ctx context.Context, dbConn *gorm.DB, dataModel *po.UserRegEmail) (err error) {
	err = dbConn.WithContext(ctx).
		Table(s.UserRegEmailSchema.TableName()).
		Where(schemas.FieldId+" = ?", dataModel.Id).
		Delete(dataModel).Error
	if err != nil {
		e := errorpkg.ErrorInternalServer("")
		err = errorpkg.Wrap(e, err)
		return err
	}
	return
}

// Delete delete one
func (s *userRegEmailDataRepo) Delete(ctx context.Context, dataModel *po.UserRegEmail) error {
	return s.delete(ctx, s.dbConn, dataModel)
}

// DeleteWithDBConn delete one
func (s *userRegEmailDataRepo) DeleteWithDBConn(ctx context.Context, dbConn *gorm.DB, dataModel *po.UserRegEmail) error {
	return s.delete(ctx, dbConn, dataModel)
}

// deleteByIds delete by ids
func (s *userRegEmailDataRepo) deleteByIds(ctx context.Context, dbConn *gorm.DB, ids interface{}) (err error) {
	err = dbConn.WithContext(ctx).
		Table(s.UserRegEmailSchema.TableName()).
		Where(schemas.FieldId+" in (?)", ids).
		Delete(po.UserRegEmail{}).Error
	if err != nil {
		e := errorpkg.ErrorInternalServer("")
		err = errorpkg.Wrap(e, err)
		return err
	}
	return
}

// DeleteByIds delete by ids
func (s *userRegEmailDataRepo) DeleteByIds(ctx context.Context, ids interface{}) error {
	return s.deleteByIds(ctx, s.dbConn, ids)
}

// DeleteByIdsWithDBConn delete by ids
func (s *userRegEmailDataRepo) DeleteByIdsWithDBConn(ctx context.Context, dbConn *gorm.DB, ids interface{}) error {
	return s.deleteByIds(ctx, dbConn, ids)
}

// =============== insert : 批量入库 ===============

var _ gormpkg.BatchInsertRepo = new(UserRegEmailSlice)

// UserRegEmailSlice 表切片
type UserRegEmailSlice []*po.UserRegEmail

// TableName 表名
func (s *UserRegEmailSlice) TableName() string {
	return schemas.UserRegEmailSchema.TableName()
}

// Len 长度
func (s *UserRegEmailSlice) Len() int {
	return len(*s)
}

// InsertColumns 批量入库的列
func (s *UserRegEmailSlice) InsertColumns() (columnList []string, placeholder string) {
	// columns
	columnList = []string{
		schemas.FieldCreatedTime, schemas.FieldUpdatedTime,
		schemas.FieldDeletedTime, schemas.FieldUserId,
		schemas.FieldUserEmail,
	}

	// placeholders
	insertPlaceholderBytes := bytes.Repeat([]byte("?, "), len(columnList))
	insertPlaceholderBytes = bytes.TrimSuffix(insertPlaceholderBytes, []byte(", "))

	return columnList, string(insertPlaceholderBytes)
}

// InsertValues 批量入库的值
func (s *UserRegEmailSlice) InsertValues(args *gormpkg.BatchInsertValueArgs) (prepareData []interface{}, placeholderSlice []string) {
	dataModels := (*s)[args.StepStart:args.StepEnd]
	for index := range dataModels {
		// placeholder
		placeholderSlice = append(placeholderSlice, "("+args.InsertPlaceholder+")")

		// prepare data
		prepareData = append(prepareData, dataModels[index].CreatedTime)
		prepareData = append(prepareData, dataModels[index].UpdatedTime)
		prepareData = append(prepareData, dataModels[index].DeletedTime)
		prepareData = append(prepareData, dataModels[index].UserId)
		prepareData = append(prepareData, dataModels[index].UserEmail)
	}
	return prepareData, placeholderSlice
}

// UpdateColumns 批量入库的列
func (s *UserRegEmailSlice) UpdateColumns() (columnList []string) {
	// columns
	columnList = []string{
		schemas.FieldCreatedTime + "=excluded." + schemas.FieldCreatedTime,
		schemas.FieldUpdatedTime + "=excluded." + schemas.FieldUpdatedTime,
		schemas.FieldDeletedTime + "=excluded." + schemas.FieldDeletedTime,
		schemas.FieldUserId + "=excluded." + schemas.FieldUserId,
		schemas.FieldUserEmail + "=excluded." + schemas.FieldUserEmail,
	}
	return columnList
}

// ConflictActionForMySQL 更新冲突时的操作
func (s *UserRegEmailSlice) ConflictActionForMySQL() (req *gormpkg.BatchInsertConflictActionReq) {
	req = &gormpkg.BatchInsertConflictActionReq{
		OnConflictValueAlias:  "AS excluded",
		OnConflictTarget:      "ON DUPLICATE KEY",
		OnConflictAction:      "UPDATE " + strings.Join(s.UpdateColumns(), ", "),
		OnConflictPrepareData: nil,
	}
	return req
}

// ConflictActionForPostgres 更新冲突时的操作
func (s *UserRegEmailSlice) ConflictActionForPostgres() (req *gormpkg.BatchInsertConflictActionReq) {
	req = &gormpkg.BatchInsertConflictActionReq{
		OnConflictValueAlias:  "",
		OnConflictTarget:      "ON CONFLICT(id)",
		OnConflictAction:      "DO UPDATE SET " + strings.Join(s.UpdateColumns(), ", "),
		OnConflictPrepareData: nil,
	}
	return req
}

// insert 批量插入
func (s *userRegEmailDataRepo) insert(ctx context.Context, dbConn *gorm.DB, dataModels UserRegEmailSlice) error {
	err := gormpkg.BatchInsertWithContext(ctx, dbConn, &dataModels)
	if err != nil {
		e := errorpkg.ErrorInternalServer("")
		err = errorpkg.Wrap(e, err)
		return err
	}
	return nil
}

// Insert 批量插入
func (s *userRegEmailDataRepo) Insert(ctx context.Context, dataModels []*po.UserRegEmail) error {
	return s.insert(ctx, s.dbConn, dataModels)
}

// InsertWithDBConn 批量插入
func (s *userRegEmailDataRepo) InsertWithDBConn(ctx context.Context, dbConn *gorm.DB, dataModels []*po.UserRegEmail) error {
	return s.insert(ctx, dbConn, dataModels)
}

// =============== conditions : 条件 ===============

// WhereConditions orm where
func (s *userRegEmailDataRepo) WhereConditions(dbConn *gorm.DB, conditions map[string]interface{}) *gorm.DB {

	// table name
	// tableName := s.UserRegEmailSchema.TableName()

	// On-demand loading

	// id
	// if data, ok := conditions[schemas.FieldId]; ok {
	// 	   dbConn = dbConn.Where(tableName+"."+schemas.FieldId+" = ?", data)
	// }

	// created_time
	// if data, ok := conditions[schemas.FieldCreatedTime]; ok {
	// 	   dbConn = dbConn.Where(tableName+"."+schemas.FieldCreatedTime+" = ?", data)
	// }

	// updated_time
	// if data, ok := conditions[schemas.FieldUpdatedTime]; ok {
	// 	   dbConn = dbConn.Where(tableName+"."+schemas.FieldUpdatedTime+" = ?", data)
	// }

	// deleted_time
	// if data, ok := conditions[schemas.FieldDeletedTime]; ok {
	// 	   dbConn = dbConn.Where(tableName+"."+schemas.FieldDeletedTime+" = ?", data)
	// }

	// user_id
	// if data, ok := conditions[schemas.FieldUserId]; ok {
	// 	   dbConn = dbConn.Where(tableName+"."+schemas.FieldUserId+" = ?", data)
	// }

	// user_email
	// if data, ok := conditions[schemas.FieldUserEmail]; ok {
	// 	   dbConn = dbConn.Where(tableName+"."+schemas.FieldUserEmail+" = ?", data)
	// }

	return dbConn
}

// UpdateColumns update columns
func (s *userRegEmailDataRepo) UpdateColumns(conditions map[string]interface{}) map[string]interface{} {

	// update columns
	updateColumns := make(map[string]interface{})

	// On-demand loading

	// id
	//if data, ok := conditions[schemas.FieldId]; ok {
	//	updateColumns[schemas.FieldId] = data
	//}

	// created_time
	//if data, ok := conditions[schemas.FieldCreatedTime]; ok {
	//	updateColumns[schemas.FieldCreatedTime] = data
	//}

	// updated_time
	//if data, ok := conditions[schemas.FieldUpdatedTime]; ok {
	//	updateColumns[schemas.FieldUpdatedTime] = data
	//}

	// deleted_time
	//if data, ok := conditions[schemas.FieldDeletedTime]; ok {
	//	updateColumns[schemas.FieldDeletedTime] = data
	//}

	// user_id
	//if data, ok := conditions[schemas.FieldUserId]; ok {
	//	updateColumns[schemas.FieldUserId] = data
	//}

	// user_email
	//if data, ok := conditions[schemas.FieldUserEmail]; ok {
	//	updateColumns[schemas.FieldUserEmail] = data
	//}

	return updateColumns
}
