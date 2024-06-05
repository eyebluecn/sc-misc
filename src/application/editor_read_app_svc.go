package application

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc/src/common/errs"
	"github.com/eyebluecn/sc-misc/src/domain"
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/repository/repo"
)

type EditorReadAppSvc struct{}

func NewEditorReadAppSvc() *EditorReadAppSvc {
	return &EditorReadAppSvc{}
}

// 使用用户名和密码进行登录。登录成功就返回对象，失败就报错。
func (r EditorReadAppSvc) Login(ctx context.Context, username string, password string) (*do.EditorDO, error) {
	editor, err := repo.NewEditorRepo().QueryByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if editor == nil {
		if username == "demo_editor" && password == "123456" {
			//测试用户手动创建。
			editor, err = domain.NewEditorDomainSvc().Create(ctx, &do.EditorDO{
				Username: username,
				Password: password,
			})
			if err != nil {
				return nil, err
			}
		}
	}
	//没有找打username对应的记录
	if editor == nil {
		klog.CtxInfof(ctx, "没有找到%v对应的记录", username)
		return nil, errs.CodeErrorf(errs.StatusCodeBadRequest, "用户名或者密码错误")
	}
	//密码错误
	if editor.Password != password {
		//这里明文比较用于示例，真实场景需要加密后用密文比较
		klog.CtxInfof(ctx, "%v对应的密码输入错误", username)
		return nil, errs.CodeErrorf(errs.StatusCodeBadRequest, "用户名或者密码错误")
	}

	return editor, nil
}
