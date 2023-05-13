package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type ArticleAGetPageReq struct {
	dto.Pagination `search:"-"`
	ArticleAOrder
}

type ArticleAOrder struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:article_a"`
	Username  string `form:"usernameOrder"  search:"type:order;column:username;table:article_a"`
	Area      string `form:"areaOrder"  search:"type:order;column:area;table:article_a"`
	Type      string `form:"typeOrder"  search:"type:order;column:type;table:article_a"`
	Status    string `form:"statusOrder"  search:"type:order;column:status;table:article_a"`
	Remark    string `form:"remarkOrder"  search:"type:order;column:remark;table:article_a"`
	CreateBy  string `form:"createByOrder"  search:"type:order;column:create_by;table:article_a"`
	UpdateBy  string `form:"updateByOrder"  search:"type:order;column:update_by;table:article_a"`
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:article_a"`
	UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:article_a"`
	DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:article_a"`
}

func (m *ArticleAGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type ArticleAInsertReq struct {
	Id       int    `json:"-" comment:"编号"` // 编号
	Username string `json:"username" comment:"用户名"`
	Area     string `json:"area" comment:"地区"`
	Type     string `json:"type" comment:"型号"`
	Status   string `json:"status" comment:"状态"`
	Remark   string `json:"remark" comment:"备注"`
	common.ControlBy
}

func (s *ArticleAInsertReq) Generate(model *models.ArticleA) {
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

func (s *ArticleAInsertReq) GetId() interface{} {
	return s.Id
}

type ArticleAUpdateReq struct {
	Id       int    `uri:"id" comment:"编号"` // 编号
	Username string `json:"username" comment:"用户名"`
	Area     string `json:"area" comment:"地区"`
	Type     string `json:"type" comment:"型号"`
	Status   string `json:"status" comment:"状态"`
	Remark   string `json:"remark" comment:"备注"`
	common.ControlBy
}

func (s *ArticleAUpdateReq) Generate(model *models.ArticleA) {
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

func (s *ArticleAUpdateReq) GetId() interface{} {
	return s.Id
}

// ArticleAGetReq 功能获取请求参数
type ArticleAGetReq struct {
	Id int `uri:"id"`
}

func (s *ArticleAGetReq) GetId() interface{} {
	return s.Id
}

// ArticleADeleteReq 功能删除请求参数
type ArticleADeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *ArticleADeleteReq) GetId() interface{} {
	return s.Ids
}
