package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type FyStatusGetPageReq struct {
	dto.Pagination `search:"-"`
	Status         string `form:"status"  search:"type:contains;column:status;table:fy_status" comment:"状态"`
	Remark         string `form:"remark"  search:"type:contains;column:remark;table:fy_status" comment:"备注"`
	FyStatusOrder
}

type FyStatusOrder struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:fy_status"`
	Status    string `form:"statusOrder"  search:"type:order;column:status;table:fy_status"`
	Remark    string `form:"remarkOrder"  search:"type:order;column:remark;table:fy_status"`
	CreateBy  string `form:"createByOrder"  search:"type:order;column:create_by;table:fy_status"`
	UpdateBy  string `form:"updateByOrder"  search:"type:order;column:update_by;table:fy_status"`
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:fy_status"`
	UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:fy_status"`
}

func (m *FyStatusGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type FyStatusInsertReq struct {
	Id     int    `json:"-" comment:"编号"` // 编号
	Status string `json:"status" comment:"状态"`
	Remark string `json:"remark" comment:"备注"`
	common.ControlBy
}

func (s *FyStatusInsertReq) Generate(model *models.FyStatus) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Status = s.Status
	model.Remark = s.Remark
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *FyStatusInsertReq) GetId() interface{} {
	return s.Id
}

type FyStatusUpdateReq struct {
	Id     int    `uri:"id" comment:"编号"` // 编号
	Status string `json:"status" comment:"状态"`
	Remark string `json:"remark" comment:"备注"`
	common.ControlBy
}

func (s *FyStatusUpdateReq) Generate(model *models.FyStatus) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Status = s.Status
	model.Remark = s.Remark
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *FyStatusUpdateReq) GetId() interface{} {
	return s.Id
}

// FyStatusGetReq 功能获取请求参数
type FyStatusGetReq struct {
	Id int `uri:"id"`
}

func (s *FyStatusGetReq) GetId() interface{} {
	return s.Id
}

// FyStatusDeleteReq 功能删除请求参数
type FyStatusDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *FyStatusDeleteReq) GetId() interface{} {
	return s.Ids
}
