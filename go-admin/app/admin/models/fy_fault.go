package models

import (
	"time"

	"go-admin/common/models"
)

type FyFault struct {
	models.Model

	Equipname    string    `json:"equipname" gorm:"type:varchar(128);comment:设备名"`
	Equipcode    string    `json:"equipcode" gorm:"type:varchar(255);comment:设备编号"`
	Area         string    `json:"area" gorm:"type:varchar(128);comment:地区"`
	Username     string    `json:"username" gorm:"type:varchar(128);comment:负责人"`
	Type         string    `json:"type" gorm:"type:varchar(128);comment:型号"`
	Status       string    `json:"status" gorm:"type:varchar(128);comment:状态"`
	Handlestatus string    `json:"handlestatus" gorm:"type:varchar(255);comment:处理结果"`
	Reason       string    `json:"reason" gorm:"type:varchar(255);comment:故障原因"`
	Faulttime    time.Time `json:"faulttime" gorm:"type:datetime;comment:故障时间"`
	Faultcode    string    `json:"faultcode" gorm:"type:int;comment:故障编号"`
	Remark       string    `json:"remark" gorm:"type:varchar(255);comment:备注"`
	CreatedAt    time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt    time.Time `json:"updatedAt" gorm:"comment:最后更新时间"`
	models.ControlBy
}

func (FyFault) TableName() string {
	return "fy_fault"
}

func (e *FyFault) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *FyFault) GetId() interface{} {
	return e.Id
}
