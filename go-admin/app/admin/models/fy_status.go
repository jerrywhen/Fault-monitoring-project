package models

import (
	"go-admin/common/models"
	"time"
)

type FyStatus struct {
	models.Model

	Status    string    `json:"status" gorm:"type:varchar(128);comment:状态"`
	Remark    string    `json:"remark" gorm:"type:varchar(255);comment:备注"`
	CreatedAt time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"comment:最后更新时间"`
	models.ControlBy
}

func (FyStatus) TableName() string {
	return "fy_status"
}

func (e *FyStatus) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *FyStatus) GetId() interface{} {
	return e.Id
}
