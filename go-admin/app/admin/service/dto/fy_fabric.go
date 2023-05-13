package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type FyFabricGetPageReq struct {
	dto.Pagination `search:"-"`
	Fabric         string `form:"fabric"  search:"type:contains;column:fabric;table:fy_fabric" comment:"光纤编号"`
	Raster         int64  `form:"raster"  search:"type:exact;column:raster;table:fy_fabric" comment:"光栅数量"`
	Instruction    string `form:"instruction"  search:"type:contains;column:instruction;table:fy_fabric" comment:"说明"`
	Remark         string `form:"remark"  search:"type:contains;column:remark;table:fy_fabric" comment:"备注"`
	FyFabricOrder
}

type FyFabricOrder struct {
	Id          string `form:"idOrder"  search:"type:order;column:id;table:fy_fabric"`
	Fabric      string `form:"fabricOrder"  search:"type:order;column:fabric;table:fy_fabric"`
	Raster      string `form:"rasterOrder"  search:"type:order;column:raster;table:fy_fabric"`
	Instruction string `form:"instructionOrder"  search:"type:order;column:instruction;table:fy_fabric"`
	Remark      string `form:"remarkOrder"  search:"type:order;column:remark;table:fy_fabric"`
	CreateBy    string `form:"createByOrder"  search:"type:order;column:create_by;table:fy_fabric"`
	UpdateBy    string `form:"updateByOrder"  search:"type:order;column:update_by;table:fy_fabric"`
	CreatedAt   string `form:"createdAtOrder"  search:"type:order;column:created_at;table:fy_fabric"`
	UpdatedAt   string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:fy_fabric"`
}

func (m *FyFabricGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type FyFabricInsertReq struct {
	Id          int    `json:"-" comment:"编号"` // 编号
	Fabric      string `json:"fabric" comment:"光纤编号"`
	Raster      int64  `json:"raster" comment:"光栅数量"`
	Instruction string `json:"instruction" comment:"说明"`
	Remark      string `json:"remark" comment:"备注"`
	common.ControlBy
}

func (s *FyFabricInsertReq) Generate(model *models.FyFabric) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Fabric = s.Fabric
	model.Raster = s.Raster
	model.Instruction = s.Instruction
	model.Remark = s.Remark
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *FyFabricInsertReq) GetId() interface{} {
	return s.Id
}

type FyFabricUpdateReq struct {
	Id          int    `uri:"id" comment:"编号"` // 编号
	Fabric      string `json:"fabric" comment:"光纤编号"`
	Raster      int64  `json:"raster" comment:"光栅数量"`
	Instruction string `json:"instruction" comment:"说明"`
	Remark      string `json:"remark" comment:"备注"`
	common.ControlBy
}

func (s *FyFabricUpdateReq) Generate(model *models.FyFabric) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Fabric = s.Fabric
	model.Raster = s.Raster
	model.Instruction = s.Instruction
	model.Remark = s.Remark
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *FyFabricUpdateReq) GetId() interface{} {
	return s.Id
}

// FyFabricGetReq 功能获取请求参数
type FyFabricGetReq struct {
	Id int `uri:"id"`
}

func (s *FyFabricGetReq) GetId() interface{} {
	return s.Id
}

// FyFabricDeleteReq 功能删除请求参数
type FyFabricDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *FyFabricDeleteReq) GetId() interface{} {
	return s.Ids
}
