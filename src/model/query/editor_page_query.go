package query

import (
	"time"
)

type EditorPageQuery struct {

	//创建时间晚于
	CreateTimeGte time.Time

	//用户名
	Username string

	//当前页 1基
	PageNum int64
	//每页大小
	PageSize int64
}
