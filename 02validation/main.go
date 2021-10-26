package main

import (
	"fmt"
	"reflect"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

type Users struct {
	//Name     string `form:"name" json:"name" validate:"required,CustomValidationErrors"` //自定义函数
	Name     string `form:"name" json:"name" `
	Age      uint8  `form:"age" json:"age" validate:"required,gt=18" label:"年龄"`
	Phone    string `form:"phone" json:"phone" validate:"required"`
	Passwd   string `form:"passwd" json:"passwd" validate:"required,max=20,min=6"`
	Repasswd string `form:"repasswd" json:"repasswd" validate:"required,max=20,min=6,eqfield=Passwd"`
	Code     string `form:"code" json:"code" validate:"required,len=6"`
}

// func CustomValidationErrors(fl validator.FieldLevel) bool {
// 	fmt.Println(fl.Field().String())
// 	return fl.Field().String() != "admin"
// }

func main() {

	user := Users{
		Name:     "admin",
		Age:      12,
		Phone:    "133123498760",
		Passwd:   "111222",
		Repasswd: "111222",
		Code:     "126122",
	}
	// 中文转换输出
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")

	validate := validator.New()
	//注册一个函数，获取struct tag里自定义的label作为字段名
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		return name
	})

	//验证器注册翻译器
	err := zh_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println(err)
	}

	//	_ = validate.RegisterValidation("CustomValidationErrors", CustomValidationErrors) //自定义验证
	//Key: 'Users.Name' Error:Field validation for 'Name' failed on the 'CustomValidationErrors' tag

	err = validate.Struct(user)
	if err != nil {
		//return
		for _, err := range err.(validator.ValidationErrors) {
			//fmt.Println(err)
			fmt.Println(err.Translate(trans))
			return
		}
	}
	return
}
