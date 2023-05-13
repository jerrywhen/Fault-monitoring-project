package main

import (
	"fmt"

	"github.com/go-admin-team/go-admin-core/sdk/api"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"gorm.io/gorm"
)

type Students struct {
	api.Api
}

var DB *gorm.DB

func main() {
	// var students Students
	// DB.First(&students)
	fmt.Println(DB)
	// fmt.Println(students.Context)
	// fmt.Println(Students.Context)

}
