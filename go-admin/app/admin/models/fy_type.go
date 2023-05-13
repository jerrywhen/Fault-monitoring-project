package models

import (
	"go-admin/common/models"
	"time"
)

type FyType struct {
	models.Model

	Type      string    `json:"type" gorm:"type:varchar(128);comment:型号"`
	Typetext  string    `json:"typetext" gorm:"type:varchar(255);comment:型号说明"`
	Remark    string    `json:"remark" gorm:"type:varchar(255);comment:备注"`
	CreatedAt time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"comment:最后更新时间"`
	models.ControlBy
}

func (FyType) TableName() string {
	return "fy_type"
}

func (e *FyType) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *FyType) GetId() interface{} {
	return e.Id
}
