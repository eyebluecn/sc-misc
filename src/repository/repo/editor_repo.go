package repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc/src/common/config"
	"github.com/eyebluecn/sc-misc/src/common/errs"
	"github.com/eyebluecn/sc-misc/src/converter/db_model_conv"
	"github.com/eyebluecn/sc-misc/src/converter/model_conv"
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/universal"
	"github.com/eyebluecn/sc-misc/src/repository/dao"
	"gorm.io/gen"
	"gorm.io/gorm"
	"time"
)

type EditorRepo struct {
}

func NewEditorRepo() EditorRepo {
	return EditorRepo{}
}

// 按照分页查询 1基
func (receiver EditorRepo) Page(
	ctx context.Context,
	req EditorPageRequest,
) (list []*do.EditorDO, pagination *universal.Pagination, err error) {

	table := dao.Use(config.DB).EditorPO
	conditions := make([]gen.Condition, 0)

	if !req.CreateTimeGte.IsZero() {
		conditions = append(conditions, table.CreateTime.Gte(req.CreateTimeGte))
	}
	if req.Username != "" {
		conditions = append(conditions, table.Username.Eq(req.Username))
	}

	tableDO := table.WithContext(ctx).Debug()
	if len(conditions) > 0 {
		tableDO = tableDO.Where(conditions...)
	}

	if req.PageNum <= 0 {
		req.PageNum = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	if req.PageSize > 100 {
		return nil, nil, errs.BadRequestErrorf("PageSize不能大于100")
	}

	offset := (req.PageNum - 1) * req.PageSize
	limit := req.PageSize
	pageData, total, err := tableDO.FindByPage(int(offset), int(limit))
	if err != nil {
		klog.CtxErrorf(ctx, "PageQuery failed, err=%v", err)
		return nil, nil, err
	}

	pagination = &universal.Pagination{
		PageNum:    req.PageNum,
		PageSize:   req.PageSize,
		TotalItems: total,
	}
	return model_conv.ConvertEditors(pageData), pagination, nil
}

// 按照用户名查找，找不到返回nil
func (receiver EditorRepo) FindByUsername(
	ctx context.Context,
	username string,
) (*do.EditorDO, error) {
	table := dao.Use(config.DB).EditorPO

	conditions := make([]gen.Condition, 0)

	conditions = append(conditions, table.Username.Eq(username))

	tableDO := table.WithContext(ctx).Debug()
	if len(conditions) > 0 {
		tableDO = tableDO.Where(conditions...)
	}
	editorDO, err := tableDO.First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 处理未找到记录的错误...
			return nil, nil
		}
		return nil, err
	}
	return model_conv.ConvertEditor(editorDO), nil
}

// 按照id查找，找不到返回nil
func (receiver EditorRepo) FindById(
	ctx context.Context,
	id int64,
) (*do.EditorDO, error) {
	table := dao.Use(config.DB).EditorPO

	conditions := make([]gen.Condition, 0)

	conditions = append(conditions, table.ID.Eq(id))

	tableDO := table.WithContext(ctx).Debug()
	if len(conditions) > 0 {
		tableDO = tableDO.Where(conditions...)
	}
	editorDO, err := tableDO.First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 处理未找到记录的错误...
			return nil, nil
		}
		return nil, err
	}
	return model_conv.ConvertEditor(editorDO), nil
}

// 按照id查找，找不到返回err
func (receiver EditorRepo) CheckById(
	ctx context.Context,
	id int64,
) (*do.EditorDO, error) {
	editor, err := receiver.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	if editor == nil {
		return nil, errs.CodeErrorf(errs.StatusCodeNotFound, fmt.Sprintf("id=%d对应的小编不存在", id))
	}
	return editor, nil
}

// 新建一个Editor
func (receiver EditorRepo) Insert(
	ctx context.Context,
	editor *do.EditorDO,
) (*do.EditorDO, error) {
	table := dao.Use(config.DB).EditorPO

	//时间置为当前
	editor.CreateTime = time.Now()
	editor.UpdateTime = time.Now()

	editorDO := db_model_conv.ConvertEditorDO(editor)

	err := table.WithContext(ctx).Debug().Create(editorDO)
	if err != nil {
		return nil, err
	}

	return model_conv.ConvertEditor(editorDO), nil
}
