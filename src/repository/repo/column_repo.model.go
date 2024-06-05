package repo

import (
	"github.com/eyebluecn/sc-misc/src/model/do/enums"
	"time"
)

type ColumnPageRequest struct {

	//创建时间晚于
	CreateTimeGte time.Time
	//标题
	Name *string
	//作者
	AuthorId *int64
	//状态
	Status *enums.ColumnStatus

	//按照时间排序
	OrderBy ColumnPageOrderBy

	//当前页 1基
	PageNum int64
	//每页大小
	PageSize int64
}

// 查询排序
type ColumnPageOrderBy int32

const (
	OrderByCreateTimeDesc ColumnPageOrderBy = 0
	OrderByCreateTimeAsc  ColumnPageOrderBy = 1
	OrderByIdAsc          ColumnPageOrderBy = 2
	OrderByIdDesc         ColumnPageOrderBy = 3
)
