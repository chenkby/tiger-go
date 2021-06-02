package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tiger-go/applications/api/internal/define"
	"tiger-go/applications/api/internal/service"
	"tiger-go/tiger/response"
)

var Article = new(articleController)

type articleController struct {
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
// @Router /v1/article/list/{account_id} [get]
func (*articleController) List(c *gin.Context) {
	var req define.ArticleListReq
	if err := c.ShouldBind(&req); err != nil {
		response.Error(c, err)
	}
	if result, err := service.Article.List(c, &req); err != nil {
		response.Error(c, err)
	} else {
		response.PageSuccess(c, result.Data, result.Paging)
	}
}

func (*articleController) Create(c *gin.Context) {
	var req define.ArticleCreateReq
	if err := c.ShouldBind(&req); err != nil {
		response.Error(c, err)
	}
	if result, err := service.Article.Create(nil, &req); err != nil {
		response.Error(c, err)
	} else {
		response.Success(c, result)
	}
}

func (*articleController) Update(c *gin.Context) {
	var req define.ArticleUpdateReq
	if err := c.ShouldBind(&req); err != nil {
		response.Error(c, err)
		return
	}
	if result, err := service.Article.Update(nil, &req); err != nil {
		response.Error(c, err)
	} else {
		response.Success(c, result)
	}
}

func (*articleController) Delete(c *gin.Context) {
	var req define.ArticleDeleteReq
	if err := c.ShouldBind(&req); err != nil {
		response.Error(c, err)
		return
	}
	fmt.Println("ids:", req.Ids)
	if result, err := service.Article.Delete(c, &req); err != nil {
		response.Error(c, err)
	} else {
		response.Success(c, result)
	}
}

func (*articleController) View(c *gin.Context) {
	var req define.ArticleViewReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Error(c, err)
		return
	}
	result, err := service.Article.View(c, &req)
	if err != nil {
		response.Error(c, err)
	} else {
		response.Success(c, result)
	}
}

func (*articleController) Toggle(c *gin.Context) {
	var req define.ArticleToggleReq
	if err := c.ShouldBind(&req); err != nil {
		response.Error(c, err)
		return
	}
	result, err := service.Article.Toggle(c, &req)
	fmt.Println(result, err)
	if err != nil {
		response.Error(c, err)
	} else {
		response.Success(c, result)
	}
}
