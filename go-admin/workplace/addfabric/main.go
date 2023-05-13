package main

import (
	"fmt"
	"math/rand"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	rand.Seed(time.Now().Unix())
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

type MyTable struct {
	Id     int
	Fabric string
	Raster string
}

func (MyTable) TableName() string {
	return "fy_equipment"
}

func main() {
	migrateTable(DB)
}

// 在创建表时为 fabric 和 raster 字段添加默认值
func migrateTable(DB *gorm.DB) error {

	// 批量更新所有记录的 fabric 和 raster 字段
	fabrics := []string{"fabric1", "fabric2", "fabric3", "fabric4", "fabric5"}
	// rasters := []string{"raster1", "raster2", "raster3", "raster4", "raster5", "raster6", "raster7", "raster8", "raster9"}
	rasters := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

	var records []MyTable
	if err := DB.Find(&records).Error; err != nil {
		return err
	}

	for i := range records {
		fmt.Println(records[i])
		updateFields := map[string]interface{}{
			"Fabric": fabrics[rand.Intn(len(fabrics))],
			"Raster": rasters[rand.Intn(len(rasters))],
		}
		if err := DB.Model(&records[i]).Where("id = ?", records[i].Id).Updates(updateFields).Error; err != nil {
			return err
		}
		fmt.Println("for之后", records[i])
	}

	return nil
}
