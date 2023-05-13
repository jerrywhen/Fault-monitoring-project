package dto

import (
	//  "time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type FyDataGetPageReq struct {
	dto.Pagination `search:"-"`
	Equipname      string `form:"equipname"  search:"type:contains;column:equipname;table:fy_data" comment:"设备名"`
	FyDataOrder
}

type FyDataOrder struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:fy_data"`
	Equipname string `form:"equipnameOrder"  search:"type:order;column:equipname;table:fy_data"`
	Datax     string `form:"dataxOrder"  search:"type:order;column:datax;table:fy_data"`
	Datay     string `form:"datayOrder"  search:"type:order;column:datay;table:fy_data"`
	CreateBy  string `form:"createByOrder"  search:"type:order;column:create_by;table:fy_data"`
	UpdateBy  string `form:"updateByOrder"  search:"type:order;column:update_by;table:fy_data"`
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:fy_data"`
	UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:fy_data"`
}

func (m *FyDataGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type FyDataInsertReq struct {
	Id        int    `json:"-" comment:""` //
	Equipname string `json:"equipname" comment:"设备名"`
	Datax     string `json:"datax" comment:"x数据"`
	Datay     string `json:"datay" comment:"y数据"`
	common.ControlBy
}

func (s *FyDataInsertReq) Generate(model *models.FyData) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Equipname = s.Equipname
	model.Datax = s.Datax
	model.Datay = s.Datay
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *FyDataInsertReq) GetId() interface{} {
	return s.Id
}

type FyDataUpdateReq struct {
	Id        int    `uri:"id" comment:""` //
	Equipname string `json:"equipname" comment:"设备名"`
	Datax     string `json:"datax" comment:"x数据"`
	Datay     string `json:"datay" comment:"y数据"`
	common.ControlBy
}

func (s *FyDataUpdateReq) Generate(model *models.FyData) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Equipname = s.Equipname
	model.Datax = s.Datax
	model.Datay = s.Datay
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *FyDataUpdateReq) GetId() interface{} {
	return s.Id
}

// FyDataGetReq 功能获取请求参数
type FyDataGetReq struct {
	Id int `uri:"id"`
}

func (s *FyDataGetReq) GetId() interface{} {
	return s.Id
}

// FyDataDeleteReq 功能删除请求参数
type FyDataDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *FyDataDeleteReq) GetId() interface{} {
	return s.Ids
}
