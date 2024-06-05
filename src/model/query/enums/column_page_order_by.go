package query

// 查询排序
type ColumnPageOrderBy int32

const (
	OrderByCreateTimeDesc ColumnPageOrderBy = 0
	OrderByCreateTimeAsc  ColumnPageOrderBy = 1
	OrderByIdAsc          ColumnPageOrderBy = 2
	OrderByIdDesc         ColumnPageOrderBy = 3
)
