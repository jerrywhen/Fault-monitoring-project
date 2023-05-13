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

type FyMonitordata struct {
	api.Api
}

// GetPage 获取监控数据列表
// @Summary 获取监控数据列表
// @Description 获取监控数据列表
// @Tags 监控数据
// @Param equipname query string false "设备名"
// @Param equipcode query string false "设备编号"
// @Param remark query string false "备注"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.FyMonitordata}} "{"code": 200, "data": [...]}"
// @Router /api/v1/fy-monitordata [get]
// @Security Bearer
func (e FyMonitordata) GetPage(c *gin.Context) {
    req := dto.FyMonitordataGetPageReq{}
    s := service.FyMonitordata{}
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
	list := make([]models.FyMonitordata, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取监控数据失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取监控数据
// @Summary 获取监控数据
// @Description 获取监控数据
// @Tags 监控数据
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.FyMonitordata} "{"code": 200, "data": [...]}"
// @Router /api/v1/fy-monitordata/{id} [get]
// @Security Bearer
func (e FyMonitordata) Get(c *gin.Context) {
	req := dto.FyMonitordataGetReq{}
	s := service.FyMonitordata{}
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
	var object models.FyMonitordata

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取监控数据失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建监控数据
// @Summary 创建监控数据
// @Description 创建监控数据
// @Tags 监控数据
// @Accept application/json
// @Product application/json
// @Param data body dto.FyMonitordataInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/fy-monitordata [post]
// @Security Bearer
func (e FyMonitordata) Insert(c *gin.Context) {
    req := dto.FyMonitordataInsertReq{}
    s := service.FyMonitordata{}
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
		e.Error(500, err, fmt.Sprintf("创建监控数据失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改监控数据
// @Summary 修改监控数据
// @Description 修改监控数据
// @Tags 监控数据
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.FyMonitordataUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/fy-monitordata/{id} [put]
// @Security Bearer
func (e FyMonitordata) Update(c *gin.Context) {
    req := dto.FyMonitordataUpdateReq{}
    s := service.FyMonitordata{}
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
		e.Error(500, err, fmt.Sprintf("修改监控数据失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除监控数据
// @Summary 删除监控数据
// @Description 删除监控数据
// @Tags 监控数据
// @Param data body dto.FyMonitordataDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/fy-monitordata [delete]
// @Security Bearer
func (e FyMonitordata) Delete(c *gin.Context) {
    s := service.FyMonitordata{}
    req := dto.FyMonitordataDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除监控数据失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
