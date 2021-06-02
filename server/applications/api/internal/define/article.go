package define

import (
	_ "tiger-go/modules/article/model"
	"tiger-go/tiger/data"
)

// 分页列表请求参数
type ArticleListReq struct {
	Keyword  string `json:"keyword" form:"keyword"`
	Field    string `json:"field" form:"field"`
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"pageSize" form:"pageSize"`
	OrderBy  string `json:"orderBy" form:"orderBy"`
}

// 分页请求响应数据
type ArticleListResp struct {
	Data   []*ArticleViewResp `json:"data"`
	Paging *data.Paging       `json:"paging"`
}

// 创建请求参数
type ArticleCreateReq struct {
	Title    string `json:"title" form:"title" binding:"required"`
	Content  string `json:"content" form:"content"`
	Template int    `json:"template" form:"template"`
	Status   int    `json:"status"`
}

// 修改请求参数
type ArticleUpdateReq struct {
	ArticleId int `json:"article_id" form:"article_id" binding:"required"`
	ArticleCreateReq
}

// 删除请求参数
type ArticleDeleteReq struct {
	Ids []int `json:"ids" form:"ids" binding:"required"`
}

// 查看请求参数
type ArticleViewReq struct {
	ArticleId int `form:"article_id" binding:"required|min=1"`
}

// 修改字段值请求参数
type ArticleToggleReq struct {
	Ids       []int  `json:"ids" form:"ids" binding:"required"`
	Attribute string `json:"attribute" form:"attribute" binding:"required"`
	Value     int    `json:"value" form:"value" binding:"min=0"`
}

// 查看响应数据
type ArticleViewResp struct {
	ArticleId int    `json:"article_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Template  int    `json:"template"`
	ViewNum   int    `json:"view_num"`
}
