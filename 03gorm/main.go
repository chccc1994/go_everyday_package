package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	ID       uint   `json:"id" sql:"id" gorm:"unique;autoIncrement"`
	UserName string `json:"username" sql:"username"`
	Age      int    `json:"age" sql:"age"  gorm:"default:18"` // 缺省 默认值18
	Gender   string `json:"gender" sql:"gender" gorm:"not null"`
}

func main() {
	// 用户名:密码@tcp(ip:port)/数据库?charset=utf8mb4&parseTime=True&loc=Local
	dsn := "root:123456@tcp(127.0.0.1:3306)/db_test_grom?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("数据库open done")
	}

	// 释放连接
	defer db.Close()

	//var err error
	db.SingularTable(true)
	db.LogMode(true) //启用logger

	// tableUser := new(User)
	//db.AutoMigrate(&User{}) //db.AutoMigrate(&User{},&Product{},&Order{}) //创建表User Product Order

	fmt.Println("数据库表格创建完成.")
	if !db.HasTable(&User{}) {
		fmt.Println("数据库表格不存在,创建表格")
		db.AutoMigrate(&User{})
	}

	// 插入数据
	// user := User{UserName: "zhangsan", Gender: "男"}
	// db.NewRecord(user)
	// db.Create(&user)

	//user := User{UserName: "wangwu", Age: 19, Gender: "男"}
	//db.Select("ID", "UserName", "Age").Create(&user)

	// 查询第一个数据
	// var u User
	// db.First(&u)
	// fmt.Println(u)
	//获取最后一条消息
	// var u User
	// db.Last(&u)

	// 查询到全部的信息 select *from user;
	// var u User
	// result := db.Find(&u)
	// fmt.Println(result.RowsAffected) //返回找到的记录数
	// fmt.Println(result.Error)
	// fmt.Println(u)

	// 检索对象
	// var users []User // 定义 user 数组
	// db.Where("user_name <> ?", "zhangsan").Find(&users)
	// //db.Find(&users)
	// for _, user := range users {
	// 	fmt.Println(user)
	// }

	//struct & map 条件检索
	var users []User
	db.Where(&User{UserName: "zhangsan"}).Find(&users)
	for _, user := range users {
		fmt.Println(user)
	}
	db.Where(map[string]interface{}{"age": 19}).Find(&users)
	for _, user := range users {
		fmt.Println(user)
	}

	//var u User
	// db.First(&u)
	// u.UserName = "wangwu"
	// db.Save(&u)

	// db.First(&u)
	// db.Model(&u).Updates(User{UserName: "litao", Age: 25, Gender: "男"})
	// fmt.Println(u)
	// db.First(&u)
	// db.Delete(&u)

	db.Where("age=?", 19).Delete(User{})
}
