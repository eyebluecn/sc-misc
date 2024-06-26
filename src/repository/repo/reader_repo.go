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
	"github.com/eyebluecn/sc-misc/src/model/result"
	"github.com/eyebluecn/sc-misc/src/repository/config"
	"github.com/eyebluecn/sc-misc/src/repository/dao"
	"gorm.io/gen"
	"gorm.io/gorm"
	"time"
)

type ReaderRepo struct {
}

func NewReaderRepo() ReaderRepo {
	return ReaderRepo{}
}

// 按照分页查询 1基
func (receiver ReaderRepo) Page(
	ctx context.Context,
	req query.ReaderPageQuery,
) (list []*do.ReaderDO, pagination *result.Pagination, err error) {

	table := dao.Use(config.DB).ReaderPO
	conditions := make([]gen.Condition, 0)

	if !req.CreateTimeGte.IsZero() {
		conditions = append(conditions, table.CreateTime.Gte(req.CreateTimeGte))
	}
	if req.Username != "" {
		conditions = append(conditions, table.Username.Eq(req.Username))
	}

	tableDO := table.WithContext(ctx).Debug()
	if len(conditions) > 0 {
		tableDO = tableDO.Where(conditions...)
	}

	if req.PageNum <= 0 {
		req.PageNum = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	if req.PageSize > 100 {
		return nil, nil, errs.BadRequestErrorf("PageSize不能大于100")
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
	return po2do.ConvertReaderDOs(pageData), pagination, nil
}

// 按照用户名查找，找不到返回nil
func (receiver ReaderRepo) QueryByUsername(
	ctx context.Context,
	username string,
) (*do.ReaderDO, error) {
	table := dao.Use(config.DB).ReaderPO

	conditions := make([]gen.Condition, 0)

	conditions = append(conditions, table.Username.Eq(username))

	tableDO := table.WithContext(ctx).Debug()
	if len(conditions) > 0 {
		tableDO = tableDO.Where(conditions...)
	}
	readerDO, err := tableDO.First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 处理未找到记录的错误...
			return nil, nil
		}
		return nil, err
	}
	return po2do.ConvertReaderDO(readerDO), nil
}

// 新建一个Reader
func (receiver ReaderRepo) Insert(
	ctx context.Context,
	reader *do.ReaderDO,
) (*do.ReaderDO, error) {
	table := dao.Use(config.DB).ReaderPO

	//时间置为当前
	reader.CreateTime = time.Now()
	reader.UpdateTime = time.Now()

	readerDO := do2po.ConvertReaderPO(reader)

	err := table.WithContext(ctx).Debug().Create(readerDO)
	if err != nil {
		return nil, err
	}

	return po2do.ConvertReaderDO(readerDO), nil
}

// 按照id查找，找不到返回nil
func (receiver ReaderRepo) QueryById(
	ctx context.Context,
	readerId int64,
) (*do.ReaderDO, error) {
	table := dao.Use(config.DB).ReaderPO

	conditions := make([]gen.Condition, 0)

	conditions = append(conditions, table.ID.Eq(readerId))

	tableDO := table.WithContext(ctx).Debug()
	if len(conditions) > 0 {
		tableDO = tableDO.Where(conditions...)
	}
	readerDO, err := tableDO.First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 处理未找到记录的错误...
			return nil, nil
		}
		return nil, err
	}
	return po2do.ConvertReaderDO(readerDO), nil
}

// 按照id查找，找不到返回NotFound的Err
func (receiver ReaderRepo) CheckById(
	ctx context.Context,
	readerId int64,
) (*do.ReaderDO, error) {
	reader, err := receiver.QueryById(ctx, readerId)
	if err != nil {
		return nil, errs.CodeErrorf(errs.StatusCodeUnknown, err.Error())
	}
	if reader == nil {
		return nil, errs.CodeErrorf(errs.StatusCodeNotFound, "没有找到%v对应的记录", readerId)
	}
	return reader, nil
}

// 根据id批量查询
func (receiver ReaderRepo) QueryByIds(
	ctx context.Context,
	ids []int64,
) (list []*do.ReaderDO, err error) {

	table := dao.Use(config.DB).ReaderPO
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

	return po2do.ConvertReaderDOs(listData), nil
}
