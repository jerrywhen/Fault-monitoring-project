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

// type FyData struct {
// 	Id        int
// 	Equipname string
// 	Datax     string
// 	Datay     string
// 	// CreateBy  string
// 	// UpdateBy  string
// 	CreatedAt string
// 	// UpdatedAt string
// 	// DeletedAt string

// }
type FyEquipment struct {
	Id        int
	Equipname string
	Equipcode string
	Username  string
	Area      string
	Type      string
	Status    string
	// Faultcount string
	// Remark     string
	CreatedAt string
}

func (FyEquipment) TableName() string {
	return "fy_equipment"
}

func main() {
	username := []string{"www", "eee", "qqq", "user1", "jerry", "jerry2"}
	area := []string{"area", "area1", "area2", "area3", "area4"}

	for i := 0; i < 50; i++ {
		time.Sleep(50 * time.Millisecond)

		var fyEquipment FyEquipment
		DB.Last(&fyEquipment)
		fmt.Println("id=", fyEquipment.Id)
		eqid := fyEquipment.Id
		nid := eqid + 1
		neqname := "eq" + fmt.Sprint(eqid)
		neqcode := "equipcode" + fmt.Sprint(eqid)

		rand.Seed(time.Now().UnixNano())
		nusername := username[rand.Intn(len(username))]
		narea := area[rand.Intn(len(area))]

		fyEquipment = FyEquipment{}

		var fyEquipmentList []FyEquipment
		fyEquipmentList = append(fyEquipmentList, FyEquipment{
			Id:        nid,
			Equipname: neqname,
			Equipcode: neqcode,
			Username:  nusername,
			Area:      narea,
			Type:      "电机a01",
			Status:    "正常运行",
			CreatedAt: string([]byte(time.Now().String())[:23]),
		})
		DB.Create(fyEquipmentList)

		// DB.First(&fyEquipment, "equipname = ?", "eq1")
		// DB.Delete(&fyEquipment)
		// fyEquipment = FyEquipment{}

	}

}
