package define

import (
	"gorm.io/gorm"
	_ "tiger-go/modules/video/model"
	"tiger-go/tiger/data"
)

// 分页列表请求参数
type VideoListReq struct {
	Keyword  string `json:"keyword" form:"keyword"`
	Field    string `json:"field" form:"field"`
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"pageSize" form:"pageSize"`
	OrderBy  string `json:"orderBy" form:"orderBy"`
}

// 分页请求响应数据
type VideoListResp struct {
	Data   []*VideoViewResp `json:"data"`
	Paging *data.Paging     `json:"paging"`
}

// 创建请求参数
type VideoCreateReq struct {
	Name       string `json:"name" form:"name" binding:"required"`
	CategoryId int    `json:"category_id" form:"category_id"`
	Template   int    `json:"template" form:"template"`
	Status     int    `json:"status"`
}

// 修改请求参数
type VideoUpdateReq struct {
	VideoId int `json:"article_id" form:"article_id" binding:"required"`
	VideoCreateReq
}

// 查看model请求参数
type VideoInfoReq struct {
	VideoId int `form:"id" binding:"required|min=1"`
}

// 查看model响应数据
type VideoInfoResp struct {
	VideoId   int    `json:"video_id" from:"video_id"`
	Name      string `json:"name" form:"name"`
	Image     string `json:"image" from:"image"`
	Video     string `json:"video" form:"video"`
	PlayNum   int    `json:"play_num" form:"play_num"`
	Status    int    `json:"status" form:"status"`
	SizeByte  int    `json:"size_byte" form:"size_byte"`
	Size      string `json:"size" form:"size"`
	Free      int    `json:"free" form:"free"`
	Rank      int    `json:"rank" form:"rank"`
	CreatedAt int    `json:"created_at" form:"created_at"`
	UpdatedAt int    `json:"updated_at" form:"updated_at"`
}

// 删除请求参数
type VideoDeleteReq struct {
	Ids []int `json:"ids" form:"ids" binding:"required"`
}

// 查看请求参数
type VideoViewReq struct {
	VideoId int `form:"article_id" binding:"required|min=1"`
}

func (m *VideoViewResp) AfterFind(tx *gorm.DB) (err error) {

	if m.Image != "" {
		m.Image = "https://xx.com" + m.Image
	}

	return nil
}

// 修改字段值请求参数
type VideoToggleReq struct {
	Ids       []int  `json:"ids" form:"ids" binding:"required"`
	Attribute string `json:"attribute" form:"attribute" binding:"required"`
	Value     int    `json:"value" form:"value" binding:"min=0"`
}

// 查看响应数据
type VideoViewResp struct {
	VideoId   int    `json:"video_id" from:"video_id"`
	Name      string `json:"name" form:"name"`
	Image     string `json:"image" from:"image"`
	Video     string `json:"video" form:"video"`
	PlayNum   int    `json:"play_num" form:"play_num"`
	Status    int    `json:"status" form:"status"`
	SizeByte  int    `json:"size_byte" form:"size_byte"`
	Size      string `json:"size" form:"size"`
	Free      int    `json:"free" form:"free"`
	Rank      int    `json:"rank" form:"rank"`
	CreatedAt int    `json:"created_at" form:"created_at"`
	UpdatedAt int    `json:"updated_at" form:"updated_at"`
}
