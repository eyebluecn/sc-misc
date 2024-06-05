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

func newEditorPO(db *gorm.DB, opts ...gen.DOOption) editorPO {
	_editorPO := editorPO{}

	_editorPO.editorPODo.UseDB(db, opts...)
	_editorPO.editorPODo.UseModel(&po.EditorPO{})

	tableName := _editorPO.editorPODo.TableName()
	_editorPO.ALL = field.NewAsterisk(tableName)
	_editorPO.ID = field.NewInt64(tableName, "id")
	_editorPO.CreateTime = field.NewTime(tableName, "create_time")
	_editorPO.UpdateTime = field.NewTime(tableName, "update_time")
	_editorPO.Username = field.NewString(tableName, "username")
	_editorPO.Password = field.NewString(tableName, "password")
	_editorPO.WorkNo = field.NewString(tableName, "work_no")

	_editorPO.fillFieldMap()

	return _editorPO
}

// editorPO 小二表
type editorPO struct {
	editorPODo editorPODo

	ALL        field.Asterisk
	ID         field.Int64  // 主键
	CreateTime field.Time   // 创建时间
	UpdateTime field.Time   // 更新时间
	Username   field.String // 昵称
	Password   field.String // 密码
	WorkNo     field.String // 工号

	fieldMap map[string]field.Expr
}

func (e editorPO) Table(newTableName string) *editorPO {
	e.editorPODo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e editorPO) As(alias string) *editorPO {
	e.editorPODo.DO = *(e.editorPODo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *editorPO) updateTableName(table string) *editorPO {
	e.ALL = field.NewAsterisk(table)
	e.ID = field.NewInt64(table, "id")
	e.CreateTime = field.NewTime(table, "create_time")
	e.UpdateTime = field.NewTime(table, "update_time")
	e.Username = field.NewString(table, "username")
	e.Password = field.NewString(table, "password")
	e.WorkNo = field.NewString(table, "work_no")

	e.fillFieldMap()

	return e
}

func (e *editorPO) WithContext(ctx context.Context) *editorPODo { return e.editorPODo.WithContext(ctx) }

func (e editorPO) TableName() string { return e.editorPODo.TableName() }

func (e editorPO) Alias() string { return e.editorPODo.Alias() }

func (e editorPO) Columns(cols ...field.Expr) gen.Columns { return e.editorPODo.Columns(cols...) }

func (e *editorPO) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *editorPO) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 6)
	e.fieldMap["id"] = e.ID
	e.fieldMap["create_time"] = e.CreateTime
	e.fieldMap["update_time"] = e.UpdateTime
	e.fieldMap["username"] = e.Username
	e.fieldMap["password"] = e.Password
	e.fieldMap["work_no"] = e.WorkNo
}

func (e editorPO) clone(db *gorm.DB) editorPO {
	e.editorPODo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e editorPO) replaceDB(db *gorm.DB) editorPO {
	e.editorPODo.ReplaceDB(db)
	return e
}

type editorPODo struct{ gen.DO }

func (e editorPODo) Debug() *editorPODo {
	return e.withDO(e.DO.Debug())
}

func (e editorPODo) WithContext(ctx context.Context) *editorPODo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e editorPODo) ReadDB() *editorPODo {
	return e.Clauses(dbresolver.Read)
}

func (e editorPODo) WriteDB() *editorPODo {
	return e.Clauses(dbresolver.Write)
}

func (e editorPODo) Session(config *gorm.Session) *editorPODo {
	return e.withDO(e.DO.Session(config))
}

func (e editorPODo) Clauses(conds ...clause.Expression) *editorPODo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e editorPODo) Returning(value interface{}, columns ...string) *editorPODo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e editorPODo) Not(conds ...gen.Condition) *editorPODo {
	return e.withDO(e.DO.Not(conds...))
}

func (e editorPODo) Or(conds ...gen.Condition) *editorPODo {
	return e.withDO(e.DO.Or(conds...))
}

func (e editorPODo) Select(conds ...field.Expr) *editorPODo {
	return e.withDO(e.DO.Select(conds...))
}

func (e editorPODo) Where(conds ...gen.Condition) *editorPODo {
	return e.withDO(e.DO.Where(conds...))
}

func (e editorPODo) Order(conds ...field.Expr) *editorPODo {
	return e.withDO(e.DO.Order(conds...))
}

func (e editorPODo) Distinct(cols ...field.Expr) *editorPODo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e editorPODo) Omit(cols ...field.Expr) *editorPODo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e editorPODo) Join(table schema.Tabler, on ...field.Expr) *editorPODo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e editorPODo) LeftJoin(table schema.Tabler, on ...field.Expr) *editorPODo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e editorPODo) RightJoin(table schema.Tabler, on ...field.Expr) *editorPODo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e editorPODo) Group(cols ...field.Expr) *editorPODo {
	return e.withDO(e.DO.Group(cols...))
}

func (e editorPODo) Having(conds ...gen.Condition) *editorPODo {
	return e.withDO(e.DO.Having(conds...))
}

func (e editorPODo) Limit(limit int) *editorPODo {
	return e.withDO(e.DO.Limit(limit))
}

func (e editorPODo) Offset(offset int) *editorPODo {
	return e.withDO(e.DO.Offset(offset))
}

func (e editorPODo) Scopes(funcs ...func(gen.Dao) gen.Dao) *editorPODo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e editorPODo) Unscoped() *editorPODo {
	return e.withDO(e.DO.Unscoped())
}

func (e editorPODo) Create(values ...*po.EditorPO) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e editorPODo) CreateInBatches(values []*po.EditorPO, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e editorPODo) Save(values ...*po.EditorPO) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e editorPODo) First() (*po.EditorPO, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*po.EditorPO), nil
	}
}

func (e editorPODo) Take() (*po.EditorPO, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*po.EditorPO), nil
	}
}

func (e editorPODo) Last() (*po.EditorPO, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*po.EditorPO), nil
	}
}

func (e editorPODo) Find() ([]*po.EditorPO, error) {
	result, err := e.DO.Find()
	return result.([]*po.EditorPO), err
}

func (e editorPODo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*po.EditorPO, err error) {
	buf := make([]*po.EditorPO, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e editorPODo) FindInBatches(result *[]*po.EditorPO, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e editorPODo) Attrs(attrs ...field.AssignExpr) *editorPODo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e editorPODo) Assign(attrs ...field.AssignExpr) *editorPODo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e editorPODo) Joins(fields ...field.RelationField) *editorPODo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e editorPODo) Preload(fields ...field.RelationField) *editorPODo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e editorPODo) FirstOrInit() (*po.EditorPO, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*po.EditorPO), nil
	}
}

func (e editorPODo) FirstOrCreate() (*po.EditorPO, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*po.EditorPO), nil
	}
}

func (e editorPODo) FindByPage(offset int, limit int) (result []*po.EditorPO, count int64, err error) {
	result, err = e.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = e.Offset(-1).Limit(-1).Count()
	return
}

func (e editorPODo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e editorPODo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e editorPODo) Delete(models ...*po.EditorPO) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *editorPODo) withDO(do gen.Dao) *editorPODo {
	e.DO = *do.(*gen.DO)
	return e
}
