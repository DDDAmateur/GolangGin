package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SimpleUser struct {
	Username string `json:username`
	Password string `json:password`
}

func main() {
	router := gin.Default()

	// 全局加载静态页面
	router.LoadHTMLGlob("templates/*")
	// 全局加载静态资源
	router.Static("/static", "./static")

	// 获取 API参数
	router.GET("/ping/:name/*action", func(ctx *gin.Context) {
		name := ctx.Param("name")
		action := ctx.Param("action")
		ctx.JSON(http.StatusOK, gin.H{
			"message": name + "is " + action,
		})
	})

	// 获取Get 参数
	// 匹配的url格式:  /welcome?name=Jane&age=17
	router.GET("/welcome", func(c *gin.Context) {
		name := c.Query("name")
		age := c.Query("age")
		c.JSON(200, gin.H{
			"name":    name,
			"age":     age,
			"message": "hello! " + name,
		})
	})

	// 创建user路由组
	userGroup := router.Group("/user")
	{
		userGroup.POST("/post_userinfo", func(c *gin.Context) {
			var user SimpleUser
			c.ShouldBindJSON(&user)

			c.JSON(200, gin.H{
				"username": user.Username,
				"password": user.Password,
			})
		})

		userGroup.POST("/login", func(c *gin.Context) {
			var user SimpleUser
			if err := c.ShouldBindJSON(&user); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "无效的JSON格式"})
				return
			}

			if user.Username == "Admin" && user.Password == "123456" {
				c.JSON(http.StatusOK, gin.H{
					"massage": "登录成功",
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"massage": "用户名或密码错误",
				})
			}
		})
	}

	// 返回一个静态页面给前端
	router.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"msg": "JSON内容提取器",
		})
	})

	// 路由重定向
	router.GET("/users", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "https://www.bing.com")
	})

	// 404页面
	router.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusNotFound, "404.html", gin.H{
			"msg": "没有找到该页面",
		})
	})

	// router.Run() //默认监听8080端口 0.0.0.0:8080
	router.Run(":34572") //监听34572

}
