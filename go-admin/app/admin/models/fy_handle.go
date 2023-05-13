package models

import (
	"time"

	"go-admin/common/models"
)

type FyHandle struct {
	models.Model

	Equipname    string    `json:"equipname" gorm:"type:varchar(128);comment:设备名"`
	Equipcode    string    `json:"equipcode" gorm:"type:varchar(255);comment:设备编号"`
	Faultcode    string    `json:"faultcode" gorm:"type:int;comment:故障编号"`
	Type         string    `json:"type" gorm:"type:varchar(128);comment:型号"`
	Handlemethod string    `json:"handlemethod" gorm:"type:varchar(255);comment:处理方法"`
	Handlestatus string    `json:"handlestatus" gorm:"type:varchar(128);comment:处理结果"`
	Handletime   time.Time `json:"handletime" gorm:"type:datetime(3);comment:处理时间"`
	Remark       string    `json:"remark" gorm:"type:varchar(255);comment:备注"`
	CreatedAt    time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt    time.Time `json:"updatedAt" gorm:"comment:最后更新时间"`
	models.ControlBy
}

func (FyHandle) TableName() string {
	return "fy_handle"
}

func (e *FyHandle) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *FyHandle) GetId() interface{} {
	return e.Id
}
