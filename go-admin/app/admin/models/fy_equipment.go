package models

import (
	"go-admin/common/models"
	"time"
)

type FyEquipment struct {
	models.Model

	Equipname  string    `json:"equipname" gorm:"type:varchar(128);comment:设备名"`
	Equipcode  string    `json:"equipcode" gorm:"type:varchar(255);comment:设备编码"`
	Username   string    `json:"username" gorm:"type:varchar(128);comment:负责人"`
	Area       string    `json:"area" gorm:"type:varchar(128);comment:地区"`
	Type       string    `json:"type" gorm:"type:varchar(128);comment:型号"`
	Status     string    `json:"status" gorm:"type:varchar(128);comment:状态"`
	Fabric     string    `json:"fabric" gorm:"type:varchar(128);comment:光纤编号"`
	Raster     string    `json:"raster" gorm:"type:varchar(128);comment:光栅编号"`
	Faultcount string    `json:"faultcount" gorm:"type:int;comment:故障次数"`
	Remark     string    `json:"remark" gorm:"type:varchar(255);comment:备注"`
	CreatedAt  time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"comment:最后更新时间"`
	models.ControlBy
}

func (FyEquipment) TableName() string {
	return "fy_equipment"
}

func (e *FyEquipment) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *FyEquipment) GetId() interface{} {
	return e.Id
}
