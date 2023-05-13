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

type FyStatus struct {
	api.Api
}

// GetPage 获取状态列表
// @Summary 获取状态列表
// @Description 获取状态列表
// @Tags 状态
// @Param status query string false "状态"
// @Param remark query string false "备注"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.FyStatus}} "{"code": 200, "data": [...]}"
// @Router /api/v1/fy-status [get]
// @Security Bearer
func (e FyStatus) GetPage(c *gin.Context) {
    req := dto.FyStatusGetPageReq{}
    s := service.FyStatus{}
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
	list := make([]models.FyStatus, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取状态失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取状态
// @Summary 获取状态
// @Description 获取状态
// @Tags 状态
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.FyStatus} "{"code": 200, "data": [...]}"
// @Router /api/v1/fy-status/{id} [get]
// @Security Bearer
func (e FyStatus) Get(c *gin.Context) {
	req := dto.FyStatusGetReq{}
	s := service.FyStatus{}
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
	var object models.FyStatus

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取状态失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建状态
// @Summary 创建状态
// @Description 创建状态
// @Tags 状态
// @Accept application/json
// @Product application/json
// @Param data body dto.FyStatusInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/fy-status [post]
// @Security Bearer
func (e FyStatus) Insert(c *gin.Context) {
    req := dto.FyStatusInsertReq{}
    s := service.FyStatus{}
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
		e.Error(500, err, fmt.Sprintf("创建状态失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改状态
// @Summary 修改状态
// @Description 修改状态
// @Tags 状态
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.FyStatusUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/fy-status/{id} [put]
// @Security Bearer
func (e FyStatus) Update(c *gin.Context) {
    req := dto.FyStatusUpdateReq{}
    s := service.FyStatus{}
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
		e.Error(500, err, fmt.Sprintf("修改状态失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除状态
// @Summary 删除状态
// @Description 删除状态
// @Tags 状态
// @Param data body dto.FyStatusDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/fy-status [delete]
// @Security Bearer
func (e FyStatus) Delete(c *gin.Context) {
    s := service.FyStatus{}
    req := dto.FyStatusDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除状态失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
