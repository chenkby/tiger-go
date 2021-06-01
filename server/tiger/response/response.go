package response

import (
	"github.com/gin-gonic/gin"
	"mang/tiger/data"
	"net/http"
)

func Success(c *gin.Context, data interface{}, attach ...map[string]interface{}) {
	response := map[string]interface{}{
		"err_code": 0,
		"data":     data,
	}
	if attach != nil {
		for key, val := range attach[0] {
			response[key] = val
		}
	}
	c.JSON(200, response)
}

// 分页响应
func PageSuccess(c *gin.Context, data interface{}, paging *data.Paging) {
	Success(c, data, map[string]interface{}{
		"pagination": paging,
	})
}

func Error(c *gin.Context, err error, code ...int) {
	if code == nil {
		code = []int{500}
	}
	c.AbortWithStatusJSON(http.StatusOK, map[string]interface{}{
		"err_code": code[0],
		"err_msg":  err.Error(),
	})

}
