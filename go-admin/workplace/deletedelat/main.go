package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	username := "root"
	password := "123456"
	host := "127.0.0.1"
	port := 3306
	Dbname := "go-admin"
	timeout := "1000ms"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("数据库连接失败,error=" + err.Error())
	}
	DB = db

}

type Article struct {
	Id        int
	DeletedAt string
}

// func (Article) TableName() string {
// 	return "fy_status"
// }

//	func main() {
//		articles := []Article{}
//		DB.Unscoped().Find(&articles)
//		for _, a := range articles {
//			if len(a.DeletedAt) > 0 {
//				DB.Delete(&a)
//			}
//		}
//	}
func (Article) TableName() string {
	return "fy_status"
}

func main() {
	articles := []Article{}
	DB.Unscoped().Find(&articles)
	for _, a := range articles {
		if len(a.DeletedAt) > 0 {
			DB.Delete(&a)
		}
	}
}
