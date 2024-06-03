package domain

import (
	"context"
	"github.com/eyebluecn/sc-misc/src/common/enums"
	"github.com/eyebluecn/sc-misc/src/common/errs"
	"github.com/eyebluecn/sc-misc/src/infra/mq"
	"github.com/eyebluecn/sc-misc/src/model"
	"github.com/eyebluecn/sc-misc/src/repository/repo"
	"github.com/eyebluecn/sc-misc/src/util"
)

type ColumnDomainService struct{}

func NewColumnDomainService() *ColumnDomainService {
	return &ColumnDomainService{}
}

// 新增用户
func (receiver ColumnDomainService) Create(ctx context.Context, column *model.Column, author *model.Author) (*model.Column, error) {

	//参数校验
	err := receiver.createParamCheck(ctx, column, author)
	if err != nil {
		return nil, err
	}

	columnRepo := repo.NewColumnRepo()
	column, err = columnRepo.Insert(ctx, column)
	if err != nil {
		return nil, err
	}

	//发出领域事件
	_ = mq.NewProducer().Publish(ctx, mq.MqTopicColumn, "column", "COLUMN_CREATE", util.ToJSON(column))

	return column, nil
}

func (receiver ColumnDomainService) createParamCheck(ctx context.Context, column *model.Column, author *model.Author) error {

	//参数校验。
	if column == nil {
		return errs.CodeErrorf(enums.StatusCodeParamsError, "column is nil")
	}
	if author == nil {
		return errs.CodeErrorf(enums.StatusCodeParamsError, "author is nil")
	}
	if column.Name == "" {
		return errs.CodeErrorf(enums.StatusCodeParamsError, "用户名不能为空")
	}
	if column.AuthorID != author.ID {
		return errs.CodeErrorf(enums.StatusCodeParamsError, "作者信息不一致")
	}

	return nil

}
