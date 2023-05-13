package router

import (
	"go-admin/app/admin/apis"

	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerTestaRouter)
}

// 需认证的路由代码
func registerTestaRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.Testa{}
	r := v1.Group("")
	{
		r.GET("/TestaList", api.GetTestaList)
	}
}
