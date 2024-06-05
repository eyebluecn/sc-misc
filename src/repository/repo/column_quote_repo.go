package repo

import (
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc/src/common/config"
	"github.com/eyebluecn/sc-misc/src/common/errs"
	"github.com/eyebluecn/sc-misc/src/converter/db_model_conv"
	"github.com/eyebluecn/sc-misc/src/converter/model_conv"
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/do/enums"
	"github.com/eyebluecn/sc-misc/src/repository/dao"
	"gorm.io/gen"
	"gorm.io/gorm"
	"time"
)

type ColumnQuoteRepo struct {
}

func NewColumnQuoteRepo() ColumnQuoteRepo {
	return ColumnQuoteRepo{}
}

// 新建一个ColumnQuote
func (receiver ColumnQuoteRepo) Insert(
	ctx context.Context,
	reader *do.ColumnQuote,
) (*do.ColumnQuote, error) {
	table := dao.Use(config.DB).ColumnQuotePO

	//时间置为当前
	reader.CreateTime = time.Now()
	reader.UpdateTime = time.Now()

	readerDO := db_model_conv.ConvertColumnQuoteDO(reader)

	err := table.WithContext(ctx).Debug().Create(readerDO)
	if err != nil {
		return nil, err
	}

	return model_conv.ConvertColumnQuote(readerDO), nil
}

// 根据id批量查询
func (receiver ColumnQuoteRepo) FindByIds(
	ctx context.Context,
	ids []int64,
) (list []*do.ColumnQuote, err error) {

	table := dao.Use(config.DB).ColumnQuotePO
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
		klog.CtxErrorf(ctx, "FindByIds failed, err=%v", err)
		return nil, err
	}

	return model_conv.ConvertColumnQuotes(listData), nil
}

// 根据专栏id批量查询生效的
func (receiver ColumnQuoteRepo) FindOkByColumnIds(
	ctx context.Context,
	columnIds []int64,
) (list []*do.ColumnQuote, err error) {

	table := dao.Use(config.DB).ColumnQuotePO
	conditions := make([]gen.Condition, 0)

	conditions = append(conditions, table.Status.Eq(db_model_conv.ColumnQuoteStatusToStorage(enums.ColumnQuoteStatusOk)))
	if len(columnIds) > 0 {
		conditions = append(conditions, table.ColumnID.In(columnIds...))
	} else {
		return nil, errs.CodeErrorf(errs.StatusCodeParamsError, "ids列表不能为空")
	}

	tableDO := table.WithContext(ctx).Debug()
	if len(conditions) > 0 {
		tableDO = tableDO.Where(conditions...)
	}

	listData, err := tableDO.Find()
	if err != nil {
		klog.CtxErrorf(ctx, "FindByIds failed, err=%v", err)
		return nil, err
	}

	return model_conv.ConvertColumnQuotes(listData), nil
}

// 按照专栏id查找，找不到返回nil
func (receiver ColumnQuoteRepo) QueryByColumnId(
	ctx context.Context,
	columnId int64,
) (*do.ColumnQuote, error) {
	table := dao.Use(config.DB).ColumnQuotePO

	conditions := make([]gen.Condition, 0)

	conditions = append(conditions, table.ColumnID.Eq(columnId))

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
	return model_conv.ConvertColumnQuote(readerDO), nil
}
