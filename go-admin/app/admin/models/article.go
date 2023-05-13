package models

import (
	"time"

	"go-admin/common/models"
)

type Article struct {
	models.Model

	Username  string    `json:"username" gorm:"type:varchar(64);comment:用户名"`
	Area      string    `json:"area" gorm:"type:varchar(128);comment:地区"`
	Type      string    `json:"type" gorm:"type:varchar(128);comment:型号"`
	Status    string    `json:"status" gorm:"type:varchar(4);comment:状态"`
	Remark    string    `json:"remark" gorm:"type:varchar(255);comment:备注"`
	CreatedAt time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"comment:最后更新时间"`
	models.ControlBy
}

func (Article) TableName() string {
	return "article"
}

func (e *Article) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Article) GetId() interface{} {
	return e.Id
}
