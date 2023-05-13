package models

import (
	"go-admin/common/models"
)

type ArticleA struct {
	models.Model

	Username string `json:"username" gorm:"type:varchar(64);comment:用户名"`
	Area     string `json:"area" gorm:"type:varchar(128);comment:地区"`
	Type     string `json:"type" gorm:"type:varchar(128);comment:型号"`
	Status   string `json:"status" gorm:"type:varchar(4);comment:状态"`
	Remark   string `json:"remark" gorm:"type:varchar(255);comment:备注"`
	models.ModelTime
	models.ControlBy
}

func (ArticleA) TableName() string {
	return "article_a"
}

func (e *ArticleA) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *ArticleA) GetId() interface{} {
	return e.Id
}
