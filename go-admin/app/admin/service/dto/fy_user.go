package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type FyUserGetPageReq struct {
	dto.Pagination `search:"-"`
	Username       string `form:"username"  search:"type:contains;column:username;table:fy_user" comment:"用户名"`
	Area           string `form:"area"  search:"type:contains;column:area;table:fy_user" comment:"负责地区"`
	Userstatus     string `form:"userstatus"  search:"type:contains;column:userstatus;table:fy_user" comment:"人员状态"`
	Duty           string `form:"duty"  search:"type:contains;column:duty;table:fy_user" comment:"职责"`
	Remark         string `form:"remark"  search:"type:exact;column:remark;table:fy_user" comment:"备注"`
	FyUserOrder
}

type FyUserOrder struct {
	Id         string `form:"idOrder"  search:"type:order;column:id;table:fy_user"`
	Username   string `form:"usernameOrder"  search:"type:order;column:username;table:fy_user"`
	Area       string `form:"areaOrder"  search:"type:order;column:area;table:fy_user"`
	Userstatus string `form:"userstatusOrder"  search:"type:order;column:userstatus;table:fy_user"`
	Duty       string `form:"dutyOrder"  search:"type:order;column:duty;table:fy_user"`
	Remark     string `form:"remarkOrder"  search:"type:order;column:remark;table:fy_user"`
	CreateBy   string `form:"createByOrder"  search:"type:order;column:create_by;table:fy_user"`
	UpdateBy   string `form:"updateByOrder"  search:"type:order;column:update_by;table:fy_user"`
	CreatedAt  string `form:"createdAtOrder"  search:"type:order;column:created_at;table:fy_user"`
	UpdatedAt  string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:fy_user"`
}

func (m *FyUserGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type FyUserInsertReq struct {
	Id         int    `json:"-" comment:"编号"` // 编号
	Username   string `json:"username" comment:"用户名"`
	Area       string `json:"area" comment:"负责地区"`
	Userstatus string `json:"userstatus" comment:"人员状态"`
	Duty       string `json:"duty" comment:"职责"`
	Remark     string `json:"remark" comment:"备注"`
	common.ControlBy
}

func (s *FyUserInsertReq) Generate(model *models.FyUser) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Username = s.Username
	model.Area = s.Area
	model.Userstatus = s.Userstatus
	model.Duty = s.Duty
	model.Remark = s.Remark
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *FyUserInsertReq) GetId() interface{} {
	return s.Id
}

type FyUserUpdateReq struct {
	Id         int    `uri:"id" comment:"编号"` // 编号
	Username   string `json:"username" comment:"用户名"`
	Area       string `json:"area" comment:"负责地区"`
	Userstatus string `json:"userstatus" comment:"人员状态"`
	Duty       string `json:"duty" comment:"职责"`
	Remark     string `json:"remark" comment:"备注"`
	common.ControlBy
}

func (s *FyUserUpdateReq) Generate(model *models.FyUser) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Username = s.Username
	model.Area = s.Area
	model.Userstatus = s.Userstatus
	model.Duty = s.Duty
	model.Remark = s.Remark
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *FyUserUpdateReq) GetId() interface{} {
	return s.Id
}

// FyUserGetReq 功能获取请求参数
type FyUserGetReq struct {
	Id int `uri:"id"`
}

func (s *FyUserGetReq) GetId() interface{} {
	return s.Id
}

// FyUserDeleteReq 功能删除请求参数
type FyUserDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *FyUserDeleteReq) GetId() interface{} {
	return s.Ids
}

//这段代码定义了一些数据传输对象（DTO），用于在应用程序的不同层之间传递数据。其中，FyUserGetPageReq、FyUserInsertReq、FyUserUpdateReq、FyUserGetReq和FyUserDeleteReq分别对应了FyUser服务结构体中的GetPage()、Insert()、Update()、Get()和Remove()方法中的参数。这些DTO的定义中，使用了search标签指定了查询参数的类型、列名和表名，使用了comment标签指定了注释，方便了解每个参数的含义。这些DTO中的Generate()方法用于将DTO转换为模型对象，GetId()方法用于获取DTO中的ID字段。这些DTO的定义是为了在应用程序的不同层之间传递数据，避免了不同层之间直接耦合的情况，提高了应用程序的灵活性和可维护性。
