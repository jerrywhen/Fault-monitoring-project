package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"

	"go-admin/app/admin/apis"
	"go-admin/common/actions"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerFyAreaRouter)
}

// registerFyAreaRouter
func registerFyAreaRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.FyArea{}
	r := v1.Group("/fy-area").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", actions.PermissionAction(), api.GetPage)
		r.GET("/:id", actions.PermissionAction(), api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", actions.PermissionAction(), api.Update)
		r.DELETE("", api.Delete)
	}
}

// 本文中的代码是 Go 语言的路由配置，用于将 FyArea 相关的 API 和请求路由进行绑定，并进行权限控制。
// registerFyAreaRouter 方法用于注册 FyArea 相关的路由，传入 gin.RouterGroup 和 jwt.GinJWTMiddleware 两个参数。在方法中，首先创建 FyArea 的 API 对象，然后创建一个路由组，将路由组和权限控制中间件绑定。
// 在路由组中，使用 GET、POST、PUT、DELETE 等方法将请求路由和对应的 API 方法进行绑定，同时使用 actions 包中的 PermissionAction 方法进行权限控制。其中，GET 方法对应获取分页数据和获取单条数据，POST 方法对应插入数据，PUT 方法对应更新数据，DELETE 方法对应删除数据。最后，将路由组注册到 gin.RouterGroup 中，完成路由配置。
