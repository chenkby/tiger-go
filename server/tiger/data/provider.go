package data

import (
	"gorm.io/gorm"
)

type Provider struct {
	Query    *gorm.DB
	Page     int
	PageSize int
	OrderBy  string
}

func (provider *Provider) Find(dest interface{}) (*Paging, error) {
	query := provider.Query
	var totalCount int64

	query.Count(&totalCount)
	if provider.PageSize <= 0 {
		provider.PageSize = 10
	}
	paging := NewPagination(provider.Page, provider.PageSize, totalCount)
	if totalCount <= 0 {
		return paging, nil
	}

	query = query.Limit(paging.PageSize).Offset(paging.Offset)

	// @todo parse orderby
	if provider.OrderBy != "" {
		query = query.Order(provider.OrderBy)
	}

	if err := query.Find(dest).Error; err != nil {
		return paging, err
	} else {
		return paging, nil
	}
}
