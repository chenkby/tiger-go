package data

import "math"

type Paging struct {
	CurrentPage int   `json:"currentPage"` // 当前页码
	PageSize    int   `json:"pageSize"`    // 每页显示多少条
	TotalCount  int64 `json:"totalCount"`  // 总条数
	PageCount   int   `json:"pageCount"`   // 总共多少页
	Offset      int   `json:"-"`           // 起始页
}

// 分页
func NewPagination(page int, pageSize int, total int64) *Paging {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	pageCount := int(math.Ceil(float64(total) / float64(pageSize)))

	if page > pageCount {
		page = pageCount
	}
	offset := pageSize * (page - 1)
	return &Paging{
		CurrentPage: page,
		PageSize:    pageSize,
		TotalCount:  total,
		PageCount:   pageCount,
		Offset:      offset,
	}
}
