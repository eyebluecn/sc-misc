package result

// 分页指示器
type Pagination struct {
	//当前页
	PageNum int64
	//每页大小
	PageSize int64
	//总记录数
	TotalItems int64
}
