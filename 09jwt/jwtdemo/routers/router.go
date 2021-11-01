//注册路由
package routers

import (
	"github.com/gin-gonic/gin"
	"main.go/pkg/setting"
	"main.go/routers/api"
)

// 初始化引擎，返回 router
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	// jwt auth
	r.GET("/auth", api.GetAuth)

	// apiv1 := r.Group("/api/v1") // 增加jwt
	// apiv1.Use(jwt.JWT())
	// {
	// 	apiv1.GET("/tags", v1.GetTag)

	// 	apiv1.POST("/tags", v1.AddTag)

	// 	apiv1.PUT("/tags/:id", v1.EditTag)

	// 	apiv1.DELETE("/tags/:id", v1.DeleteTag)

	// 	apiv1.GET("/articles", v1.GetArticles)
	// 	apiv1.GET("/articles/:id", v1.GetArticle)
	// 	apiv1.POST("/articles", v1.AddArticles)
	// 	apiv1.PUT("/articles/:id", v1.EditArticle)
	// 	apiv1.DELETE("/articles/:id", v1.DeleteArticle)

	// }

	return r
}
