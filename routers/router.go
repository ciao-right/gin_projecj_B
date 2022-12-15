package routers

import (
	"gin_project_B/pkg/setting"
	v1 "gin_project_B/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/tags", v1.GetTagList)
		apiV1.POST("/tags", v1.AddTag)
	}
	return r
}
