package domain

import (
	"context"
	"github.com/eyebluecn/sc-misc/src/common/errs"
	"github.com/eyebluecn/sc-misc/src/common/util"
	mq2 "github.com/eyebluecn/sc-misc/src/infrastructure/middleware/mq"
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/repository/repo"
)

type AuthorDomainService struct{}

func NewAuthorDomainService() *AuthorDomainService {
	return &AuthorDomainService{}
}

// 新增用户
func (receiver AuthorDomainService) Create(ctx context.Context, author *do.AuthorDO) (*do.AuthorDO, error) {

	//参数校验
	err := receiver.createParamCheck(ctx, author)
	if err != nil {
		return nil, err
	}

	authorRepo := repo.NewAuthorRepo()
	author, err = authorRepo.Insert(ctx, author)
	if err != nil {
		return nil, err
	}

	//发出领域事件
	_ = mq2.NewProducer().Publish(ctx, mq2.MqTopicAuthor, "author", "AUTHOR_CREATE", util.ToJSON(author))

	return author, nil
}

func (receiver AuthorDomainService) createParamCheck(ctx context.Context, author *do.AuthorDO) error {

	//参数校验。
	if author == nil {
		return errs.CodeErrorf(errs.StatusCodeParamsError, "author is nil")
	}
	if author.Username == "" {
		return errs.CodeErrorf(errs.StatusCodeParamsError, "用户名不能为空")
	}
	if author.Password == "" {
		return errs.CodeErrorf(errs.StatusCodeParamsError, "密码不能为空")
	}
	return nil

}
