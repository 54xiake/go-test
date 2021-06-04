package test

import (
	"fmt"
	"github.com/54xiake/gotest/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

type UserController struct {
	BaseController
}

func (u *UserController) Create() {
	dsn := "root:asdfghjkl@tcp(127.0.0.1:3306)/basesrv?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err == nil {
		fmt.Println(err)
	}

	//var users = []models.Users{{Username: "jinzhu1"}, {Username: "jinzhu2"}, {Username: "jinzhu3"}}

	userList := make([]models.Users, 0)
	for i := 0; i <= 10; i++ {
		user := models.Users{}
		user.Username = "test" + strconv.Itoa(i)
		user.Password = "123132"
		userList = append(userList, user)
	}

	db.Create(&userList)
	//db.Create(&users)

	//for _, user := range users {
	//	fmt.Println(user.Id) // 1,2,3
	//}

	for _, user := range userList {
		fmt.Println(user.Id) // 1,2,3
	}

	// 根据 `[]map[string]interface{}{}` 批量插入
	db.Model(&models.Users{}).Create([]map[string]interface{}{
		{"Username": "jinzhu_1"},
		{"Username": "jinzhu_2"},
	})

	res := JsonResponse{
		200, "success", nil,
	}
	u.Data["json"] = res
	u.ServeJSON() //对json进行序列化输出
	u.StopRun()
}
