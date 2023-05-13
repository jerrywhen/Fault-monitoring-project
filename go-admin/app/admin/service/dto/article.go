package dto

import (

	//  "time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type ArticleGetPageReq struct {
	dto.Pagination `search:"-"`
	Username       string `form:"username"  search:"type:exact;column:username;table:article" comment:"用户名"`
	Area           string `form:"area"  search:"type:exact;column:area;table:article" comment:"地区"`
	Type           string `form:"type"  search:"type:exact;column:type;table:article" comment:"型号"`
	Status         string `form:"status"  search:"type:exact;column:status;table:article" comment:"状态"`
	ArticleOrder
}

type ArticleOrder struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:article"`
	Username  string `form:"usernameOrder"  search:"type:order;column:username;table:article"`
	Area      string `form:"areaOrder"  search:"type:order;column:area;table:article"`
	Type      string `form:"typeOrder"  search:"type:order;column:type;table:article"`
	Status    string `form:"statusOrder"  search:"type:order;column:status;table:article"`
	Remark    string `form:"remarkOrder"  search:"type:order;column:remark;table:article"`
	CreateBy  string `form:"createByOrder"  search:"type:order;column:create_by;table:article"`
	UpdateBy  string `form:"updateByOrder"  search:"type:order;column:update_by;table:article"`
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:article"`
	UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:article"`
}

func (m *ArticleGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type ArticleInsertReq struct {
	Id       int    `json:"-" comment:"编号"` // 编号
	Username string `json:"username" comment:"用户名"`
	Area     string `json:"area" comment:"地区"`
	Type     string `json:"type" comment:"型号"`
	Status   string `json:"status" comment:"状态"`
	Remark   string `json:"remark" comment:"备注"`
	common.ControlBy
}

func (s *ArticleInsertReq) Generate(model *models.Article) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Username = s.Username
	model.Area = s.Area
	model.Type = s.Type
	model.Status = s.Status
	model.Remark = s.Remark
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *ArticleInsertReq) GetId() interface{} {
	return s.Id
}

type ArticleUpdateReq struct {
	Id       int    `uri:"id" comment:"编号"` // 编号
	Username string `json:"username" comment:"用户名"`
	Area     string `json:"area" comment:"地区"`
	Type     string `json:"type" comment:"型号"`
	Status   string `json:"status" comment:"状态"`
	Remark   string `json:"remark" comment:"备注"`
	common.ControlBy
}

func (s *ArticleUpdateReq) Generate(model *models.Article) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Username = s.Username
	model.Area = s.Area
	model.Type = s.Type
	model.Status = s.Status
	model.Remark = s.Remark
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *ArticleUpdateReq) GetId() interface{} {
	return s.Id
}

// ArticleGetReq 功能获取请求参数
type ArticleGetReq struct {
	Id int `uri:"id"`
}

func (s *ArticleGetReq) GetId() interface{} {
	return s.Id
}

// ArticleDeleteReq 功能删除请求参数
type ArticleDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *ArticleDeleteReq) GetId() interface{} {
	return s.Ids
}
