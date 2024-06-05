package application

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc/src/common/errs"
	"github.com/eyebluecn/sc-misc/src/domain"
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/repository/repo"
)

type ReaderReadApp struct{}

func NewReaderReadApp() *ReaderReadApp {
	return &ReaderReadApp{}
}

// 使用用户名和密码进行登录。登录成功就返回对象，失败就报错。
func (receiver ReaderReadApp) Login(ctx context.Context, username string, password string) (*do.Reader, error) {
	reader, err := repo.NewReaderRepo().QueryByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if reader == nil {
		if username == "demo_reader" && password == "123456" {
			//测试用户手动创建。
			reader, err = domain.NewReaderDomainService().Create(ctx, &do.Reader{
				Username: username,
				Password: password,
			})
			if err != nil {
				return nil, err
			}
		}
	}
	//没有找打username对应的记录
	if reader == nil {
		klog.CtxInfof(ctx, "没有找到%v对应的记录", username)
		return nil, errs.CodeErrorf(errs.StatusCodeBadRequest, "用户名或者密码错误")
	}
	//密码错误
	if reader.Password != password {
		//这里明文比较用于示例，真实场景需要加密后用密文比较
		klog.CtxInfof(ctx, "%v对应的密码输入错误", username)
		return nil, errs.CodeErrorf(errs.StatusCodeBadRequest, "用户名或者密码错误")
	}

	return reader, nil
}

// 根据id来查询读者信息
func (receiver ReaderReadApp) QueryById(ctx context.Context, readerId int64) (*do.Reader, error) {
	reader, err := repo.NewReaderRepo().QueryById(ctx, readerId)
	if err != nil {
		return nil, err
	}
	return reader, nil
}

// 根据id来查询读者信息
func (receiver ReaderReadApp) QueryByIds(ctx context.Context, readerIds []int64) ([]*do.Reader, error) {
	readers, err := repo.NewReaderRepo().QueryByIds(ctx, readerIds)
	if err != nil {
		return nil, err
	}
	return readers, nil
}
