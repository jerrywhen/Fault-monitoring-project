package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type FyAreaGetPageReq struct {
	dto.Pagination `search:"-"`
	Area           string `form:"area"  search:"type:contains;column:area;table:fy_area" comment:"地区"`
	Username       string `form:"username"  search:"type:contains;column:username;table:fy_area" comment:"负责人"`
	Equipcount     string `form:"equipcount"  search:"type:exact;column:equipcount;table:fy_area" comment:"设备数量"`
	Remark         string `form:"remark"  search:"type:contains;column:remark;table:fy_area" comment:"备注"`
	FyAreaOrder
}

type FyAreaOrder struct {
	Id         string `form:"idOrder"  search:"type:order;column:id;table:fy_area"`
	Area       string `form:"areaOrder"  search:"type:order;column:area;table:fy_area"`
	Username   string `form:"usernameOrder"  search:"type:order;column:username;table:fy_area"`
	Equipcount string `form:"equipcountOrder"  search:"type:order;column:equipcount;table:fy_area"`
	Remark     string `form:"remarkOrder"  search:"type:order;column:remark;table:fy_area"`
	CreateBy   string `form:"createByOrder"  search:"type:order;column:create_by;table:fy_area"`
	UpdateBy   string `form:"updateByOrder"  search:"type:order;column:update_by;table:fy_area"`
	CreatedAt  string `form:"createdAtOrder"  search:"type:order;column:created_at;table:fy_area"`
	UpdatedAt  string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:fy_area"`
}

func (m *FyAreaGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type FyAreaInsertReq struct {
	Id         int    `json:"-" comment:"编号"` // 编号
	Area       string `json:"area" comment:"地区"`
	Username   string `json:"username" comment:"负责人"`
	Equipcount string `json:"equipcount" comment:"设备数量"`
	Remark     string `json:"remark" comment:"备注"`
	common.ControlBy
}

func (s *FyAreaInsertReq) Generate(model *models.FyArea) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Area = s.Area
	model.Username = s.Username
	model.Equipcount = s.Equipcount
	model.Remark = s.Remark
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *FyAreaInsertReq) GetId() interface{} {
	return s.Id
}

type FyAreaUpdateReq struct {
	Id         int    `uri:"id" comment:"编号"` // 编号
	Area       string `json:"area" comment:"地区"`
	Username   string `json:"username" comment:"负责人"`
	Equipcount string `json:"equipcount" comment:"设备数量"`
	Remark     string `json:"remark" comment:"备注"`
	common.ControlBy
}

func (s *FyAreaUpdateReq) Generate(model *models.FyArea) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Area = s.Area
	model.Username = s.Username
	model.Equipcount = s.Equipcount
	model.Remark = s.Remark
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *FyAreaUpdateReq) GetId() interface{} {
	return s.Id
}

// FyAreaGetReq 功能获取请求参数
type FyAreaGetReq struct {
	Id int `uri:"id"`
}

func (s *FyAreaGetReq) GetId() interface{} {
	return s.Id
}

// FyAreaDeleteReq 功能删除请求参数
type FyAreaDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *FyAreaDeleteReq) GetId() interface{} {
	return s.Ids
}

//本文中的代码是 Go 语言的结构体和方法，用于处理请求参数和数据模型之间的转换。
// FyAreaGetPageReq 结构体用于获取分页数据的请求参数，包含了分页信息、查询条件等。GetNeedSearch 方法返回结构体本身，以便在处理请求时获取具体的请求参数。
// FyAreaInsertReq 结构体用于插入数据的请求参数，包含了地区、负责人、设备数量、备注等信息。Generate 方法将请求参数转换为数据模型，用于存储到数据库中。GetId 方法返回结构体中的编号。
// FyAreaUpdateReq 结构体用于更新数据的请求参数，包含了编号、地区、负责人、设备数量、备注等信息。Generate 方法将请求参数转换为数据模型，用于更新数据库中的数据。GetId 方法返回结构体中的编号。
// FyAreaGetReq 结构体用于获取单条数据的请求参数，仅包含编号信息。GetId 方法返回结构体中的编号。
// FyAreaDeleteReq 结构体用于删除数据的请求参数，包含了多个编号。GetId 方法返回结构体中的编号列表。
