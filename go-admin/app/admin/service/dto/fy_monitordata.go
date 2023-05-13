package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type FyMonitordataGetPageReq struct {
	dto.Pagination `search:"-"`
	Equipname      string `form:"equipname"  search:"type:contains;column:equipname;table:fy_monitordata" comment:"设备名"`
	Equipcode      string `form:"equipcode"  search:"type:contains;column:equipcode;table:fy_monitordata" comment:"设备编号"`
	Remark         string `form:"remark"  search:"type:contains;column:remark;table:fy_monitordata" comment:"备注"`
	FyMonitordataOrder
}

type FyMonitordataOrder struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:fy_monitordata"`
	Equipname string `form:"equipnameOrder"  search:"type:order;column:equipname;table:fy_monitordata"`
	Equipcode string `form:"equipcodeOrder"  search:"type:order;column:equipcode;table:fy_monitordata"`
	Record1   string `form:"record1Order"  search:"type:order;column:record1;table:fy_monitordata"`
	Record2   string `form:"record2Order"  search:"type:order;column:record2;table:fy_monitordata"`
	Record3   string `form:"record3Order"  search:"type:order;column:record3;table:fy_monitordata"`
	Record4   string `form:"record4Order"  search:"type:order;column:record4;table:fy_monitordata"`
	Record5   string `form:"record5Order"  search:"type:order;column:record5;table:fy_monitordata"`
	Record6   string `form:"record6Order"  search:"type:order;column:record6;table:fy_monitordata"`
	Remark    string `form:"remarkOrder"  search:"type:order;column:remark;table:fy_monitordata"`
	CreateBy  string `form:"createByOrder"  search:"type:order;column:create_by;table:fy_monitordata"`
	UpdateBy  string `form:"updateByOrder"  search:"type:order;column:update_by;table:fy_monitordata"`
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:fy_monitordata"`
	UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:fy_monitordata"`
}

func (m *FyMonitordataGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type FyMonitordataInsertReq struct {
	Id        int    `json:"-" comment:"编号"` // 编号
	Equipname string `json:"equipname" comment:"设备名"`
	Equipcode string `json:"equipcode" comment:"设备编号"`
	Record1   string `json:"record1" comment:"数据1"`
	Record2   string `json:"record2" comment:"数据2"`
	Record3   string `json:"record3" comment:"数据3"`
	Record4   string `json:"record4" comment:"数据4"`
	Record5   string `json:"record5" comment:"数据5"`
	Record6   string `json:"record6" comment:"数据6"`
	Remark    string `json:"remark" comment:"备注"`
	common.ControlBy
}

func (s *FyMonitordataInsertReq) Generate(model *models.FyMonitordata) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Equipname = s.Equipname
	model.Equipcode = s.Equipcode
	model.Record1 = s.Record1
	model.Record2 = s.Record2
	model.Record3 = s.Record3
	model.Record4 = s.Record4
	model.Record5 = s.Record5
	model.Record6 = s.Record6
	model.Remark = s.Remark
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *FyMonitordataInsertReq) GetId() interface{} {
	return s.Id
}

type FyMonitordataUpdateReq struct {
	Id        int    `uri:"id" comment:"编号"` // 编号
	Equipname string `json:"equipname" comment:"设备名"`
	Equipcode string `json:"equipcode" comment:"设备编号"`
	Record1   string `json:"record1" comment:"数据1"`
	Record2   string `json:"record2" comment:"数据2"`
	Record3   string `json:"record3" comment:"数据3"`
	Record4   string `json:"record4" comment:"数据4"`
	Record5   string `json:"record5" comment:"数据5"`
	Record6   string `json:"record6" comment:"数据6"`
	Remark    string `json:"remark" comment:"备注"`
	common.ControlBy
}

func (s *FyMonitordataUpdateReq) Generate(model *models.FyMonitordata) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Equipname = s.Equipname
	model.Equipcode = s.Equipcode
	model.Record1 = s.Record1
	model.Record2 = s.Record2
	model.Record3 = s.Record3
	model.Record4 = s.Record4
	model.Record5 = s.Record5
	model.Record6 = s.Record6
	model.Remark = s.Remark
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *FyMonitordataUpdateReq) GetId() interface{} {
	return s.Id
}

// FyMonitordataGetReq 功能获取请求参数
type FyMonitordataGetReq struct {
	Id int `uri:"id"`
}

func (s *FyMonitordataGetReq) GetId() interface{} {
	return s.Id
}

// FyMonitordataDeleteReq 功能删除请求参数
type FyMonitordataDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *FyMonitordataDeleteReq) GetId() interface{} {
	return s.Ids
}
