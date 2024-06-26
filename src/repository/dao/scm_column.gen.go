// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/eyebluecn/sc-misc/src/model/po"
)

func newColumnPO(db *gorm.DB, opts ...gen.DOOption) columnPO {
	_columnPO := columnPO{}

	_columnPO.columnPODo.UseDB(db, opts...)
	_columnPO.columnPODo.UseModel(&po.ColumnPO{})

	tableName := _columnPO.columnPODo.TableName()
	_columnPO.ALL = field.NewAsterisk(tableName)
	_columnPO.ID = field.NewInt64(tableName, "id")
	_columnPO.CreateTime = field.NewTime(tableName, "create_time")
	_columnPO.UpdateTime = field.NewTime(tableName, "update_time")
	_columnPO.Name = field.NewString(tableName, "name")
	_columnPO.AuthorID = field.NewInt64(tableName, "author_id")
	_columnPO.Status = field.NewInt32(tableName, "status")

	_columnPO.fillFieldMap()

	return _columnPO
}

// columnPO 专栏表
type columnPO struct {
	columnPODo columnPODo

	ALL        field.Asterisk
	ID         field.Int64  // 主键
	CreateTime field.Time   // 创建时间
	UpdateTime field.Time   // 更新时间
	Name       field.String // 专栏名称
	AuthorID   field.Int64  // 作者id
	Status     field.Int32  // 状态 NEW/OK/DISABLED 未发布/已生效/已禁用

	fieldMap map[string]field.Expr
}

func (c columnPO) Table(newTableName string) *columnPO {
	c.columnPODo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c columnPO) As(alias string) *columnPO {
	c.columnPODo.DO = *(c.columnPODo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *columnPO) updateTableName(table string) *columnPO {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt64(table, "id")
	c.CreateTime = field.NewTime(table, "create_time")
	c.UpdateTime = field.NewTime(table, "update_time")
	c.Name = field.NewString(table, "name")
	c.AuthorID = field.NewInt64(table, "author_id")
	c.Status = field.NewInt32(table, "status")

	c.fillFieldMap()

	return c
}

func (c *columnPO) WithContext(ctx context.Context) *columnPODo { return c.columnPODo.WithContext(ctx) }

func (c columnPO) TableName() string { return c.columnPODo.TableName() }

func (c columnPO) Alias() string { return c.columnPODo.Alias() }

func (c columnPO) Columns(cols ...field.Expr) gen.Columns { return c.columnPODo.Columns(cols...) }

func (c *columnPO) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *columnPO) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 6)
	c.fieldMap["id"] = c.ID
	c.fieldMap["create_time"] = c.CreateTime
	c.fieldMap["update_time"] = c.UpdateTime
	c.fieldMap["name"] = c.Name
	c.fieldMap["author_id"] = c.AuthorID
	c.fieldMap["status"] = c.Status
}

func (c columnPO) clone(db *gorm.DB) columnPO {
	c.columnPODo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c columnPO) replaceDB(db *gorm.DB) columnPO {
	c.columnPODo.ReplaceDB(db)
	return c
}

type columnPODo struct{ gen.DO }

func (c columnPODo) Debug() *columnPODo {
	return c.withDO(c.DO.Debug())
}

func (c columnPODo) WithContext(ctx context.Context) *columnPODo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c columnPODo) ReadDB() *columnPODo {
	return c.Clauses(dbresolver.Read)
}

func (c columnPODo) WriteDB() *columnPODo {
	return c.Clauses(dbresolver.Write)
}

func (c columnPODo) Session(config *gorm.Session) *columnPODo {
	return c.withDO(c.DO.Session(config))
}

func (c columnPODo) Clauses(conds ...clause.Expression) *columnPODo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c columnPODo) Returning(value interface{}, columns ...string) *columnPODo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c columnPODo) Not(conds ...gen.Condition) *columnPODo {
	return c.withDO(c.DO.Not(conds...))
}

func (c columnPODo) Or(conds ...gen.Condition) *columnPODo {
	return c.withDO(c.DO.Or(conds...))
}

func (c columnPODo) Select(conds ...field.Expr) *columnPODo {
	return c.withDO(c.DO.Select(conds...))
}

func (c columnPODo) Where(conds ...gen.Condition) *columnPODo {
	return c.withDO(c.DO.Where(conds...))
}

func (c columnPODo) Order(conds ...field.Expr) *columnPODo {
	return c.withDO(c.DO.Order(conds...))
}

func (c columnPODo) Distinct(cols ...field.Expr) *columnPODo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c columnPODo) Omit(cols ...field.Expr) *columnPODo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c columnPODo) Join(table schema.Tabler, on ...field.Expr) *columnPODo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c columnPODo) LeftJoin(table schema.Tabler, on ...field.Expr) *columnPODo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c columnPODo) RightJoin(table schema.Tabler, on ...field.Expr) *columnPODo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c columnPODo) Group(cols ...field.Expr) *columnPODo {
	return c.withDO(c.DO.Group(cols...))
}

func (c columnPODo) Having(conds ...gen.Condition) *columnPODo {
	return c.withDO(c.DO.Having(conds...))
}

func (c columnPODo) Limit(limit int) *columnPODo {
	return c.withDO(c.DO.Limit(limit))
}

func (c columnPODo) Offset(offset int) *columnPODo {
	return c.withDO(c.DO.Offset(offset))
}

func (c columnPODo) Scopes(funcs ...func(gen.Dao) gen.Dao) *columnPODo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c columnPODo) Unscoped() *columnPODo {
	return c.withDO(c.DO.Unscoped())
}

func (c columnPODo) Create(values ...*po.ColumnPO) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c columnPODo) CreateInBatches(values []*po.ColumnPO, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c columnPODo) Save(values ...*po.ColumnPO) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c columnPODo) First() (*po.ColumnPO, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*po.ColumnPO), nil
	}
}

func (c columnPODo) Take() (*po.ColumnPO, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*po.ColumnPO), nil
	}
}

func (c columnPODo) Last() (*po.ColumnPO, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*po.ColumnPO), nil
	}
}

func (c columnPODo) Find() ([]*po.ColumnPO, error) {
	result, err := c.DO.Find()
	return result.([]*po.ColumnPO), err
}

func (c columnPODo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*po.ColumnPO, err error) {
	buf := make([]*po.ColumnPO, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c columnPODo) FindInBatches(result *[]*po.ColumnPO, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c columnPODo) Attrs(attrs ...field.AssignExpr) *columnPODo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c columnPODo) Assign(attrs ...field.AssignExpr) *columnPODo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c columnPODo) Joins(fields ...field.RelationField) *columnPODo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c columnPODo) Preload(fields ...field.RelationField) *columnPODo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c columnPODo) FirstOrInit() (*po.ColumnPO, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*po.ColumnPO), nil
	}
}

func (c columnPODo) FirstOrCreate() (*po.ColumnPO, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*po.ColumnPO), nil
	}
}

func (c columnPODo) FindByPage(offset int, limit int) (result []*po.ColumnPO, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c columnPODo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c columnPODo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c columnPODo) Delete(models ...*po.ColumnPO) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *columnPODo) withDO(do gen.Dao) *columnPODo {
	c.DO = *do.(*gen.DO)
	return c
}
