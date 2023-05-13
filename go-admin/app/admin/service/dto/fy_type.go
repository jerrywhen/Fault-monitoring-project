package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type FyTypeGetPageReq struct {
	dto.Pagination `search:"-"`
	Type           string `form:"type"  search:"type:exact;column:type;table:fy_type" comment:"型号"`
	Typetext       string `form:"typetext"  search:"type:contains;column:typetext;table:fy_type" comment:"型号说明"`
	Remark         string `form:"remark"  search:"type:contains;column:remark;table:fy_type" comment:"备注"`
	FyTypeOrder
}

type FyTypeOrder struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:fy_type"`
	Type      string `form:"typeOrder"  search:"type:order;column:type;table:fy_type"`
	Typetext  string `form:"typetextOrder"  search:"type:order;column:typetext;table:fy_type"`
	Remark    string `form:"remarkOrder"  search:"type:order;column:remark;table:fy_type"`
	CreateBy  string `form:"createByOrder"  search:"type:order;column:create_by;table:fy_type"`
	UpdateBy  string `form:"updateByOrder"  search:"type:order;column:update_by;table:fy_type"`
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:fy_type"`
	UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:fy_type"`
}

func (m *FyTypeGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type FyTypeInsertReq struct {
	Id       int    `json:"-" comment:"编号"` // 编号
	Type     string `json:"type" comment:"型号"`
	Typetext string `json:"typetext" comment:"型号说明"`
	Remark   string `json:"remark" comment:"备注"`
	common.ControlBy
}

func (s *FyTypeInsertReq) Generate(model *models.FyType) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Type = s.Type
	model.Typetext = s.Typetext
	model.Remark = s.Remark
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *FyTypeInsertReq) GetId() interface{} {
	return s.Id
}

type FyTypeUpdateReq struct {
	Id       int    `uri:"id" comment:"编号"` // 编号
	Type     string `json:"type" comment:"型号"`
	Typetext string `json:"typetext" comment:"型号说明"`
	Remark   string `json:"remark" comment:"备注"`
	common.ControlBy
}

func (s *FyTypeUpdateReq) Generate(model *models.FyType) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Type = s.Type
	model.Typetext = s.Typetext
	model.Remark = s.Remark
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *FyTypeUpdateReq) GetId() interface{} {
	return s.Id
}

// FyTypeGetReq 功能获取请求参数
type FyTypeGetReq struct {
	Id int `uri:"id"`
}

func (s *FyTypeGetReq) GetId() interface{} {
	return s.Id
}

// FyTypeDeleteReq 功能删除请求参数
type FyTypeDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *FyTypeDeleteReq) GetId() interface{} {
	return s.Ids
}
