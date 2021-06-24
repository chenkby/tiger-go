package service

import (
	"context"
	"github.com/pkgng/copit"
	"tiger-go/applications/admin-api/internal/define"
	"tiger-go/modules/video/dao"
	"tiger-go/modules/video/model"
	"tiger-go/modules/video/service"
	"tiger-go/tiger/data"
	"tiger-go/tiger/util/convert"
)

var Video = new(videoService)

type videoService struct {
	service.VideoService
}

func (s *videoService) List(ctx context.Context, req *define.VideoListReq) (*define.VideoListResp, error) {
	query := dao.Video
	// @todo 链式操作返回对象不为db.DB(bug)
	if req.Keyword != "" {
		query.Where("title = ?", req.Keyword)
	}
	provider := data.Provider{
		Query:    query,
		Page:     req.Page,
		PageSize: req.PageSize,
		OrderBy:  "video_id desc",
	}
	result := ([]*define.VideoViewResp)(nil)
	paging, err := provider.Find(&result)
	return &define.VideoListResp{
		Data:   result,
		Paging: paging,
	}, err
}

// 添加记录，返回新增的数据
func (s *videoService) Create(ctx context.Context, req *define.VideoCreateReq) (*define.VideoViewResp, error) {
	var insertData = model.Video{}
	if err := copit.Copy(&insertData, &req); err != nil {
		return nil, err
	}
	err := dao.Video.Create(&insertData).Error
	if err != nil {
		return nil, err
	}
	var result = define.VideoViewResp{}
	if err := convert.Copy(insertData, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// 更新记录，返回更新后的数据
func (s *videoService) Update(ctx context.Context, req *define.VideoUpdateReq) (*define.VideoViewResp, error) {
	err := dao.Video.Select("*").Omit(dao.Video.PrimaryKey).Where(req.VideoId).Updates(&req).Error
	if err != nil {
		return nil, err
	}
	var result = define.VideoViewResp{}
	if err := convert.Copy(&req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// 批量删除记录，返回成功删除的记录数
func (s *videoService) Delete(ctx context.Context, req *define.VideoDeleteReq) (int64, error) {
	db := dao.Video.Delete(&model.Video{}, req.Ids)
	if err := db.Error; err != nil {
		return 0, err
	} else {
		return db.RowsAffected, nil
	}
}

// 返回单条数据
func (s *videoService) View(ctx context.Context, req *define.VideoViewReq) (*define.VideoViewResp, error) {
	var result *define.VideoViewResp
	err := dao.Video.First(&result, "video_id=?", req.VideoId).Error
	if err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

// 批量更新单个字段的值，返回影响的记录数
func (s *videoService) Toggle(ctx context.Context, req *define.VideoToggleReq) (int64, error) {
	db := dao.Video.Where("video_id in ?", req.Ids).Updates(map[string]interface{}{
		req.Attribute: req.Value,
	})
	if err := db.Error; err != nil {
		return 0, err
	} else {
		return db.RowsAffected, nil
	}
}
