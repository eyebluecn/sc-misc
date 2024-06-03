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

type ReaderDomainService struct{}

func NewReaderDomainService() *ReaderDomainService {
	return &ReaderDomainService{}
}

// 新增用户
func (receiver ReaderDomainService) Create(ctx context.Context, reader *model.Reader) (*model.Reader, error) {

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
	_ = mq.NewProducer().Publish(ctx, mq.MqTopicReader, "reader", "READER_CREATE", util.ToJSON(reader))

	return reader, nil
}

func (receiver ReaderDomainService) createParamCheck(ctx context.Context, reader *model.Reader) error {

	//参数校验。
	if reader == nil {
		return errs.CodeErrorf(enums.StatusCodeParamsError, "reader is nil")
	}
	if reader.Username == "" {
		return errs.CodeErrorf(enums.StatusCodeParamsError, "用户名不能为空")
	}
	if reader.Password == "" {
		return errs.CodeErrorf(enums.StatusCodeParamsError, "密码不能为空")
	}
	return nil

}
