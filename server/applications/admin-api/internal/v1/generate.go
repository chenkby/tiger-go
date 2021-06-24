package v1

import (
	"github.com/gin-gonic/gin"
	"tiger-go/applications/admin-api/module/generate_code/define"
	"tiger-go/applications/admin-api/module/generate_code/service"
	"tiger-go/tiger/response"
)

var Generate = new(generateApi)

type generateApi struct {}

func (*generateApi) Create(c *gin.Context) {
	var req define.GenerateCodeCreateReq
	if err := c.ShouldBind(&req); err != nil {
		response.Error(c, err)
		return
	}
	if result, err := service.GenerateCode.Model(req.DbName,req.TableName); err != nil {
		response.Error(c, err)
	} else {
		response.Success(c, result)
	}
}
