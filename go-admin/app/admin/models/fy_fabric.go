package models

import (
	"go-admin/common/models"
	"time"
)

type FyFabric struct {
	models.Model

	Fabric      string    `json:"fabric" gorm:"type:varchar(128);comment:光纤编号"`
	Raster      int64     `json:"raster" gorm:"type:int;comment:光栅数量"`
	Instruction string    `json:"instruction" gorm:"type:varchar(128);comment:说明"`
	Remark      string    `json:"remark" gorm:"type:varchar(255);comment:备注"`
	CreatedAt   time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"comment:最后更新时间"`
	models.ControlBy
}

func (FyFabric) TableName() string {
	return "fy_fabric"
}

func (e *FyFabric) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *FyFabric) GetId() interface{} {
	return e.Id
}
