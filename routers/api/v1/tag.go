package v1

import (
	"gin_project_B/models"
	"gin_project_B/pkg/error"
	"gin_project_B/pkg/setting"
	"gin_project_B/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

// GetTagList 获取标签列表：GET("/tags")
func GetTagList(c *gin.Context) {
	name := c.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}
	state := -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}
	code := error.SUCCESS
	data["data"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": data,
	})
}

// AddTag 新建标签：POST("/tags")
func AddTag(c *gin.Context) {
	//_data := make(map[string]interface{})
	name := c.PostForm("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.PostForm("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("最长不能超过一百个字符")
	valid.Required(createdBy, "createdBy").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "createdBy").Message("最长不能超过一百个字符")
	valid.Range(state, 0, 1, "state").Message("状态错误")
	code := error.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistTagByName(name) {
			code = error.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = error.ERROR_EXIST_TAG
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  error.GetMsg(code),
		"code": code,
		"data": 1,
	})
}

//更新指定标签：PUT("/tags/:id")
//删除指定标签：DELETE("/tags/:id")

//c.Query可用于获取?name=test&state=1这类 URL 参数，而c.DefaultQuery则支持设置一个默认值
//code变量使用了e模块的错误编码，这正是先前规划好的错误码，方便排错和识别记录
//util.GetPage保证了各接口的page处理是一致的
//c *gin.Context是Gin很重要的组成部分，可以理解为上下文，它允许我们在中间件之间传递变量、管理流、验证请求的 JSON 和呈现 JSON 响应
