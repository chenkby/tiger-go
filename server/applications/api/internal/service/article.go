package service

import (
	"context"
	"github.com/pkgng/copit"
	"tiger-go/applications/api/internal/define"
	"tiger-go/modules/article/dao"
	"tiger-go/modules/article/model"
	"tiger-go/modules/article/services"
	"tiger-go/tiger/data"
	"tiger-go/tiger/util/convert"
)

var Article = new(articleService)

type articleService struct {
	services.ArticleService
}

func (s *articleService) List(ctx context.Context, req *define.ArticleListReq) (*define.ArticleListResp, error) {
	query := dao.Article
	provider := data.Provider{
		Query:    query,
		Page:     req.Page,
		PageSize: req.PageSize,
		OrderBy:  "article_id desc",
	}
	result := ([]*define.ArticleViewResp)(nil)
	paging, err := provider.Find(&result)
	return &define.ArticleListResp{
		Data:   result,
		Paging: paging,
	}, err
}

// 添加记录，返回新增的数据
func (s *articleService) Create(ctx context.Context, req *define.ArticleCreateReq) (*define.ArticleViewResp, error) {
	var insertData = model.Article{}
	if err := copit.Copy(&insertData, &req); err != nil {
		return nil, err
	}
	err := dao.Article.Create(&insertData).Error
	if err != nil {
		return nil, err
	}
	var result = define.ArticleViewResp{}
	if err := convert.Copy(insertData, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// 更新记录，返回更新后的数据
func (s *articleService) Update(ctx context.Context, req *define.ArticleUpdateReq) (*define.ArticleViewResp, error) {
	err := dao.Article.Select("*").Omit(dao.Article.PrimaryKey).Where(req.ArticleId).Updates(&req).Error
	if err != nil {
		return nil, err
	}
	var result = define.ArticleViewResp{}
	if err := convert.Copy(&req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// 批量删除记录，返回成功删除的记录数
func (s *articleService) Delete(ctx context.Context, req *define.ArticleDeleteReq) (int64, error) {
	db := dao.Article.Delete(&model.Article{}, req.Ids)
	if err := db.Error; err != nil {
		return 0, err
	} else {
		return db.RowsAffected, nil
	}
}

// 返回单条数据
func (s *articleService) View(ctx context.Context, req *define.ArticleViewReq) (*define.ArticleViewResp, error) {
	var result *define.ArticleViewResp
	err := dao.Article.First(&result, "article_id=?", req.ArticleId).Error
	if err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

// 批量更新单个字段的值，返回影响的记录数
func (s *articleService) Toggle(ctx context.Context, req *define.ArticleToggleReq) (int64, error) {
	db := dao.Article.Where("article_id in ?", req.Ids).Updates(map[string]interface{}{
		req.Attribute: req.Value,
	})
	if err := db.Error; err != nil {
		return 0, err
	} else {
		return db.RowsAffected, nil
	}
}
