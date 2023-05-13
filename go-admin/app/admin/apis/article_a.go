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

type ArticleA struct {
	api.Api
}

// GetPage 获取ArticleA列表
// @Summary 获取ArticleA列表
// @Description 获取ArticleA列表
// @Tags ArticleA
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.ArticleA}} "{"code": 200, "data": [...]}"
// @Router /api/v1/article-a [get]
// @Security Bearer
func (e ArticleA) GetPage(c *gin.Context) {
    req := dto.ArticleAGetPageReq{}
    s := service.ArticleA{}
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
	list := make([]models.ArticleA, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取ArticleA失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取ArticleA
// @Summary 获取ArticleA
// @Description 获取ArticleA
// @Tags ArticleA
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.ArticleA} "{"code": 200, "data": [...]}"
// @Router /api/v1/article-a/{id} [get]
// @Security Bearer
func (e ArticleA) Get(c *gin.Context) {
	req := dto.ArticleAGetReq{}
	s := service.ArticleA{}
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
	var object models.ArticleA

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取ArticleA失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建ArticleA
// @Summary 创建ArticleA
// @Description 创建ArticleA
// @Tags ArticleA
// @Accept application/json
// @Product application/json
// @Param data body dto.ArticleAInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/article-a [post]
// @Security Bearer
func (e ArticleA) Insert(c *gin.Context) {
    req := dto.ArticleAInsertReq{}
    s := service.ArticleA{}
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
		e.Error(500, err, fmt.Sprintf("创建ArticleA失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改ArticleA
// @Summary 修改ArticleA
// @Description 修改ArticleA
// @Tags ArticleA
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.ArticleAUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/article-a/{id} [put]
// @Security Bearer
func (e ArticleA) Update(c *gin.Context) {
    req := dto.ArticleAUpdateReq{}
    s := service.ArticleA{}
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
		e.Error(500, err, fmt.Sprintf("修改ArticleA失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除ArticleA
// @Summary 删除ArticleA
// @Description 删除ArticleA
// @Tags ArticleA
// @Param data body dto.ArticleADeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/article-a [delete]
// @Security Bearer
func (e ArticleA) Delete(c *gin.Context) {
    s := service.ArticleA{}
    req := dto.ArticleADeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除ArticleA失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
