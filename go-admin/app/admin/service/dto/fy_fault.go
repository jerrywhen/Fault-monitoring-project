package dto

import (
	"time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type FyFaultGetPageReq struct {
	dto.Pagination `search:"-"`
	Equipname      string    `form:"equipname"  search:"type:contains;column:equipname;table:fy_fault" comment:"设备名"`
	Equipcode      string    `form:"equipcode"  search:"type:contains;column:equipcode;table:fy_fault" comment:"设备编号"`
	Area           string    `form:"area"  search:"type:contains;column:area;table:fy_fault" comment:"地区"`
	Username       string    `form:"username"  search:"type:contains;column:username;table:fy_fault" comment:"负责人"`
	Type           string    `form:"type"  search:"type:contains;column:type;table:fy_fault" comment:"型号"`
	Status         string    `form:"status"  search:"type:contains;column:status;table:fy_fault" comment:"状态"`
	Handlestatus   string    `form:"handlestatus"  search:"type:contains;column:handlestatus;table:fy_fault" comment:"处理结果"`
	Reason         string    `form:"reason"  search:"type:contains;column:reason;table:fy_fault" comment:"故障原因"`
	Faulttime      time.Time `form:"faulttime"  search:"type:contains;column:faulttime;table:fy_fault" comment:"故障时间"`
	Faultcode      string    `form:"faultcode"  search:"type:exact;column:faultcode;table:fy_fault" comment:"故障编号"`
	Remark         string    `form:"remark"  search:"type:contains;column:remark;table:fy_fault" comment:"备注"`
	FyFaultOrder
}

type FyFaultOrder struct {
	Id           string `form:"idOrder"  search:"type:order;column:id;table:fy_fault"`
	Equipname    string `form:"equipnameOrder"  search:"type:order;column:equipname;table:fy_fault"`
	Equipcode    string `form:"equipcodeOrder"  search:"type:order;column:equipcode;table:fy_fault"`
	Area         string `form:"areaOrder"  search:"type:order;column:area;table:fy_fault"`
	Username     string `form:"usernameOrder"  search:"type:order;column:username;table:fy_fault"`
	Type         string `form:"typeOrder"  search:"type:order;column:type;table:fy_fault"`
	Status       string `form:"statusOrder"  search:"type:order;column:status;table:fy_fault"`
	Handlestatus string `form:"handlestatusOrder"  search:"type:order;column:handlestatus;table:fy_fault"`
	Reason       string `form:"reasonOrder"  search:"type:order;column:reason;table:fy_fault"`
	Faulttime    string `form:"faulttimeOrder"  search:"type:order;column:faulttime;table:fy_fault"`
	Faultcode    string `form:"faultcodeOrder"  search:"type:order;column:faultcode;table:fy_fault"`
	Remark       string `form:"remarkOrder"  search:"type:order;column:remark;table:fy_fault"`
	CreateBy     string `form:"createByOrder"  search:"type:order;column:create_by;table:fy_fault"`
	UpdateBy     string `form:"updateByOrder"  search:"type:order;column:update_by;table:fy_fault"`
	CreatedAt    string `form:"createdAtOrder"  search:"type:order;column:created_at;table:fy_fault"`
	UpdatedAt    string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:fy_fault"`
}

func (m *FyFaultGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type FyFaultInsertReq struct {
	Id           int       `json:"-" comment:"编号"` // 编号
	Equipname    string    `json:"equipname" comment:"设备名"`
	Equipcode    string    `json:"equipcode" comment:"设备编号"`
	Area         string    `json:"area" comment:"地区"`
	Username     string    `json:"username" comment:"负责人"`
	Type         string    `json:"type" comment:"型号"`
	Status       string    `json:"status" comment:"状态"`
	Handlestatus string    `json:"handlestatus" comment:"处理结果"`
	Reason       string    `json:"reason" comment:"故障原因"`
	Faulttime    time.Time `json:"faulttime" comment:"故障时间"`
	Faultcode    string    `json:"faultcode" comment:"故障编号"`
	Remark       string    `json:"remark" comment:"备注"`
	common.ControlBy
}

func (s *FyFaultInsertReq) Generate(model *models.FyFault) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Equipname = s.Equipname
	model.Equipcode = s.Equipcode
	model.Area = s.Area
	model.Username = s.Username
	model.Type = s.Type
	model.Status = s.Status
	model.Handlestatus = s.Handlestatus
	model.Reason = s.Reason
	model.Faulttime = s.Faulttime
	model.Faultcode = s.Faultcode
	model.Remark = s.Remark
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *FyFaultInsertReq) GetId() interface{} {
	return s.Id
}

type FyFaultUpdateReq struct {
	Id           int       `uri:"id" comment:"编号"` // 编号
	Equipname    string    `json:"equipname" comment:"设备名"`
	Equipcode    string    `json:"equipcode" comment:"设备编号"`
	Area         string    `json:"area" comment:"地区"`
	Username     string    `json:"username" comment:"负责人"`
	Type         string    `json:"type" comment:"型号"`
	Status       string    `json:"status" comment:"状态"`
	Handlestatus string    `json:"handlestatus" comment:"处理结果"`
	Reason       string    `json:"reason" comment:"故障原因"`
	Faulttime    time.Time `json:"faulttime" comment:"故障时间"`
	Faultcode    string    `json:"faultcode" comment:"故障编号"`
	Remark       string    `json:"remark" comment:"备注"`
	common.ControlBy
}

func (s *FyFaultUpdateReq) Generate(model *models.FyFault) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Equipname = s.Equipname
	model.Equipcode = s.Equipcode
	model.Area = s.Area
	model.Username = s.Username
	model.Type = s.Type
	model.Status = s.Status
	model.Handlestatus = s.Handlestatus
	model.Reason = s.Reason
	model.Faulttime = s.Faulttime
	model.Faultcode = s.Faultcode
	model.Remark = s.Remark
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *FyFaultUpdateReq) GetId() interface{} {
	return s.Id
}

// FyFaultGetReq 功能获取请求参数
type FyFaultGetReq struct {
	Id int `uri:"id"`
}

func (s *FyFaultGetReq) GetId() interface{} {
	return s.Id
}

// FyFaultDeleteReq 功能删除请求参数
type FyFaultDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *FyFaultDeleteReq) GetId() interface{} {
	return s.Ids
}
