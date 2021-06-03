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
		group.GET("/delete", v1.Article.Delete)
	}
}
