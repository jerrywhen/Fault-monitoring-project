package dto

import (
	"time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type FyHandleGetPageReq struct {
	dto.Pagination `search:"-"`
	Equipname      string    `form:"equipname"  search:"type:contains;column:equipname;table:fy_handle" comment:"设备名"`
	Equipcode      string    `form:"equipcode"  search:"type:contains;column:equipcode;table:fy_handle" comment:"设备编号"`
	Faultcode      string    `form:"faultcode"  search:"type:exact;column:faultcode;table:fy_handle" comment:"故障编号"`
	Type           string    `form:"type"  search:"type:contains;column:type;table:fy_handle" comment:"型号"`
	Handlemethod   string    `form:"handlemethod"  search:"type:contains;column:handlemethod;table:fy_handle" comment:"处理方法"`
	Handlestatus   string    `form:"handlestatus"  search:"type:contains;column:handlestatus;table:fy_handle" comment:"处理结果"`
	Handletime     time.Time `form:"handletime"  search:"type:contains;column:handletime;table:fy_handle" comment:"处理时间"`
	Remark         string    `form:"remark"  search:"type:contains;column:remark;table:fy_handle" comment:"备注"`
	FyHandleOrder
}

type FyHandleOrder struct {
	Id           string `form:"idOrder"  search:"type:order;column:id;table:fy_handle"`
	Equipname    string `form:"equipnameOrder"  search:"type:order;column:equipname;table:fy_handle"`
	Equipcode    string `form:"equipcodeOrder"  search:"type:order;column:equipcode;table:fy_handle"`
	Faultcode    string `form:"faultcodeOrder"  search:"type:order;column:faultcode;table:fy_handle"`
	Type         string `form:"typeOrder"  search:"type:order;column:type;table:fy_handle"`
	Handlemethod string `form:"handlemethodOrder"  search:"type:order;column:handlemethod;table:fy_handle"`
	Handlestatus string `form:"handlestatusOrder"  search:"type:order;column:handlestatus;table:fy_handle"`
	Handletime   string `form:"handletimeOrder"  search:"type:order;column:handletime;table:fy_handle"`
	Remark       string `form:"remarkOrder"  search:"type:order;column:remark;table:fy_handle"`
	CreateBy     string `form:"createByOrder"  search:"type:order;column:create_by;table:fy_handle"`
	UpdateBy     string `form:"updateByOrder"  search:"type:order;column:update_by;table:fy_handle"`
	CreatedAt    string `form:"createdAtOrder"  search:"type:order;column:created_at;table:fy_handle"`
	UpdatedAt    string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:fy_handle"`
}

func (m *FyHandleGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type FyHandleInsertReq struct {
	Id           int       `json:"-" comment:"编号"` // 编号
	Equipname    string    `json:"equipname" comment:"设备名"`
	Equipcode    string    `json:"equipcode" comment:"设备编号"`
	Faultcode    string    `json:"faultcode" comment:"故障编号"`
	Type         string    `json:"type" comment:"型号"`
	Handlemethod string    `json:"handlemethod" comment:"处理方法"`
	Handlestatus string    `json:"handlestatus" comment:"处理结果"`
	Handletime   time.Time `json:"handletime" comment:"处理时间"`
	Remark       string    `json:"remark" comment:"备注"`
	common.ControlBy
}

func (s *FyHandleInsertReq) Generate(model *models.FyHandle) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Equipname = s.Equipname
	model.Equipcode = s.Equipcode
	model.Faultcode = s.Faultcode
	model.Type = s.Type
	model.Handlemethod = s.Handlemethod
	model.Handlestatus = s.Handlestatus
	model.Handletime = s.Handletime
	model.Remark = s.Remark
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *FyHandleInsertReq) GetId() interface{} {
	return s.Id
}

type FyHandleUpdateReq struct {
	Id           int       `uri:"id" comment:"编号"` // 编号
	Equipname    string    `json:"equipname" comment:"设备名"`
	Equipcode    string    `json:"equipcode" comment:"设备编号"`
	Faultcode    string    `json:"faultcode" comment:"故障编号"`
	Type         string    `json:"type" comment:"型号"`
	Handlemethod string    `json:"handlemethod" comment:"处理方法"`
	Handlestatus string    `json:"handlestatus" comment:"处理结果"`
	Handletime   time.Time `json:"handletime" comment:"处理时间"`
	Remark       string    `json:"remark" comment:"备注"`
	common.ControlBy
}

func (s *FyHandleUpdateReq) Generate(model *models.FyHandle) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Equipname = s.Equipname
	model.Equipcode = s.Equipcode
	model.Faultcode = s.Faultcode
	model.Type = s.Type
	model.Handlemethod = s.Handlemethod
	model.Handlestatus = s.Handlestatus
	model.Handletime = s.Handletime
	model.Remark = s.Remark
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *FyHandleUpdateReq) GetId() interface{} {
	return s.Id
}

// FyHandleGetReq 功能获取请求参数
type FyHandleGetReq struct {
	Id int `uri:"id"`
}

func (s *FyHandleGetReq) GetId() interface{} {
	return s.Id
}

// FyHandleDeleteReq 功能删除请求参数
type FyHandleDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *FyHandleDeleteReq) GetId() interface{} {
	return s.Ids
}
