### 验证规则

- required ：必填
- email：验证字符串是email格式；例：“email”
- url：这将验证字符串值包含有效的网址;例：“url”
- max：字符串最大长度；例：“max=20”
- min:字符串最小长度；例：“min=6”
- excludesall:不能包含特殊字符；例：“excludesall=0x2C”//注意这里用十六进制表示。
- len：字符长度必须等于n，或者数组、切片、map的len值为n，即包含的项目数；例：“len=6”
- eq：数字等于n，或者或者数组、切片、map的len值为n，即包含的项目数；例：“eq=6”
- ne：数字不等于n，或者或者数组、切片、map的len值不等于为n，即包含的项目数不为n，其和eq相反；例：“ne=6”
- gt：数字大于n，或者或者数组、切片、map的len值大于n，即包含的项目数大于n；例：“gt=6”
- gte：数字大于或等于n，或者或者数组、切片、map的len值大于或等于n，即包含的项目数大于或等于n；例：“gte=6”
- lt：数字小于n，或者或者数组、切片、map的len值小于n，即包含的项目数小于n；例：“lt=6”
- lte：数字小于或等于n，或者或者数组、切片、map的len值小于或等于n，即包含的项目数小于或等于n；例：“lte=6”

### 跨字段验证

- eqfield=Field: 必须等于 Field 的值；
- nefield=Field: 必须不等于 Field 的值；
- gtfield=Field: 必须大于 Field 的值；
- gtefield=Field: 必须大于等于 Field 的值；
- ltfield=Field: 必须小于 Field 的值；
- ltefield=Field: 必须小于等于 Field 的值；
- eqcsfield=Other.Field: 必须等于 struct Other 中 Field 的值；
- necsfield=Other.Field: 必须不等于 struct Other 中 Field 的值；
- gtcsfield=Other.Field: 必须大于 struct Other 中 Field 的值；
- gtecsfield=Other.Field: 必须大于等于 struct Other 中 Field 的值；
- ltcsfield=Other.Field: 必须小于 struct Other 中 Field 的值；
- ltecsfield=Other.Field: 必须小于等于 struct Other 中 Field 的值；

```go
type UserReg struct {
    Passwd   string `form:"passwd" json:"passwd" validate:"required,max=20,min=6"`
    Repasswd   string `form:"repasswd" json:"repasswd" validate:"required,max=20,min=6,eqfield=Passwd"`
}

```

### 自定义验证类型
```go
package main

import (
    "fmt"
    "github.com/go-playground/validator/v10"
)

type Users struct {
    Name   string `form:"name" json:"name" validate:"required,CustomValidationErrors"`//包含自定义函数
    Age   uint8 `form:"age" json:"age" validate:"required,gt=18"`
    Passwd   string `form:"passwd" json:"passwd" validate:"required,max=20,min=6"`
    Code   string `form:"code" json:"code" validate:"required,len=6"`
}

func main() {

    users := &Users{
        Name:      "admin",
        Age:        12,
        Passwd:       "123",
        Code:            "123456",
    }
    validate := validator.New()
    //注册自定义函数
    _=validate.RegisterValidation("CustomValidationErrors", CustomValidationErrors)
    err := validate.Struct(users)
    if err != nil {
        for _, err := range err.(validator.ValidationErrors) {
            fmt.Println(err)//Key: 'Users.Name' Error:Field validation for 'Name' failed on the 'CustomValidationErrors' tag
            return
        }
    }
    return
}

func CustomValidationErrors(fl validator.FieldLevel) bool {
return fl.Field().String() != "admin"
}


```

```go
package main

import (
    "fmt"
    "github.com/go-playground/locales/zh"
    ut "github.com/go-playground/universal-translator"
    "github.com/go-playground/validator/v10"
    zh_translations "github.com/go-playground/validator/v10/translations/zh"
    "reflect"
)

type Users struct {
    Name   string `form:"name" json:"name" validate:"required" label:"用户名"`
    Age   uint8 `form:"age" json:"age" validate:"required,gt=18" label:"年龄"`
    Passwd   string `form:"passwd" json:"passwd" validate:"required,max=20,min=6"`
    Code   string `form:"code" json:"code" validate:"required,len=6"`
}

func main() {
    users := &Users{
        Name:      "admin",
        Age:        12,
        Passwd:       "123",
        Code:            "123456",
    }
    uni := ut.New(zh.New())
    trans, _ := uni.GetTranslator("zh")
    validate := validator.New()
    //注册一个函数，获取struct tag里自定义的label作为字段名
    validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
        name:=fld.Tag.Get("label")
        return name
    })
    //注册翻译器
    err := zh_translations.RegisterDefaultTranslations(validate, trans)
    if err!=nil {
        fmt.Println(err)
    }
    err = validate.Struct(users)
    if err != nil {
        for _, err := range err.(validator.ValidationErrors) {
            fmt.Println(err.Translate(trans))//年龄必须大于18
            return
        }
    }
    return
}
```
### 参考
[golang之数据验证validator](https://blog.csdn.net/guyan0319/article/details/105918559/)
[go-playground/validator](https://github.com/go-playground/validator)
[go-playground / locales](https://github.com/go-playground/locales)
[go-playground / universal-translator](https://github.com/go-playground/universal-translator)
