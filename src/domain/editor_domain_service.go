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

type EditorDomainService struct{}

func NewEditorDomainService() *EditorDomainService {
	return &EditorDomainService{}
}

// 新增用户
func (receiver EditorDomainService) Create(ctx context.Context, editor *model.Editor) (*model.Editor, error) {

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
	_ = mq.NewProducer().Publish(ctx, mq.MqTopicEditor, "editor", "EDITOR_CREATE", util.ToJSON(editor))

	return editor, nil
}

func (receiver EditorDomainService) createParamCheck(ctx context.Context, editor *model.Editor) error {

	//参数校验。
	if editor == nil {
		return errs.CodeErrorf(enums.StatusCodeParamsError, "editor is nil")
	}
	if editor.Username == "" {
		return errs.CodeErrorf(enums.StatusCodeParamsError, "用户名不能为空")
	}
	if editor.Password == "" {
		return errs.CodeErrorf(enums.StatusCodeParamsError, "密码不能为空")
	}
	return nil

}
