package models

import (
	"time"

	"go-admin/common/models"
)

type FyMonitordata struct {
	models.Model

	Equipname string    `json:"equipname" gorm:"type:varchar(128);comment:设备名"`
	Equipcode string    `json:"equipcode" gorm:"type:varchar(128);comment:设备编号"`
	Record1   string    `json:"record1" gorm:"type:double(128,0);comment:数据1"`
	Record2   string    `json:"record2" gorm:"type:double(128,0);comment:数据2"`
	Record3   string    `json:"record3" gorm:"type:double(128,0);comment:数据3"`
	Record4   string    `json:"record4" gorm:"type:double(128,0);comment:数据4"`
	Record5   string    `json:"record5" gorm:"type:double(128,0);comment:数据5"`
	Record6   string    `json:"record6" gorm:"type:double(128,0);comment:数据6"`
	Remark    string    `json:"remark" gorm:"type:varchar(128);comment:备注"`
	CreatedAt time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"comment:最后更新时间"`
	models.ControlBy
}

func (FyMonitordata) TableName() string {
	return "fy_monitordata"
}

func (e *FyMonitordata) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *FyMonitordata) GetId() interface{} {
	return e.Id
}
