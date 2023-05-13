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

type FyData struct {
	Id        int
	Equipname string
	Datax     string
	Datay     string
	// CreateBy  string
	// UpdateBy  string
	CreatedAt string
	// UpdatedAt string
	// DeletedAt string

	// Id        string `form:"idOrder"  search:"type:order;column:id;table:fy_data"`
	// Equipname string `form:"equipnameOrder"  search:"type:order;column:equipname;table:fy_data"`
	// Datax     string `form:"dataxOrder"  search:"type:order;column:datax;table:fy_data"`
	// Datay     string `form:"datayOrder"  search:"type:order;column:datay;table:fy_data"`
	// CreateBy  string `form:"createByOrder"  search:"type:order;column:create_by;table:fy_data"`
	// UpdateBy  string `form:"updateByOrder"  search:"type:order;column:update_by;table:fy_data"`
	// CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:fy_data"`
	// UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:fy_data"`
	// DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:fy_data"`
}

func decimalf(num float64) float64 {
	num, _ = strconv.ParseFloat(fmt.Sprintf("%.4f", num), 64)
	return num
}

func main() {

	// fmt.Println("start")
	// timer := time.NewTimer(2000 * time.Microsecond)
	// // go func() {
	// // 	<-timer.C
	// // 	fmt.Println("3 seconds after")
	// // }()
	// // fmt.Println("停止任务结果:", timer.Stop())
	// // time.Sleep(4 * time.Second)
	// // fmt.Println("程序结束了")
	// timer.Reset(2 * time.Second)

	// for i := 0; i < 3; i++ {

	// 	x := rand.Intn(50)
	// 	y := rand.Intn(50)
	// 	var fydataList []FyData
	// 	fydataList = append(fydataList, FyData{
	// 		// Id:        fmt.Sprint(i),
	// 		Equipname: "jerry",
	// 		Datax:     fmt.Sprint(x),
	// 		Datay:     fmt.Sprint(y),
	// 	})
	// 	DB.Create(fydataList)
	// }
	// // fmt.Println(1 * i)
	// time.Sleep(2000 * time.Microsecond)

	for i := 0; i < 200000; i++ {
		time.Sleep(50 * time.Millisecond)
		// fmt.Println(i)
		// x := rand.Float64() * 10
		// y := rand.Float64() * 100

		var fydata FyData
		DB.Last(&fydata, "equipname = ?", "eq1")
		fmt.Println("id=", fydata.Id)
		// fmt.Println("id=", fydata.Id, "name=", fydata.Equipname)
		// fmt.Println(fydata)
		xo := fydata.Datax
		yo := fydata.Datay
		// fmt.Println(xo, yo)
		xm, _ := strconv.ParseFloat(xo, 64)
		ym, _ := strconv.ParseFloat(yo, 64)
		x := xm + 0.1
		y := ym + rand.Float64()*200 - 100
		// y := 1500 + rand.Float64()*500 - 250 + math.Sin(x)*50
		x = decimalf(x + 0.000001)
		y = decimalf(y)
		// fmt.Println(x, y)
		fydata = FyData{}

		var fydataList []FyData
		fydataList = append(fydataList, FyData{
			// Id:        fmt.Sprint(i),
			Equipname: "eq1",
			Datax:     fmt.Sprint(x),
			Datay:     fmt.Sprint(y),
			CreatedAt: string([]byte(time.Now().String())[:23]),
		})
		DB.Create(fydataList)
		// fmt.Println(x, y, fydataList)

		DB.First(&fydata, "equipname = ?", "eq1")
		DB.Delete(&fydata)
		fydata = FyData{}

		//eq2表格操作
		DB.Last(&fydata, "equipname = ?", "eq2")
		// fmt.Println("id=", fydata.Id, "name=", fydata.Equipname)
		xo2 := fydata.Datax
		yo2 := fydata.Datay
		xm2, _ := strconv.ParseFloat(xo2, 64)
		ym2, _ := strconv.ParseFloat(yo2, 64)
		x2 := xm2 + 0.01
		y2 := ym2 + rand.Float64()*400 - 195
		x2 = decimalf(x2 + 0.000001)
		y2 = decimalf(y2)
		fydata = FyData{}
		// fmt.Println("x2,y2=", x2, y2)
		var fydataList2 []FyData
		fydataList2 = append(fydataList2, FyData{
			// Id:        fmt.Sprint(i),
			Equipname: "eq2",
			Datax:     fmt.Sprint(x2),
			Datay:     fmt.Sprint(y2),
			CreatedAt: string([]byte(time.Now().String())[:23]),
		})
		DB.Create(fydataList2)
		// fmt.Println(x, y, fydataList)

		DB.First(&fydata, "equipname = ?", "eq2")
		DB.Delete(&fydata)
	}

}
