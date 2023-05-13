package models

import (
	//  "time"

	"go-admin/common/models"
	"time"
)

type FyData struct {
	models.Model

	Equipname string    `json:"equipname" gorm:"type:varchar(64);comment:设备名"`
	Datax     string    `json:"datax" gorm:"type:double;comment:x数据"`
	Datay     string    `json:"datay" gorm:"type:double;comment:y数据"`
	CreatedAt time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"comment:最后更新时间"`
	models.ControlBy
}

func (FyData) TableName() string {
	return "fy_data"
}

func (e *FyData) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *FyData) GetId() interface{} {
	return e.Id
}
