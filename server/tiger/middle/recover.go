package middle

import (
	"errors"
	"github.com/gin-gonic/gin"
	"tiger-go/tiger/response"
)

// 处理服务器异常(500)
func Recover(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			response.Error(c, errors.New("服务器开小差了"))
		}
	}()
	c.Next()
}
