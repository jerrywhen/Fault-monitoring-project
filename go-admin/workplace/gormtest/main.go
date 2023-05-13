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

// type Jiance_test struct {
// 	Id   int64  `gorm:"size:10" db:"id"`
// 	Name string `gorm:"size:16" db:"name" `
// 	Use  string `gorm:"size:64" db:"use"`
// 	Area string `gorm:"size:128" db:"area"`
// }

type Students struct {
	Id     uint   `gorm:"size:3" json:"id"`
	Name   string `gorm:"size:8" json:"name"`
	Age    int    `gorm:"size:3" json:"age"`
	Gender bool
	Email  *string `gorm:"size:32" json:"email"`
}

// 钩子函数，数据插入前做一些事情
// func (user *Students) BeforeCreate(tx *gorm.DB) (err error) {
// 	email := "test@qq.com"
// 	user.Email = &email

// 	return nil
// }

func main() {
	//日期提取
	//2023-03-14 22:42:40.275
	// t := time.Now().String()
	// t = string([]byte(t)[:23])
	// fmt.Println(t)

	// s2 := s[1:3]
	// fmt.Println(s2)
	// fmt.Printf("%T", s2)

	// fmt.Println(s1)

	//确认连接数据库
	fmt.Println(DB)

	//创建student表
	//DB.AutoMigrate(&Students{})

	//单行插入
	// email := "813680666@qq.com"
	// s1 := Students{
	// 	Name:   "jerry",
	// 	Age:    21,
	// 	Gender: true,
	// 	Email:  &email,
	// // }
	// err := DB.Create(&s1).Error
	// fmt.Println(err)

	//批量插入
	// email := "813680666@qq.com"
	// var StudentList []Students
	// for i := 0; i < 10; i++ {
	// 	StudentList = append(StudentList, Students{
	// 		Name:   fmt.Sprintf("jerry%d", i+1),
	// 		Age:    21 + i + rand.Intn(10) - 5,
	// 		Gender: true,
	// 		Email:  &email,
	// 	})
	// 	fmt.Println(rand.Intn(10), StudentList[i].Age)
	// }
	// DB.Create(StudentList)

	//单条记录查询
	//内置查询
	var students Students
	DB.Take(&students)
	fmt.Println(students)
	// students = Students{}
	// DB.First(&students)
	// fmt.Println(students)
	// students = Students{}
	// DB.Last(&students)
	// fmt.Println(students)
	// students = Students{}

	//条件查询
	// DB.Take(&students, "name = ?", "jerry9")
	// fmt.Println(students)
	// students.Id = 16
	// DB.Take(&students)
	// fmt.Println(students)

	//返回查询数量
	// count := DB.Find(&students).RowsAffected
	// fmt.Println(count)

	//多条记录查询
	// var studentList []Students
	// DB.Find(&studentList)
	// for _, students := range studentList {
	// 	fmt.Println(students)
	// }
	// count := DB.Find(&studentList).RowsAffected
	// fmt.Println(count)
	// data, _ := json.Marshal(studentList)
	// fmt.Println(string(data))
	// DB.Find(&studentList, []int{4, 6, 7, 12, 8})
	// DB.Find(&studentList, "name in ?", []string{"jerry", "jerry8", "jerry6"})
	// fmt.Println(studentList)

	//save更新
	// var student Students
	// DB.Take(&student, 21)
	// fmt.Println(&student, 21)
	// student.Name = "jerry110"
	// fmt.Println(&student, 21)

	//update 批量更新
	// var studentList []Students
	// DB.Find(&studentList, []int{1, 2, 3, 4, 5}).Update("gender", false)
	// fmt.Println(&studentList)

	//该方法不更新0值（gender）
	// email := "888866666@qq.com"
	// DB.Model(&Students{}).Where("name = ?", "jerry9").Updates(Students{
	// 	Email:  &email,
	// 	Gender: false,
	// })
	// DB.Find(&studentList, "name in ?", []string{"jerry9"})
	// fmt.Println(&studentList)

	//使用select更新0值（gender）
	// email := "888876666@qq.com"
	// DB.Model(&Students{}).Where("name = ?", "jerry9").Select("gender", "email").Updates(Students{
	// 	Email:  &email,
	// 	Gender: false,
	// })
	// DB.Find(&studentList, "name in ?", []string{"jerry9"})
	// fmt.Println(&studentList)

	//删除delete操作
	// var students Students
	// DB.Delete(&students, []int{12, 13})

	//删除第一行
	// var students Students
	// DB.Take(&students)
	// DB.Delete(&students)

	// DB.Create(&Students{
	// 	Name: "yrrej",
	// 	Age:  13,
	// })
}
