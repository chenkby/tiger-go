package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tiger-go/applications/admin-api/internal/define"
	"tiger-go/applications/admin-api/internal/service"
	"tiger-go/tiger/response"
)

var Video = new(videoApi)

type videoApi struct {
}

// 文章列表
// @Summary 文章列表
// @Description 这个是接口的额外说明
// @Tags 文章
// @Accept json
// @Produce json
// @Param group_id path int true "Group ID"
// @Param account_id path int true "Account ID"
// @Success 200 {string} string "answer"
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /v1/video/list/{account_id} [get]
func (*videoApi) List(c *gin.Context) {
	var req define.VideoListReq
	if err := c.ShouldBind(&req); err != nil {
		response.Error(c, err)
		return
	}
	if result, err := service.Video.List(c, &req); err != nil {
		response.Error(c, err)
	} else {
		response.PageSuccess(c, result.Data, result.Paging)
	}
}

func (*videoApi) Create(c *gin.Context) {
	var req define.VideoCreateReq
	if err := c.ShouldBind(&req); err != nil {
		response.Error(c, err)
		return
	}
	if result, err := service.Video.Create(nil, &req); err != nil {
		response.Error(c, err)
	} else {
		response.Success(c, result)
	}
}

func (*videoApi) Update(c *gin.Context) {
	var req define.VideoUpdateReq
	if err := c.ShouldBind(&req); err != nil {
		response.Error(c, err)
		return
	}
	if result, err := service.Video.Update(nil, &req); err != nil {
		response.Error(c, err)
	} else {
		response.Success(c, result)
	}
}

func (*videoApi) Delete(c *gin.Context) {
	var req define.VideoDeleteReq
	if err := c.ShouldBind(&req); err != nil {
		response.Error(c, err)
		return
	}
	fmt.Println("ids:", req.Ids)
	if result, err := service.Video.Delete(c, &req); err != nil {
		response.Error(c, err)
	} else {
		response.Success(c, result)
	}
}

func (*videoApi) Info(c *gin.Context) {
	var req define.VideoInfoReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Error(c, err)
		return
	}
	result, err := service.Video.Info(c, &req)
	if err != nil {
		response.Error(c, err)
	} else {
		response.Success(c, result)
	}
}

func (*videoApi) View(c *gin.Context) {
	var req define.VideoViewReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Error(c, err)
		return
	}
	result, err := service.Video.View(c, &req)
	if err != nil {
		response.Error(c, err)
	} else {
		response.Success(c, result)
	}
}

func (*videoApi) Toggle(c *gin.Context) {
	var req define.VideoToggleReq
	if err := c.ShouldBind(&req); err != nil {
		response.Error(c, err)
		return
	}
	result, err := service.Video.Toggle(c, &req)
	fmt.Println(result, err)
	if err != nil {
		response.Error(c, err)
	} else {
		response.Success(c, result)
	}
}