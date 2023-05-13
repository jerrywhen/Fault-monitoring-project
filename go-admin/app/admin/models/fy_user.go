package models

import (
	"go-admin/common/models"
	"time"
)

type FyUser struct {
	models.Model

	Username   string    `json:"username" gorm:"type:varchar(128);comment:用户名"`
	Area       string    `json:"area" gorm:"type:varchar(128);comment:负责地区"`
	Userstatus string    `json:"userstatus" gorm:"type:varchar(128);comment:人员状态"`
	Duty       string    `json:"duty" gorm:"type:varchar(128);comment:职责"`
	Remark     string    `json:"remark" gorm:"type:varchar(255);comment:备注"`
	CreatedAt  time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"comment:最后更新时间"`
	models.ControlBy
}

func (FyUser) TableName() string {
	return "fy_user"
}

func (e *FyUser) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *FyUser) GetId() interface{} {
	return e.Id
}

// 这段代码定义了一个FyUser结构体，它继承了common/models包中的Model结构体，Model结构体中定义了一些通用的数据库字段，例如ID、创建时间、更新时间等。FyUser结构体中还定义了一些字段，例如用户名、负责地区、人员状态、职责和备注等。它们的类型都是varchar，并且都有一个注释。此外，FyUser结构体还实现了Generate()和GetId()方法，这两个方法用于生成ActiveRecord对象和获取ID。其中，Generate()方法用于生成一个新的FyUser对象，GetId()方法用于获取FyUser对象的ID字段。最后，它还定义了TableName()方法，用于返回表名。在这个例子中，它返回的表名是“fy_user”。
