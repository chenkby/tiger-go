package router

import (
	"github.com/gin-gonic/gin"
	v1 "tiger-go/applications/admin-api/internal/v1"
)

func Load(r *gin.Engine) {
	//r.Use(middle.Recover)
	routerGroup := r.Group("/v1")

	group := routerGroup.Group("/article")
	{
		group.GET("/list", v1.Article.List)
		group.POST("/create", v1.Article.Create)
		group.POST("/update", v1.Article.Update)
		group.GET("/view", v1.Article.View)
		group.POST("/toggle", v1.Article.Toggle)
		group.POST("/delete", v1.Article.Delete)
	}

	group = routerGroup.Group("/video")
	{
		group.GET("/list", v1.Video.List)
		group.POST("/create", v1.Video.Create)
		group.POST("/update", v1.Video.Update)
		group.GET("/info", v1.Video.Info)
		group.GET("/view", v1.Video.View)
		group.POST("/toggle", v1.Video.Toggle)
		group.POST("/delete", v1.Video.Delete)
	}

	group = routerGroup.Group("/generate")
	{
		group.GET("/list", v1.Article.List)
		group.POST("/create", v1.Generate.Create)
		group.POST("/update", v1.Article.Update)
		group.GET("/view", v1.Article.View)
		group.POST("/toggle", v1.Article.Toggle)
		group.POST("/delete", v1.Article.Delete)
	}
}
