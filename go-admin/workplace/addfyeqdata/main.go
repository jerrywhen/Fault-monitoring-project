package main

import (
	"fmt"
	"math/rand"
	"strconv"
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

type FyMonitordata struct {
	Id        int
	Equipname string
	Equipcode string
	Record1   float64
	Record2   float64
	Record3   float64
	Record4   float64
	// Record5   float64
	// Record6   float64
	Remark    string
	CreatedAt time.Time
	// UpdatedAt time.Time
}

func (FyMonitordata) TableName() string {
	return "fy_monitordata"
}

func decimalf(num float64) float64 {
	num, _ = strconv.ParseFloat(fmt.Sprintf("%.4f", num), 64)
	return num
}

func main() {

	for i := 0; i < 10000; i++ {
		time.Sleep(250 * time.Millisecond)
		// fmt.Println(i)
		// x := rand.Float64() * 10
		// y := rand.Float64() * 100

		var fydata FyMonitordata
		DB.Last(&fydata, "equipcode = ?", "eqcode1")
		fmt.Println("id=", fydata.Id)
		x1 := fydata.Record1
		// fmt.Println("x1xxxxxxxx1", x1)
		x2 := fydata.Record2
		x3 := fydata.Record3
		x4 := fydata.Record4

		x1d := decimalf(x1 + rand.Float64()*200 - 95)
		x2d := decimalf(x2 + rand.Float64()*200 - 95)
		x3d := decimalf(x3 + rand.Float64()*200 - 95)
		x4d := decimalf(x4 + rand.Float64()*200 - 95)

		// fmt.Printf("bbbbbbbb %v:%T", x1d, x1d)
		fmt.Println("insert:", x1d, x2d, x3d, x4d)

		fydata = FyMonitordata{}

		var fydataList []FyMonitordata
		fydataList = append(fydataList, FyMonitordata{
			// Id:        fmt.Sprint(i),
			Equipname: "eq1",
			Equipcode: "eqcode1",
			Record1:   x1d,
			Record2:   x2d,
			Record3:   x3d,
			Record4:   x4d,
			// Record5:   fmt.Sprint(x5d),
			// Record6:   fmt.Sprint(x6d),
			CreatedAt: time.Now(),
		})
		DB.Create(fydataList)
		// fmt.Println(x, y, fydataList)

		DB.First(&fydata, "equipcode = ?", "eqcode1")
		DB.Delete(&fydata)
		fydata = FyMonitordata{}

		// //eq2表格操作
		// DB.Last(&fydata, "equipname = ?", "eq2")
		// // fmt.Println("id=", fydata.Id, "name=", fydata.Equipname)
		// xo2 := fydata.Datax
		// yo2 := fydata.Datay
		// xm2, _ := strconv.ParseFloat(xo2, 64)
		// ym2, _ := strconv.ParseFloat(yo2, 64)
		// x2 := xm2 + 0.01
		// y2 := ym2 + rand.Float64()*400 - 195
		// x2 = decimalf(x2 + 0.000001)
		// y2 = decimalf(y2)
		// fydata = FyMonitordata{}
		// // fmt.Println("x2,y2=", x2, y2)
		// var fydataList2 []FyMonitordata
		// fydataList2 = append(fydataList2, FyMonitordata{
		// 	// Id:        fmt.Sprint(i),
		// 	Equipname: "eq2",
		// 	Datax:     fmt.Sprint(x2),
		// 	Datay:     fmt.Sprint(y2),
		// 	CreatedAt: string([]byte(time.Now().String())[:23]),
		// })
		// DB.Create(fydataList2)
		// // fmt.Println(x, y, fydataList)

		// DB.First(&fydata, "equipname = ?", "eq2")
		// DB.Delete(&fydata)
	}

}
