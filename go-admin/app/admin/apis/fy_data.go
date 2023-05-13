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

type FyData struct {
	api.Api
}

// GetPage 获取FyData列表
// @Summary 获取FyData列表
// @Description 获取FyData列表
// @Tags FyData
// @Param equipname query string false "设备名"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.FyData}} "{"code": 200, "data": [...]}"
// @Router /api/v1/fy-data [get]
// @Security Bearer
func (e FyData) GetPage(c *gin.Context) {
    req := dto.FyDataGetPageReq{}
    s := service.FyData{}
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
	list := make([]models.FyData, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取FyData失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取FyData
// @Summary 获取FyData
// @Description 获取FyData
// @Tags FyData
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.FyData} "{"code": 200, "data": [...]}"
// @Router /api/v1/fy-data/{id} [get]
// @Security Bearer
func (e FyData) Get(c *gin.Context) {
	req := dto.FyDataGetReq{}
	s := service.FyData{}
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
	var object models.FyData

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取FyData失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建FyData
// @Summary 创建FyData
// @Description 创建FyData
// @Tags FyData
// @Accept application/json
// @Product application/json
// @Param data body dto.FyDataInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/fy-data [post]
// @Security Bearer
func (e FyData) Insert(c *gin.Context) {
    req := dto.FyDataInsertReq{}
    s := service.FyData{}
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
		e.Error(500, err, fmt.Sprintf("创建FyData失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改FyData
// @Summary 修改FyData
// @Description 修改FyData
// @Tags FyData
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.FyDataUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/fy-data/{id} [put]
// @Security Bearer
func (e FyData) Update(c *gin.Context) {
    req := dto.FyDataUpdateReq{}
    s := service.FyData{}
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
		e.Error(500, err, fmt.Sprintf("修改FyData失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除FyData
// @Summary 删除FyData
// @Description 删除FyData
// @Tags FyData
// @Param data body dto.FyDataDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/fy-data [delete]
// @Security Bearer
func (e FyData) Delete(c *gin.Context) {
    s := service.FyData{}
    req := dto.FyDataDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除FyData失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
