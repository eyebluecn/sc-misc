package domain

import (
	"context"
	"github.com/eyebluecn/sc-misc/src/common/errs"
	"github.com/eyebluecn/sc-misc/src/common/util"
	mq2 "github.com/eyebluecn/sc-misc/src/infrastructure/middleware/mq"
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/repository/repo"
)

type ReaderDomainService struct{}

func NewReaderDomainService() *ReaderDomainService {
	return &ReaderDomainService{}
}

// 新增用户
func (receiver ReaderDomainService) Create(ctx context.Context, reader *do.ReaderDO) (*do.ReaderDO, error) {

	//参数校验
	err := receiver.createParamCheck(ctx, reader)
	if err != nil {
		return nil, err
	}

	readerRepo := repo.NewReaderRepo()
	reader, err = readerRepo.Insert(ctx, reader)
	if err != nil {
		return nil, err
	}

	//发出领域事件
	_ = mq2.NewProducer().Publish(ctx, mq2.MqTopicReader, "reader", "READER_CREATE", util.ToJSON(reader))

	return reader, nil
}

func (receiver ReaderDomainService) createParamCheck(ctx context.Context, reader *do.ReaderDO) error {

	//参数校验。
	if reader == nil {
		return errs.CodeErrorf(errs.StatusCodeParamsError, "reader is nil")
	}
	if reader.Username == "" {
		return errs.CodeErrorf(errs.StatusCodeParamsError, "用户名不能为空")
	}
	if reader.Password == "" {
		return errs.CodeErrorf(errs.StatusCodeParamsError, "密码不能为空")
	}
	return nil

}
