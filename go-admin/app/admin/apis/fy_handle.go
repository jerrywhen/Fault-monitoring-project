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

type FyHandle struct {
	api.Api
}

// GetPage 获取处理台账列表
// @Summary 获取处理台账列表
// @Description 获取处理台账列表
// @Tags 处理台账
// @Param equipname query string false "设备名"
// @Param equipcode query string false "设备编号"
// @Param faultcode query string false "故障编号"
// @Param type query string false "型号"
// @Param handlemethod query string false "处理方法"
// @Param handlestatus query string false "处理结果"
// @Param handletime query time.Time false "处理时间"
// @Param remark query string false "备注"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.FyHandle}} "{"code": 200, "data": [...]}"
// @Router /api/v1/fy-handle [get]
// @Security Bearer
func (e FyHandle) GetPage(c *gin.Context) {
    req := dto.FyHandleGetPageReq{}
    s := service.FyHandle{}
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
	list := make([]models.FyHandle, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取处理台账失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取处理台账
// @Summary 获取处理台账
// @Description 获取处理台账
// @Tags 处理台账
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.FyHandle} "{"code": 200, "data": [...]}"
// @Router /api/v1/fy-handle/{id} [get]
// @Security Bearer
func (e FyHandle) Get(c *gin.Context) {
	req := dto.FyHandleGetReq{}
	s := service.FyHandle{}
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
	var object models.FyHandle

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取处理台账失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建处理台账
// @Summary 创建处理台账
// @Description 创建处理台账
// @Tags 处理台账
// @Accept application/json
// @Product application/json
// @Param data body dto.FyHandleInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/fy-handle [post]
// @Security Bearer
func (e FyHandle) Insert(c *gin.Context) {
    req := dto.FyHandleInsertReq{}
    s := service.FyHandle{}
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
		e.Error(500, err, fmt.Sprintf("创建处理台账失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改处理台账
// @Summary 修改处理台账
// @Description 修改处理台账
// @Tags 处理台账
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.FyHandleUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/fy-handle/{id} [put]
// @Security Bearer
func (e FyHandle) Update(c *gin.Context) {
    req := dto.FyHandleUpdateReq{}
    s := service.FyHandle{}
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
		e.Error(500, err, fmt.Sprintf("修改处理台账失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除处理台账
// @Summary 删除处理台账
// @Description 删除处理台账
// @Tags 处理台账
// @Param data body dto.FyHandleDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/fy-handle [delete]
// @Security Bearer
func (e FyHandle) Delete(c *gin.Context) {
    s := service.FyHandle{}
    req := dto.FyHandleDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除处理台账失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
