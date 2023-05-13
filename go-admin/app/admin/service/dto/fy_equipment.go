package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type FyEquipmentGetPageReq struct {
	dto.Pagination `search:"-"`
	Equipname      string `form:"equipname"  search:"type:contains;column:equipname;table:fy_equipment" comment:"设备名"`
	Equipcode      string `form:"equipcode"  search:"type:contains;column:equipcode;table:fy_equipment" comment:"设备编码"`
	Username       string `form:"username"  search:"type:contains;column:username;table:fy_equipment" comment:"负责人"`
	Area           string `form:"area"  search:"type:contains;column:area;table:fy_equipment" comment:"地区"`
	Type           string `form:"type"  search:"type:contains;column:type;table:fy_equipment" comment:"型号"`
	Status         string `form:"status"  search:"type:contains;column:status;table:fy_equipment" comment:"状态"`
	Fabric         string `form:"fabric"  search:"type:contains;column:fabric;table:fy_equipment" comment:"状态"`
	Raster         string `form:"raster"  search:"type:contains;column:raster;table:fy_equipment" comment:"状态"`
	Faultcount     string `form:"faultcount"  search:"type:gte;column:faultcount;table:fy_equipment" comment:"故障次数"`
	Remark         string `form:"remark"  search:"type:contains;column:remark;table:fy_equipment" comment:"备注"`
	FyEquipmentOrder
}

type FyEquipmentOrder struct {
	Id         string `form:"idOrder"  search:"type:order;column:id;table:fy_equipment"`
	Equipname  string `form:"equipnameOrder"  search:"type:order;column:equipname;table:fy_equipment"`
	Equipcode  string `form:"equipcodeOrder"  search:"type:order;column:equipcode;table:fy_equipment"`
	Username   string `form:"usernameOrder"  search:"type:order;column:username;table:fy_equipment"`
	Area       string `form:"areaOrder"  search:"type:order;column:area;table:fy_equipment"`
	Type       string `form:"typeOrder"  search:"type:order;column:type;table:fy_equipment"`
	Status     string `form:"statusOrder"  search:"type:order;column:status;table:fy_equipment"`
	Fabric     string `form:"fabricOrder"  search:"type:order;column:fabric;table:fy_equipment"`
	Raster     string `form:"rasterOrder"  search:"type:order;column:raster;table:fy_equipment"`
	Faultcount string `form:"faultcountOrder"  search:"type:order;column:faultcount;table:fy_equipment"`
	Remark     string `form:"remarkOrder"  search:"type:order;column:remark;table:fy_equipment"`
	CreateBy   string `form:"createByOrder"  search:"type:order;column:create_by;table:fy_equipment"`
	UpdateBy   string `form:"updateByOrder"  search:"type:order;column:update_by;table:fy_equipment"`
	CreatedAt  string `form:"createdAtOrder"  search:"type:order;column:created_at;table:fy_equipment"`
	UpdatedAt  string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:fy_equipment"`
}

func (m *FyEquipmentGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type FyEquipmentInsertReq struct {
	Id         int    `json:"-" comment:"编号"` // 编号
	Equipname  string `json:"equipname" comment:"设备名"`
	Equipcode  string `json:"equipcode" comment:"设备编码"`
	Username   string `json:"username" comment:"负责人"`
	Area       string `json:"area" comment:"地区"`
	Type       string `json:"type" comment:"型号"`
	Status     string `json:"status" comment:"状态"`
	Fabric     string `json:"fabric" comment:"光纤编号"`
	Raster     string `json:"raster" comment:"光栅编号"`
	Faultcount string `json:"faultcount" comment:"故障次数"`
	Remark     string `json:"remark" comment:"备注"`
	common.ControlBy
}

func (s *FyEquipmentInsertReq) Generate(model *models.FyEquipment) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Equipname = s.Equipname
	model.Equipcode = s.Equipcode
	model.Username = s.Username
	model.Area = s.Area
	model.Type = s.Type
	model.Status = s.Status
	model.Fabric = s.Fabric
	model.Raster = s.Raster
	model.Faultcount = s.Faultcount
	model.Remark = s.Remark
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *FyEquipmentInsertReq) GetId() interface{} {
	return s.Id
}

type FyEquipmentUpdateReq struct {
	Id         int    `uri:"id" comment:"编号"` // 编号
	Equipname  string `json:"equipname" comment:"设备名"`
	Equipcode  string `json:"equipcode" comment:"设备编码"`
	Username   string `json:"username" comment:"负责人"`
	Area       string `json:"area" comment:"地区"`
	Type       string `json:"type" comment:"型号"`
	Status     string `json:"status" comment:"状态"`
	Fabric     string `json:"fabric" comment:"光纤编号"`
	Raster     string `json:"raster" comment:"光栅编号"`
	Faultcount string `json:"faultcount" comment:"故障次数"`
	Remark     string `json:"remark" comment:"备注"`
	common.ControlBy
}

func (s *FyEquipmentUpdateReq) Generate(model *models.FyEquipment) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Equipname = s.Equipname
	model.Equipcode = s.Equipcode
	model.Username = s.Username
	model.Area = s.Area
	model.Type = s.Type
	model.Status = s.Status
	model.Fabric = s.Fabric
	model.Raster = s.Raster
	model.Faultcount = s.Faultcount
	model.Remark = s.Remark
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *FyEquipmentUpdateReq) GetId() interface{} {
	return s.Id
}

// FyEquipmentGetReq 功能获取请求参数
type FyEquipmentGetReq struct {
	Id int `uri:"id"`
}

func (s *FyEquipmentGetReq) GetId() interface{} {
	return s.Id
}

// FyEquipmentDeleteReq 功能删除请求参数
type FyEquipmentDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *FyEquipmentDeleteReq) GetId() interface{} {
	return s.Ids
}
