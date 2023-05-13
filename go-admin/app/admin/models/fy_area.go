package models

import (
	"go-admin/common/models"
	"time"
)

type FyArea struct {
	models.Model

	Area       string    `json:"area" gorm:"type:varchar(128);comment:地区"`
	Username   string    `json:"username" gorm:"type:varchar(128);comment:负责人"`
	Equipcount string    `json:"equipcount" gorm:"type:int;comment:设备数量"`
	Remark     string    `json:"remark" gorm:"type:varchar(255);comment:备注"`
	CreatedAt  time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"comment:最后更新时间"`
	models.ControlBy
}

func (FyArea) TableName() string {
	return "fy_area"
}

func (e *FyArea) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *FyArea) GetId() interface{} {
	return e.Id
}

//本文中的代码是 Go 语言的数据模型，对应了 FyArea 数据库表中的字段。其中，FyArea 结构体继承了 common/models 包中的 Model，包含了一些通用的数据库字段，如 ID、创建时间、更新时间等。
// FyArea 结构体中的其他字段包括了地区、负责人、设备数量、备注等。这些字段与数据库表中的字段进行了对应，使用了 gorm 标签进行了类型、长度、注释等属性的声明。
// Generate 方法用于将 FyArea 对象转换为 ActiveRecord 接口对象，方便进行数据库操作。GetId 方法返回 FyArea 对象的 ID。TableName 方法返回 FyArea 对象对应的数据库表名。
