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
func (receiver AuthorRepo) FindByUsername(
	ctx context.Context,
	username string,
) (*do.Author, error) {
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
	return model_conv.ConvertAuthor(readerDO), nil
}

// 新建一个Author
func (receiver AuthorRepo) Insert(
	ctx context.Context,
	reader *do.Author,
) (*do.Author, error) {
	table := dao.Use(config.DB).AuthorPO

	//时间置为当前
	reader.CreateTime = time.Now()
	reader.UpdateTime = time.Now()

	readerDO := db_model_conv.ConvertAuthorDO(reader)

	err := table.WithContext(ctx).Debug().Create(readerDO)
	if err != nil {
		return nil, err
	}

	return model_conv.ConvertAuthor(readerDO), nil
}

// 根据id批量查询
func (receiver AuthorRepo) FindByIds(
	ctx context.Context,
	ids []int64,
) (list []*do.Author, err error) {

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
		klog.CtxErrorf(ctx, "FindByIds failed, err=%v", err)
		return nil, err
	}

	return model_conv.ConvertAuthors(listData), nil
}
