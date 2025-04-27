package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 模拟用户数据
var mockUser = gin.H{
	"Username":     "测试用户",
	"Email":        "test@example.com",
	"QQ":           "123456789",
	"Points":       100,
	"IsRegistered": false,
	"Group":        "A组",
	"Address":      "测试地址",
	"Phone":        "13800138000",
}

func init() {
	// 静态文件服务
	engine.LoadHTMLGlob("app/views/**/*")
	engine.Static("/static", "./static")

	// 页面路由
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "明日方舟斗蛐蛐活动",
			"user":  nil,
		})
	})

	engine.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "登录",
			"user":  nil,
		})
	})

	engine.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", gin.H{
			"title": "注册",
			"user":  nil,
		})
	})

	engine.GET("/profile", func(c *gin.Context) {
		c.HTML(http.StatusOK, "profile.html", gin.H{
			"title": "个人中心",
			"user":  mockUser,
		})
	})

	// 模拟 API 响应
	engine.POST("/v1/common/register", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "注册成功",
			"data":    mockUser,
		})
	})

	engine.POST("/v1/common/address", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "地址更新成功",
		})
	})

	engine.POST("/v1/cricket-fighting-6th/license-accept", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "已接受活动规则",
		})
	})

	engine.POST("/v1/cricket-fighting-6th/participant", func(c *gin.Context) {
		mockUser["IsRegistered"] = true
		c.JSON(http.StatusOK, gin.H{
			"message": "报名成功",
		})
	})

	// 其他 API 路由
	v1 := engine.Group("/v1")
	manageApi := v1.Group("/manage")
	manageApi.POST("/group", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "分组更新成功",
		})
	})
	manageApi.POST("/reward", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "奖励发放成功",
		})
	})
}
