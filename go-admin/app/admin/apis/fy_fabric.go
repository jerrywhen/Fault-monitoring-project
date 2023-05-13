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

type FyFabric struct {
	api.Api
}

// GetPage 获取光纤列表
// @Summary 获取光纤列表
// @Description 获取光纤列表
// @Tags 光纤
// @Param fabric query string false "光纤编号"
// @Param raster query int64 false "光栅数量"
// @Param instruction query string false "说明"
// @Param remark query string false "备注"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.FyFabric}} "{"code": 200, "data": [...]}"
// @Router /api/v1/fy-fabric [get]
// @Security Bearer
func (e FyFabric) GetPage(c *gin.Context) {
    req := dto.FyFabricGetPageReq{}
    s := service.FyFabric{}
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
	list := make([]models.FyFabric, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取光纤失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取光纤
// @Summary 获取光纤
// @Description 获取光纤
// @Tags 光纤
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.FyFabric} "{"code": 200, "data": [...]}"
// @Router /api/v1/fy-fabric/{id} [get]
// @Security Bearer
func (e FyFabric) Get(c *gin.Context) {
	req := dto.FyFabricGetReq{}
	s := service.FyFabric{}
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
	var object models.FyFabric

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取光纤失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建光纤
// @Summary 创建光纤
// @Description 创建光纤
// @Tags 光纤
// @Accept application/json
// @Product application/json
// @Param data body dto.FyFabricInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/fy-fabric [post]
// @Security Bearer
func (e FyFabric) Insert(c *gin.Context) {
    req := dto.FyFabricInsertReq{}
    s := service.FyFabric{}
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
		e.Error(500, err, fmt.Sprintf("创建光纤失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改光纤
// @Summary 修改光纤
// @Description 修改光纤
// @Tags 光纤
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.FyFabricUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/fy-fabric/{id} [put]
// @Security Bearer
func (e FyFabric) Update(c *gin.Context) {
    req := dto.FyFabricUpdateReq{}
    s := service.FyFabric{}
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
		e.Error(500, err, fmt.Sprintf("修改光纤失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除光纤
// @Summary 删除光纤
// @Description 删除光纤
// @Tags 光纤
// @Param data body dto.FyFabricDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/fy-fabric [delete]
// @Security Bearer
func (e FyFabric) Delete(c *gin.Context) {
    s := service.FyFabric{}
    req := dto.FyFabricDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除光纤失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
