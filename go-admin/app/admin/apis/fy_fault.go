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

type FyFault struct {
	api.Api
}

// GetPage 获取故障台账列表
// @Summary 获取故障台账列表
// @Description 获取故障台账列表
// @Tags 故障台账
// @Param equipname query string false "设备名"
// @Param equipcode query string false "设备编号"
// @Param area query string false "地区"
// @Param username query string false "负责人"
// @Param type query string false "型号"
// @Param status query string false "状态"
// @Param handlestatus query string false "处理结果"
// @Param reason query string false "故障原因"
// @Param faulttime query time.Time false "故障时间"
// @Param faultcode query string false "故障编号"
// @Param remark query string false "备注"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.FyFault}} "{"code": 200, "data": [...]}"
// @Router /api/v1/fy-fault [get]
// @Security Bearer
func (e FyFault) GetPage(c *gin.Context) {
    req := dto.FyFaultGetPageReq{}
    s := service.FyFault{}
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
	list := make([]models.FyFault, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取故障台账失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取故障台账
// @Summary 获取故障台账
// @Description 获取故障台账
// @Tags 故障台账
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.FyFault} "{"code": 200, "data": [...]}"
// @Router /api/v1/fy-fault/{id} [get]
// @Security Bearer
func (e FyFault) Get(c *gin.Context) {
	req := dto.FyFaultGetReq{}
	s := service.FyFault{}
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
	var object models.FyFault

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取故障台账失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建故障台账
// @Summary 创建故障台账
// @Description 创建故障台账
// @Tags 故障台账
// @Accept application/json
// @Product application/json
// @Param data body dto.FyFaultInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/fy-fault [post]
// @Security Bearer
func (e FyFault) Insert(c *gin.Context) {
    req := dto.FyFaultInsertReq{}
    s := service.FyFault{}
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
		e.Error(500, err, fmt.Sprintf("创建故障台账失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改故障台账
// @Summary 修改故障台账
// @Description 修改故障台账
// @Tags 故障台账
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.FyFaultUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/fy-fault/{id} [put]
// @Security Bearer
func (e FyFault) Update(c *gin.Context) {
    req := dto.FyFaultUpdateReq{}
    s := service.FyFault{}
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
		e.Error(500, err, fmt.Sprintf("修改故障台账失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除故障台账
// @Summary 删除故障台账
// @Description 删除故障台账
// @Tags 故障台账
// @Param data body dto.FyFaultDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/fy-fault [delete]
// @Security Bearer
func (e FyFault) Delete(c *gin.Context) {
    s := service.FyFault{}
    req := dto.FyFaultDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除故障台账失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
