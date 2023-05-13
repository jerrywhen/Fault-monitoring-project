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

type Article struct {
	api.Api
}

// GetPage 获取Article列表
// @Summary 获取Article列表
// @Description 获取Article列表
// @Tags Article
// @Param username query string false "用户名"
// @Param area query string false "地区"
// @Param type query string false "型号"
// @Param status query string false "状态"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Article}} "{"code": 200, "data": [...]}"
// @Router /api/v1/article [get]
// @Security Bearer
func (e Article) GetPage(c *gin.Context) {
    req := dto.ArticleGetPageReq{}
    s := service.Article{}
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
	list := make([]models.Article, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Article失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取Article
// @Summary 获取Article
// @Description 获取Article
// @Tags Article
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.Article} "{"code": 200, "data": [...]}"
// @Router /api/v1/article/{id} [get]
// @Security Bearer
func (e Article) Get(c *gin.Context) {
	req := dto.ArticleGetReq{}
	s := service.Article{}
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
	var object models.Article

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Article失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建Article
// @Summary 创建Article
// @Description 创建Article
// @Tags Article
// @Accept application/json
// @Product application/json
// @Param data body dto.ArticleInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/article [post]
// @Security Bearer
func (e Article) Insert(c *gin.Context) {
    req := dto.ArticleInsertReq{}
    s := service.Article{}
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
		e.Error(500, err, fmt.Sprintf("创建Article失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改Article
// @Summary 修改Article
// @Description 修改Article
// @Tags Article
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.ArticleUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/article/{id} [put]
// @Security Bearer
func (e Article) Update(c *gin.Context) {
    req := dto.ArticleUpdateReq{}
    s := service.Article{}
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
		e.Error(500, err, fmt.Sprintf("修改Article失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除Article
// @Summary 删除Article
// @Description 删除Article
// @Tags Article
// @Param data body dto.ArticleDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/article [delete]
// @Security Bearer
func (e Article) Delete(c *gin.Context) {
    s := service.Article{}
    req := dto.ArticleDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除Article失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
