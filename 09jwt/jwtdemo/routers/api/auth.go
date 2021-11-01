package api

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"main.go/models"
	"main.go/pkg/e"
	"main.go/pkg/logging"
	"main.go/pkg/util"
)

type auth struct {
	Username string `valid:"Required;MaxSize(50)"`
	Password string `valid:"Required;MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	data := make(map[string]interface{})
	code := e.INVALID_PARAMS

	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	if ok {
		isExit := models.CheckAuth(username, password)
		if isExit {
			token, err := util.GenerateToken(username, password)

			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			//log.Println(err.Key, err.Message)
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}
