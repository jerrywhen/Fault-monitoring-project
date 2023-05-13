package apis

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type FyArea struct {
	api.Api
}

// GetPage 获取地区列表
// @Summary 获取地区列表
// @Description 获取地区列表
// @Tags 地区
// @Param area query string false "地区"
// @Param username query string false "负责人"
// @Param equipcount query string false "设备数量"
// @Param remark query string false "备注"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.FyArea}} "{"code": 200, "data": [...]}"
// @Router /api/v1/fy-area [get]
// @Security Bearer
func (e FyArea) GetPage(c *gin.Context) {
	req := dto.FyAreaGetPageReq{}
	s := service.FyArea{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.FyArea, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取地区失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取地区
// @Summary 获取地区
// @Description 获取地区
// @Tags 地区
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.FyArea} "{"code": 200, "data": [...]}"
// @Router /api/v1/fy-area/{id} [get]
// @Security Bearer
func (e FyArea) Get(c *gin.Context) {
	req := dto.FyAreaGetReq{}
	s := service.FyArea{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.FyArea

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取地区失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建地区
// @Summary 创建地区
// @Description 创建地区
// @Tags 地区
// @Accept application/json
// @Product application/json
// @Param data body dto.FyAreaInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/fy-area [post]
// @Security Bearer
func (e FyArea) Insert(c *gin.Context) {
	req := dto.FyAreaInsertReq{}
	s := service.FyArea{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建地区失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改地区
// @Summary 修改地区
// @Description 修改地区
// @Tags 地区
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.FyAreaUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/fy-area/{id} [put]
// @Security Bearer
func (e FyArea) Update(c *gin.Context) {
	req := dto.FyAreaUpdateReq{}
	s := service.FyArea{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改地区失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除地区
// @Summary 删除地区
// @Description 删除地区
// @Tags 地区
// @Param data body dto.FyAreaDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/fy-area [delete]
// @Security Bearer
func (e FyArea) Delete(c *gin.Context) {
	s := service.FyArea{}
	req := dto.FyAreaDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除地区失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}

//以上代码是一个基于 Gin 框架开发的 RESTful API 的实现，主要是对地区信息的增删改查操作。该代码使用了自定义的 dto 和 service 结构体，其中 dto 结构体主要用于数据传输，service 结构体主要用于业务逻辑处理。

// 在 GetPage 方法中，通过输入参数 req 和 s 获取分页查询所需的条件和 Service 实例，然后调用 Service 中的 GetPage 方法进行分页查询，并将查询结果返回给客户端。

// 在 Get 方法中，通过输入参数 req 和 s 获取查询所需的条件和 Service 实例，然后调用 Service 中的 Get 方法进行查询，并将查询结果返回给客户端。

// 在 Insert 方法中，通过输入参数 req 和 s 获取插入所需的数据和 Service 实例，然后调用 Service 中的 Insert 方法进行插入操作，并将插入结果返回给客户端。

// 在 Update 方法中，通过输入参数 req 和 s 获取更新所需的数据和 Service 实例，然后调用 Service 中的 Update 方法进行更新操作，并将更新结果返回给客户端。

// 在 Delete 方法中，通过输入参数 req 和 s 获取删除所需的条件和 Service 实例，然后调用 Service 中的 Remove 方法进行删除操作，并将删除结果返回给客户端。

// 在以上方法中，还使用了 MakeContext、MakeOrm 和 MakeService 方法初始化了 Context、ORM 和 Service，以及调用了 GetPermissionFromContext 方法获取用户权限信息。如果出现错误，将会记录日志并返回相应的错误信息。
