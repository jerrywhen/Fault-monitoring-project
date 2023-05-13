package models

import "gorm.io/gorm/schema"

type ActiveRecord interface {
	schema.Tabler
	SetCreateBy(createBy int)
	SetUpdateBy(updateBy int)
	Generate() ActiveRecord
	GetId() interface{}
}

//这段代码定义了一个ActiveRecord接口，它定义了四个方法：schema.Tabler、SetCreateBy()、SetUpdateBy()、Generate()和GetId()。schema.Tabler方法是从gorm.io/gorm/schema包中引入的，它定义了一个表名方法，任何实现了schema.Tabler接口的对象都可以返回它们自己的表名。SetCreateBy()和SetUpdateBy()方法用于设置创建人和更新人的ID，这两个方法在实现的结构体中会根据业务逻辑进行具体实现。Generate()方法用于生成一个新的ActiveRecord对象，它在实现的结构体中会根据具体业务逻辑生成新的对象。GetId()方法用于获取对象的ID，它在实现的结构体中会返回对象的ID字段。这个接口的定义是为了规范化模型对象的基本操作，使得模型对象可以被更加灵活地操作。
