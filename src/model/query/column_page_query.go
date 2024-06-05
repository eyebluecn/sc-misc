package query

import (
	"github.com/eyebluecn/sc-misc/src/model/do/enums"
	query "github.com/eyebluecn/sc-misc/src/model/query/enums"
	"time"
)

type ColumnPageQuery struct {

	//创建时间晚于
	CreateTimeGte time.Time
	//标题
	Name *string
	//作者
	AuthorId *int64
	//状态
	Status *enums.ColumnStatus

	//按照时间排序
	OrderBy query.ColumnPageOrderBy

	//当前页 1基
	PageNum int64
	//每页大小
	PageSize int64
}
