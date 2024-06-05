package repo

import (
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc/src/common/errs"
	"github.com/eyebluecn/sc-misc/src/converter/do2po"
	"github.com/eyebluecn/sc-misc/src/converter/po2do"
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/repository/config"
	"github.com/eyebluecn/sc-misc/src/repository/dao"
	"gorm.io/gen"
	"gorm.io/gorm"
	"time"
)

type AuthorRepo struct {
}

func NewAuthorRepo() AuthorRepo {
	return AuthorRepo{}
}

// 按照用户名查找，找不到返回nil
func (receiver AuthorRepo) QueryByUsername(
	ctx context.Context,
	username string,
) (*do.AuthorDO, error) {
	table := dao.Use(config.DB).AuthorPO

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
	return po2do.ConvertAuthorDO(readerDO), nil
}

// 新建一个Author
func (receiver AuthorRepo) Insert(
	ctx context.Context,
	reader *do.AuthorDO,
) (*do.AuthorDO, error) {
	table := dao.Use(config.DB).AuthorPO

	//时间置为当前
	reader.CreateTime = time.Now()
	reader.UpdateTime = time.Now()

	readerDO := do2po.ConvertAuthorPO(reader)

	err := table.WithContext(ctx).Debug().Create(readerDO)
	if err != nil {
		return nil, err
	}

	return po2do.ConvertAuthorDO(readerDO), nil
}

// 根据id批量查询
func (receiver AuthorRepo) QueryByIds(
	ctx context.Context,
	ids []int64,
) (list []*do.AuthorDO, err error) {

	table := dao.Use(config.DB).AuthorPO
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

	return po2do.ConvertAuthorDOs(listData), nil
}
