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

type AuthorDomainService struct{}

func NewAuthorDomainService() *AuthorDomainService {
	return &AuthorDomainService{}
}

// 新增用户
func (receiver AuthorDomainService) Create(ctx context.Context, author *model.Author) (*model.Author, error) {

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
	_ = mq.NewProducer().Publish(ctx, mq.MqTopicAuthor, "author", "AUTHOR_CREATE", util.ToJSON(author))

	return author, nil
}

func (receiver AuthorDomainService) createParamCheck(ctx context.Context, author *model.Author) error {

	//参数校验。
	if author == nil {
		return errs.CodeErrorf(enums.StatusCodeParamsError, "author is nil")
	}
	if author.Username == "" {
		return errs.CodeErrorf(enums.StatusCodeParamsError, "用户名不能为空")
	}
	if author.Password == "" {
		return errs.CodeErrorf(enums.StatusCodeParamsError, "密码不能为空")
	}
	return nil

}
