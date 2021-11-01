package main

import (
	"fmt"
	"log"
	"net/http"
	"syscall"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"main.go/pkg/setting"
	"main.go/routers"
)

func testHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "test",
	})
}
func main() {
	fmt.Println("Start")

	endless.DefaultReadTimeOut = setting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.HTTPPort)
	fmt.Printf("端口：%d", setting.HTTPPort)
	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("pid is %d", syscall.Getpid())
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("server err %v", err)
	}
}
