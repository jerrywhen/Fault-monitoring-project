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

// 这段代码定义了一个名为FyUser的结构体，它嵌入了api.Api这个结构体。这种方式被称为结构体嵌入或匿名字段，它可以让FyUser结构体继承api.Api结构体的所有属性和方法。这意味着FyUser结构体可以直接使用api.Api结构体的方法和字段，而不需要重新定义一遍。
type FyUser struct {
	api.Api
}

// GetPage 获取负责人列表
// @Summary 获取负责人列表
// @Description 获取负责人列表
// @Tags 负责人
// @Param username query string false "用户名"
// @Param area query string false "负责地区"
// @Param userstatus query string false "人员状态"
// @Param duty query string false "职责"
// @Param remark query string false "备注"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.FyUser}} "{"code": 200, "data": [...]}"
// @Router /api/v1/fy-user [get]
// @Security Bearer
func (e FyUser) GetPage(c *gin.Context) {
	req := dto.FyUserGetPageReq{}
	s := service.FyUser{}
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
	list := make([]models.FyUser, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取负责人失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取负责人
// @Summary 获取负责人
// @Description 获取负责人
// @Tags 负责人
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.FyUser} "{"code": 200, "data": [...]}"
// @Router /api/v1/fy-user/{id} [get]
// @Security Bearer
func (e FyUser) Get(c *gin.Context) {
	req := dto.FyUserGetReq{}
	s := service.FyUser{}
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
	var object models.FyUser

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取负责人失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建负责人
// @Summary 创建负责人
// @Description 创建负责人
// @Tags 负责人
// @Accept application/json
// @Product application/json
// @Param data body dto.FyUserInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/fy-user [post]
// @Security Bearer
func (e FyUser) Insert(c *gin.Context) {
	req := dto.FyUserInsertReq{}
	s := service.FyUser{}
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
		e.Error(500, err, fmt.Sprintf("创建负责人失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改负责人
// @Summary 修改负责人
// @Description 修改负责人
// @Tags 负责人
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.FyUserUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/fy-user/{id} [put]
// @Security Bearer
func (e FyUser) Update(c *gin.Context) {
	req := dto.FyUserUpdateReq{}
	s := service.FyUser{}
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
		e.Error(500, err, fmt.Sprintf("修改负责人失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除负责人
// @Summary 删除负责人
// @Description 删除负责人
// @Tags 负责人
// @Param data body dto.FyUserDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/fy-user [delete]
// @Security Bearer
func (e FyUser) Delete(c *gin.Context) {
	s := service.FyUser{}
	req := dto.FyUserDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除负责人失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}

// 这是一个FyUser结构体的方法集合，包含了GetPage、Get、Insert、Update和Delete等方法，用于实现负责人（FyUser）相关的API接口。
// 在这些方法中，都会通过MakeContext、MakeOrm和MakeService等方法来初始化上下文、ORM连接和服务对象。然后根据不同的接口需求，通过调用service.FyUser结构体中的不同方法来实现查询、新增、修改和删除负责人等操作。如果发生错误，则通过e.Error方法返回错误信息。
// 这个方法集合是整个API的核心部分，通过调用不同的方法来实现各种不同的功能，对于开发者来说，只需要关注API接口的定义和需求，无需过多关注具体实现的细节。
