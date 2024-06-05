package repo

import (
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc/src/common/errs"
	"github.com/eyebluecn/sc-misc/src/converter/do2po"
	"github.com/eyebluecn/sc-misc/src/converter/po2do"
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/query"
	query2 "github.com/eyebluecn/sc-misc/src/model/query/enums"
	"github.com/eyebluecn/sc-misc/src/model/result"
	"github.com/eyebluecn/sc-misc/src/repository/config"
	"github.com/eyebluecn/sc-misc/src/repository/dao"
	"gorm.io/gen"
	"gorm.io/gorm"
	"time"
)

type ColumnRepo struct {
}

func NewColumnRepo() ColumnRepo {
	return ColumnRepo{}
}

// 新建一个Reader
func (receiver ColumnRepo) Insert(
	ctx context.Context,
	column *do.ColumnDO,
) (*do.ColumnDO, error) {
	table := dao.Use(config.DB).ColumnPO

	//时间置为当前
	column.CreateTime = time.Now()
	column.UpdateTime = time.Now()

	columnDO := do2po.ConvertColumnPO(column)

	err := table.WithContext(ctx).Debug().Create(columnDO)
	if err != nil {
		return nil, err
	}

	return po2do.ConvertColumnDO(columnDO), nil
}

// 按照分页查询 1基
func (receiver ColumnRepo) Page(
	ctx context.Context,
	req query.ColumnPageQuery,
) (list []*do.ColumnDO, pagination *result.Pagination, err error) {

	table := dao.Use(config.DB).ColumnPO
	conditions := make([]gen.Condition, 0)

	if !req.CreateTimeGte.IsZero() {
		conditions = append(conditions, table.CreateTime.Gte(req.CreateTimeGte))
	}

	if req.Name != nil {
		conditions = append(conditions, table.Name.Like("%"+(*req.Name)+"%"))
	}

	if req.AuthorId != nil {
		conditions = append(conditions, table.AuthorID.Eq(*req.AuthorId))
	}

	if req.Status != nil {
		status := do2po.ConvertColumnStatus(*req.Status)
		conditions = append(conditions, table.Status.Eq(status))
	}

	tableDO := table.WithContext(ctx).Debug()
	if len(conditions) > 0 {
		tableDO = tableDO.Where(conditions...)
	}

	//默认按照创建时间倒序排列
	orderExpr := table.CreateTime.Desc()
	if req.OrderBy == query2.OrderByCreateTimeAsc {
		orderExpr = table.CreateTime.Asc()
	} else if req.OrderBy == query2.OrderByIdDesc {
		orderExpr = table.CreateTime.Desc()
	} else if req.OrderBy == query2.OrderByIdAsc {
		orderExpr = table.CreateTime.Asc()
	}
	tableDO = tableDO.Order(orderExpr)

	if req.PageNum <= 0 {
		req.PageNum = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	if req.PageSize > 100 {
		return nil, nil, errs.CodeErrorf(errs.StatusCodeParamsError, "PageSize不能大于100")
	}

	offset := (req.PageNum - 1) * req.PageSize
	limit := req.PageSize
	pageData, total, err := tableDO.FindByPage(int(offset), int(limit))
	if err != nil {
		klog.CtxErrorf(ctx, "PageQuery failed, err=%v", err)
		return nil, nil, err
	}

	pagination = &result.Pagination{
		PageNum:    req.PageNum,
		PageSize:   req.PageSize,
		TotalItems: total,
	}
	return po2do.ConvertColumnDOs(pageData), pagination, nil
}

// 按照id查找，找不到返回nil
func (receiver ColumnRepo) QueryById(
	ctx context.Context,
	columnId int64,
) (*do.ColumnDO, error) {
	table := dao.Use(config.DB).ColumnPO

	conditions := make([]gen.Condition, 0)

	conditions = append(conditions, table.ID.Eq(columnId))

	tableDO := table.WithContext(ctx).Debug()
	if len(conditions) > 0 {
		tableDO = tableDO.Where(conditions...)
	}
	columnDO, err := tableDO.First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 处理未找到记录的错误...
			return nil, nil
		}
		return nil, err
	}
	return po2do.ConvertColumnDO(columnDO), nil
}

// 按照id查找，找不到返回NotFound的Err
func (receiver ColumnRepo) CheckById(
	ctx context.Context,
	columnId int64,
) (*do.ColumnDO, error) {
	column, err := receiver.QueryById(ctx, columnId)
	if err != nil {
		return nil, errs.CodeErrorf(errs.StatusCodeUnknown, err.Error())
	}
	if column == nil {
		return nil, errs.CodeErrorf(errs.StatusCodeNotFound, "没有找到%v对应的记录", columnId)
	}
	return column, nil
}

// 根据id批量查询
func (receiver ColumnRepo) QueryByIds(
	ctx context.Context,
	ids []int64,
) (list []*do.ColumnDO, err error) {

	table := dao.Use(config.DB).ColumnPO
	conditions := make([]gen.Condition, 0)

	if len(ids) > 0 {
		conditions = append(conditions, table.ID.In(ids...))
	} else {
		return nil, errs.CodeErrorf(errs.StatusCodeParamsError, "ids列表不能为空")
	}

	tableDO := table.WithContext(ctx).Debug()
	if len(conditions) > 0 {
		tableDO = tableDO.Where(conditions...)
	}

	listData, err := tableDO.Find()
	if err != nil {
		klog.CtxErrorf(ctx, "QueryByIds failed, err=%v", err)
		return nil, err
	}

	return po2do.ConvertColumnDOs(listData), nil
}
