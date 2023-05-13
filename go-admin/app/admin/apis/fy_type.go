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

type FyType struct {
	api.Api
}

// GetPage 获取型号列表
// @Summary 获取型号列表
// @Description 获取型号列表
// @Tags 型号
// @Param type query string false "型号"
// @Param typetext query string false "型号说明"
// @Param remark query string false "备注"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.FyType}} "{"code": 200, "data": [...]}"
// @Router /api/v1/fy-type [get]
// @Security Bearer
func (e FyType) GetPage(c *gin.Context) {
    req := dto.FyTypeGetPageReq{}
    s := service.FyType{}
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
	list := make([]models.FyType, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取型号失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取型号
// @Summary 获取型号
// @Description 获取型号
// @Tags 型号
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.FyType} "{"code": 200, "data": [...]}"
// @Router /api/v1/fy-type/{id} [get]
// @Security Bearer
func (e FyType) Get(c *gin.Context) {
	req := dto.FyTypeGetReq{}
	s := service.FyType{}
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
	var object models.FyType

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取型号失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建型号
// @Summary 创建型号
// @Description 创建型号
// @Tags 型号
// @Accept application/json
// @Product application/json
// @Param data body dto.FyTypeInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/fy-type [post]
// @Security Bearer
func (e FyType) Insert(c *gin.Context) {
    req := dto.FyTypeInsertReq{}
    s := service.FyType{}
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
		e.Error(500, err, fmt.Sprintf("创建型号失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改型号
// @Summary 修改型号
// @Description 修改型号
// @Tags 型号
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.FyTypeUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/fy-type/{id} [put]
// @Security Bearer
func (e FyType) Update(c *gin.Context) {
    req := dto.FyTypeUpdateReq{}
    s := service.FyType{}
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
		e.Error(500, err, fmt.Sprintf("修改型号失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除型号
// @Summary 删除型号
// @Description 删除型号
// @Tags 型号
// @Param data body dto.FyTypeDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/fy-type [delete]
// @Security Bearer
func (e FyType) Delete(c *gin.Context) {
    s := service.FyType{}
    req := dto.FyTypeDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除型号失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
