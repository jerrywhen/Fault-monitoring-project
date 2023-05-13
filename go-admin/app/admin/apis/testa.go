package apis

import (
	"go-admin/common/apis"

	"github.com/gin-gonic/gin"
)

type Testa struct {
	apis.Api
}

// GetTestaList 获取文章列表
func (e Testa) GetTestaList(c *gin.Context) {
	err := e.MakeContext(c).
		Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK("hello world ！", "success")
}
