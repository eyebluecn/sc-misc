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

func newColumnQuotePO(db *gorm.DB, opts ...gen.DOOption) columnQuotePO {
	_columnQuotePO := columnQuotePO{}

	_columnQuotePO.columnQuotePODo.UseDB(db, opts...)
	_columnQuotePO.columnQuotePODo.UseModel(&po.ColumnQuotePO{})

	tableName := _columnQuotePO.columnQuotePODo.TableName()
	_columnQuotePO.ALL = field.NewAsterisk(tableName)
	_columnQuotePO.ID = field.NewInt64(tableName, "id")
	_columnQuotePO.CreateTime = field.NewTime(tableName, "create_time")
	_columnQuotePO.UpdateTime = field.NewTime(tableName, "update_time")
	_columnQuotePO.ColumnID = field.NewInt64(tableName, "column_id")
	_columnQuotePO.EditorID = field.NewInt64(tableName, "editor_id")
	_columnQuotePO.Price = field.NewInt64(tableName, "price")
	_columnQuotePO.Status = field.NewInt32(tableName, "status")

	_columnQuotePO.fillFieldMap()

	return _columnQuotePO
}

// columnQuotePO 专栏报价表
type columnQuotePO struct {
	columnQuotePODo columnQuotePODo

	ALL        field.Asterisk
	ID         field.Int64 // 主键
	CreateTime field.Time  // 创建时间
	UpdateTime field.Time  // 更新时间
	ColumnID   field.Int64 // 专栏id
	EditorID   field.Int64 // 编辑id
	Price      field.Int64 // 价格（单位：分）
	Status     field.Int32 // 报价状态 0/1  未生效/已生效

	fieldMap map[string]field.Expr
}

func (c columnQuotePO) Table(newTableName string) *columnQuotePO {
	c.columnQuotePODo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c columnQuotePO) As(alias string) *columnQuotePO {
	c.columnQuotePODo.DO = *(c.columnQuotePODo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *columnQuotePO) updateTableName(table string) *columnQuotePO {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt64(table, "id")
	c.CreateTime = field.NewTime(table, "create_time")
	c.UpdateTime = field.NewTime(table, "update_time")
	c.ColumnID = field.NewInt64(table, "column_id")
	c.EditorID = field.NewInt64(table, "editor_id")
	c.Price = field.NewInt64(table, "price")
	c.Status = field.NewInt32(table, "status")

	c.fillFieldMap()

	return c
}

func (c *columnQuotePO) WithContext(ctx context.Context) *columnQuotePODo {
	return c.columnQuotePODo.WithContext(ctx)
}

func (c columnQuotePO) TableName() string { return c.columnQuotePODo.TableName() }

func (c columnQuotePO) Alias() string { return c.columnQuotePODo.Alias() }

func (c columnQuotePO) Columns(cols ...field.Expr) gen.Columns {
	return c.columnQuotePODo.Columns(cols...)
}

func (c *columnQuotePO) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *columnQuotePO) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 7)
	c.fieldMap["id"] = c.ID
	c.fieldMap["create_time"] = c.CreateTime
	c.fieldMap["update_time"] = c.UpdateTime
	c.fieldMap["column_id"] = c.ColumnID
	c.fieldMap["editor_id"] = c.EditorID
	c.fieldMap["price"] = c.Price
	c.fieldMap["status"] = c.Status
}

func (c columnQuotePO) clone(db *gorm.DB) columnQuotePO {
	c.columnQuotePODo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c columnQuotePO) replaceDB(db *gorm.DB) columnQuotePO {
	c.columnQuotePODo.ReplaceDB(db)
	return c
}

type columnQuotePODo struct{ gen.DO }

func (c columnQuotePODo) Debug() *columnQuotePODo {
	return c.withDO(c.DO.Debug())
}

func (c columnQuotePODo) WithContext(ctx context.Context) *columnQuotePODo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c columnQuotePODo) ReadDB() *columnQuotePODo {
	return c.Clauses(dbresolver.Read)
}

func (c columnQuotePODo) WriteDB() *columnQuotePODo {
	return c.Clauses(dbresolver.Write)
}

func (c columnQuotePODo) Session(config *gorm.Session) *columnQuotePODo {
	return c.withDO(c.DO.Session(config))
}

func (c columnQuotePODo) Clauses(conds ...clause.Expression) *columnQuotePODo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c columnQuotePODo) Returning(value interface{}, columns ...string) *columnQuotePODo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c columnQuotePODo) Not(conds ...gen.Condition) *columnQuotePODo {
	return c.withDO(c.DO.Not(conds...))
}

func (c columnQuotePODo) Or(conds ...gen.Condition) *columnQuotePODo {
	return c.withDO(c.DO.Or(conds...))
}

func (c columnQuotePODo) Select(conds ...field.Expr) *columnQuotePODo {
	return c.withDO(c.DO.Select(conds...))
}

func (c columnQuotePODo) Where(conds ...gen.Condition) *columnQuotePODo {
	return c.withDO(c.DO.Where(conds...))
}

func (c columnQuotePODo) Order(conds ...field.Expr) *columnQuotePODo {
	return c.withDO(c.DO.Order(conds...))
}

func (c columnQuotePODo) Distinct(cols ...field.Expr) *columnQuotePODo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c columnQuotePODo) Omit(cols ...field.Expr) *columnQuotePODo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c columnQuotePODo) Join(table schema.Tabler, on ...field.Expr) *columnQuotePODo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c columnQuotePODo) LeftJoin(table schema.Tabler, on ...field.Expr) *columnQuotePODo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c columnQuotePODo) RightJoin(table schema.Tabler, on ...field.Expr) *columnQuotePODo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c columnQuotePODo) Group(cols ...field.Expr) *columnQuotePODo {
	return c.withDO(c.DO.Group(cols...))
}

func (c columnQuotePODo) Having(conds ...gen.Condition) *columnQuotePODo {
	return c.withDO(c.DO.Having(conds...))
}

func (c columnQuotePODo) Limit(limit int) *columnQuotePODo {
	return c.withDO(c.DO.Limit(limit))
}

func (c columnQuotePODo) Offset(offset int) *columnQuotePODo {
	return c.withDO(c.DO.Offset(offset))
}

func (c columnQuotePODo) Scopes(funcs ...func(gen.Dao) gen.Dao) *columnQuotePODo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c columnQuotePODo) Unscoped() *columnQuotePODo {
	return c.withDO(c.DO.Unscoped())
}

func (c columnQuotePODo) Create(values ...*po.ColumnQuotePO) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c columnQuotePODo) CreateInBatches(values []*po.ColumnQuotePO, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c columnQuotePODo) Save(values ...*po.ColumnQuotePO) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c columnQuotePODo) First() (*po.ColumnQuotePO, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*po.ColumnQuotePO), nil
	}
}

func (c columnQuotePODo) Take() (*po.ColumnQuotePO, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*po.ColumnQuotePO), nil
	}
}

func (c columnQuotePODo) Last() (*po.ColumnQuotePO, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*po.ColumnQuotePO), nil
	}
}

func (c columnQuotePODo) Find() ([]*po.ColumnQuotePO, error) {
	result, err := c.DO.Find()
	return result.([]*po.ColumnQuotePO), err
}

func (c columnQuotePODo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*po.ColumnQuotePO, err error) {
	buf := make([]*po.ColumnQuotePO, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c columnQuotePODo) FindInBatches(result *[]*po.ColumnQuotePO, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c columnQuotePODo) Attrs(attrs ...field.AssignExpr) *columnQuotePODo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c columnQuotePODo) Assign(attrs ...field.AssignExpr) *columnQuotePODo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c columnQuotePODo) Joins(fields ...field.RelationField) *columnQuotePODo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c columnQuotePODo) Preload(fields ...field.RelationField) *columnQuotePODo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c columnQuotePODo) FirstOrInit() (*po.ColumnQuotePO, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*po.ColumnQuotePO), nil
	}
}

func (c columnQuotePODo) FirstOrCreate() (*po.ColumnQuotePO, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*po.ColumnQuotePO), nil
	}
}

func (c columnQuotePODo) FindByPage(offset int, limit int) (result []*po.ColumnQuotePO, count int64, err error) {
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

func (c columnQuotePODo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c columnQuotePODo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c columnQuotePODo) Delete(models ...*po.ColumnQuotePO) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *columnQuotePODo) withDO(do gen.Dao) *columnQuotePODo {
	c.DO = *do.(*gen.DO)
	return c
}