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
		apiV1.PUT("/tags", v1.EditTag)
		apiV1.DELETE("/tags", v1.DeleteTag)
		//获取文章列表
		apiV1.GET("/articlesList", v1.GetArticlesList)
		//获取指定文章
		apiV1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiV1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiV1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiV1.DELETE("/articles/:id", v1.DeleteArticle)
	}
	return r
}
