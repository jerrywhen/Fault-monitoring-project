package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"

	"go-admin/app/admin/apis"
	"go-admin/common/actions"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerFyUserRouter)
}

// registerFyUserRouter
func registerFyUserRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.FyUser{}
	r := v1.Group("/fy-user").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", actions.PermissionAction(), api.GetPage)
		r.GET("/:id", actions.PermissionAction(), api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", actions.PermissionAction(), api.Update)
		r.DELETE("", api.Delete)
	}
}

//这是一个使用 Gin 框架实现的路由注册函数。其中，通过 init 函数将 registerFyUserRouter 函数注册进路由树中。registerFyUserRouter 函数使用了 authMiddleware 和 middleware.AuthCheckRole 中间件，用于验证用户是否登录并是否有相应的权限。在路由中，使用 GET 方法获取用户列表和单个用户信息，使用 POST 方法添加用户，使用 PUT 方法更新用户信息，使用 DELETE 方法删除用户。同时，使用 actions.PermissionAction() 方法验证用户是否有相应的权限进行操作。这样，就可以通过访问相应的 API 实现对用户数据的 CRUD 操作。
