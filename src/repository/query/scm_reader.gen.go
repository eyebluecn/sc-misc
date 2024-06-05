// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gen/helper"

	"gorm.io/plugin/dbresolver"
)

func newReaderPO(db *gorm.DB, opts ...gen.DOOption) readerPO {
	_readerPO := readerPO{}

	_readerPO.readerPODo.UseDB(db, opts...)
	_readerPO.readerPODo.UseModel(&po.ReaderPO{})

	tableName := _readerPO.readerPODo.TableName()
	_readerPO.ALL = field.NewAsterisk(tableName)
	_readerPO.ID = field.NewInt64(tableName, "id")
	_readerPO.CreateTime = field.NewTime(tableName, "create_time")
	_readerPO.UpdateTime = field.NewTime(tableName, "update_time")
	_readerPO.Username = field.NewString(tableName, "username")
	_readerPO.Password = field.NewString(tableName, "password")

	_readerPO.fillFieldMap()

	return _readerPO
}

// readerPO 读者表
type readerPO struct {
	readerPODo readerPODo

	ALL        field.Asterisk
	ID         field.Int64  // 主键
	CreateTime field.Time   // 创建时间
	UpdateTime field.Time   // 更新时间
	Username   field.String // 用户名
	Password   field.String // 密码

	fieldMap map[string]field.Expr
}

func (r readerPO) Table(newTableName string) *readerPO {
	r.readerPODo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r readerPO) As(alias string) *readerPO {
	r.readerPODo.DO = *(r.readerPODo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *readerPO) updateTableName(table string) *readerPO {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewInt64(table, "id")
	r.CreateTime = field.NewTime(table, "create_time")
	r.UpdateTime = field.NewTime(table, "update_time")
	r.Username = field.NewString(table, "username")
	r.Password = field.NewString(table, "password")

	r.fillFieldMap()

	return r
}

func (r *readerPO) WithContext(ctx context.Context) *readerPODo { return r.readerPODo.WithContext(ctx) }

func (r readerPO) TableName() string { return r.readerPODo.TableName() }

func (r readerPO) Alias() string { return r.readerPODo.Alias() }

func (r readerPO) Columns(cols ...field.Expr) gen.Columns { return r.readerPODo.Columns(cols...) }

func (r *readerPO) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *readerPO) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 5)
	r.fieldMap["id"] = r.ID
	r.fieldMap["create_time"] = r.CreateTime
	r.fieldMap["update_time"] = r.UpdateTime
	r.fieldMap["username"] = r.Username
	r.fieldMap["password"] = r.Password
}

func (r readerPO) clone(db *gorm.DB) readerPO {
	r.readerPODo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r readerPO) replaceDB(db *gorm.DB) readerPO {
	r.readerPODo.ReplaceDB(db)
	return r
}

type readerPODo struct{ gen.DO }

func (r readerPODo) Debug() *readerPODo {
	return r.withDO(r.DO.Debug())
}

func (r readerPODo) WithContext(ctx context.Context) *readerPODo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r readerPODo) ReadDB() *readerPODo {
	return r.Clauses(dbresolver.Read)
}

func (r readerPODo) WriteDB() *readerPODo {
	return r.Clauses(dbresolver.Write)
}

func (r readerPODo) Session(config *gorm.Session) *readerPODo {
	return r.withDO(r.DO.Session(config))
}

func (r readerPODo) Clauses(conds ...clause.Expression) *readerPODo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r readerPODo) Returning(value interface{}, columns ...string) *readerPODo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r readerPODo) Not(conds ...gen.Condition) *readerPODo {
	return r.withDO(r.DO.Not(conds...))
}

func (r readerPODo) Or(conds ...gen.Condition) *readerPODo {
	return r.withDO(r.DO.Or(conds...))
}

func (r readerPODo) Select(conds ...field.Expr) *readerPODo {
	return r.withDO(r.DO.Select(conds...))
}

func (r readerPODo) Where(conds ...gen.Condition) *readerPODo {
	return r.withDO(r.DO.Where(conds...))
}

func (r readerPODo) Order(conds ...field.Expr) *readerPODo {
	return r.withDO(r.DO.Order(conds...))
}

func (r readerPODo) Distinct(cols ...field.Expr) *readerPODo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r readerPODo) Omit(cols ...field.Expr) *readerPODo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r readerPODo) Join(table schema.Tabler, on ...field.Expr) *readerPODo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r readerPODo) LeftJoin(table schema.Tabler, on ...field.Expr) *readerPODo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r readerPODo) RightJoin(table schema.Tabler, on ...field.Expr) *readerPODo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r readerPODo) Group(cols ...field.Expr) *readerPODo {
	return r.withDO(r.DO.Group(cols...))
}

func (r readerPODo) Having(conds ...gen.Condition) *readerPODo {
	return r.withDO(r.DO.Having(conds...))
}

func (r readerPODo) Limit(limit int) *readerPODo {
	return r.withDO(r.DO.Limit(limit))
}

func (r readerPODo) Offset(offset int) *readerPODo {
	return r.withDO(r.DO.Offset(offset))
}

func (r readerPODo) Scopes(funcs ...func(gen.Dao) gen.Dao) *readerPODo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r readerPODo) Unscoped() *readerPODo {
	return r.withDO(r.DO.Unscoped())
}

func (r readerPODo) Create(values ...*po.ReaderPO) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r readerPODo) CreateInBatches(values []*po.ReaderPO, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r readerPODo) Save(values ...*po.ReaderPO) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r readerPODo) First() (*po.ReaderPO, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*po.ReaderPO), nil
	}
}

func (r readerPODo) Take() (*po.ReaderPO, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*po.ReaderPO), nil
	}
}

func (r readerPODo) Last() (*po.ReaderPO, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*po.ReaderPO), nil
	}
}

func (r readerPODo) Find() ([]*po.ReaderPO, error) {
	result, err := r.DO.Find()
	return result.([]*po.ReaderPO), err
}

func (r readerPODo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*po.ReaderPO, err error) {
	buf := make([]*po.ReaderPO, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r readerPODo) FindInBatches(result *[]*po.ReaderPO, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r readerPODo) Attrs(attrs ...field.AssignExpr) *readerPODo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r readerPODo) Assign(attrs ...field.AssignExpr) *readerPODo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r readerPODo) Joins(fields ...field.RelationField) *readerPODo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r readerPODo) Preload(fields ...field.RelationField) *readerPODo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r readerPODo) FirstOrInit() (*po.ReaderPO, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*po.ReaderPO), nil
	}
}

func (r readerPODo) FirstOrCreate() (*po.ReaderPO, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*po.ReaderPO), nil
	}
}

func (r readerPODo) FindByPage(offset int, limit int) (result []*po.ReaderPO, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r readerPODo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r readerPODo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r readerPODo) Delete(models ...*po.ReaderPO) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *readerPODo) withDO(do gen.Dao) *readerPODo {
	r.DO = *do.(*gen.DO)
	return r
}
