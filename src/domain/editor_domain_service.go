package domain

import (
	"context"
	"github.com/eyebluecn/sc-misc/src/common/errs"
	"github.com/eyebluecn/sc-misc/src/common/util"
	mq2 "github.com/eyebluecn/sc-misc/src/infrastructure/middleware/mq"
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/repository/repo"
)

type EditorDomainService struct{}

func NewEditorDomainService() *EditorDomainService {
	return &EditorDomainService{}
}

// 新增用户
func (receiver EditorDomainService) Create(ctx context.Context, editor *do.Editor) (*do.Editor, error) {

	//参数校验
	err := receiver.createParamCheck(ctx, editor)
	if err != nil {
		return nil, err
	}

	editorRepo := repo.NewEditorRepo()
	editor, err = editorRepo.Insert(ctx, editor)
	if err != nil {
		return nil, err
	}

	//发出领域事件
	_ = mq2.NewProducer().Publish(ctx, mq2.MqTopicEditor, "editor", "EDITOR_CREATE", util.ToJSON(editor))

	return editor, nil
}

func (receiver EditorDomainService) createParamCheck(ctx context.Context, editor *do.Editor) error {

	//参数校验。
	if editor == nil {
		return errs.CodeErrorf(errs.StatusCodeParamsError, "editor is nil")
	}
	if editor.Username == "" {
		return errs.CodeErrorf(errs.StatusCodeParamsError, "用户名不能为空")
	}
	if editor.Password == "" {
		return errs.CodeErrorf(errs.StatusCodeParamsError, "密码不能为空")
	}
	return nil

}
