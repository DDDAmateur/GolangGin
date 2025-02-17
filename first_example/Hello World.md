# 创建第一个Gin程序

参考文档：	[Quickstart | Gin Web Framework](https://gin-gonic.com/docs/quickstart/)

## 1、安装Gin

```powershell
go get -u github.com/gin-gonic/gin
```

## 2、在项目中引用

```go
go install github.com/gin-gonic/gin@latest
```

## 3、Hello World

```go
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
```

## 知识扩展

### `gin.New()` 和 `gin.Default()` 的区别

- **`gin.New()`**  创建一个纯净的Engine实例，**不包含任何默认中间件**。适用于需要完全自定义中间件链的场景，例如：
  
   ```go 
  router := gin.New()
  router.Use(customLogger, customRecovery) // 手动添加所需中间件 

- `gin.Default()` 在 `gin.New()`基础上自动添加两个关键中间件：

  - **Logger中间件**：记录请求日志。

  - **Recovery中间件**：自动捕获panic并返回500响应，防止服务崩溃

### 获取 `API` 参数

```go
// 此规则能够匹配/user/john这种格式，但不能匹配/user/ 或 /user这种格式
router.GET("/ping/:name/*action", func(ctx *gin.Context) {
    name := ctx.Param("name")
    action := ctx.Param("action")
    ctx.JSON(http.StatusOK, gin.H{
        "message": name + "is " + action,
    })
})
```



###  获取`GET`参数

- URL参数可以通过`DefaultQuery()`或`Query()`方法获取
- `DefaultQuery()`若参数不村则，返回默认值，`Query()`若不存在，返回空串
- `API ?name=zs`

```go
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
```

### 获取 Post 参数

- 

### 返回一个静态页面给前端

文件架构：

```
.
├── main.go            // 主程序文件
├── templates/         // HTML 模板目录
│   └── index.html    // 静态页面
└── static/           // 静态资源目录
│   └── css/    // css文件夹
│   	└── index_css.css    // css文件
│   └── js/    // js文件夹
│   	└── script.js    // js文件
```

代码示例：

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 全局加载静态页面
	router.LoadHTMLGlob("templates/*")
	// 全局加载静态资源
	router.Static("/static", "./static")

	// 返回一个静态页面给前端
	router.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"msg": "hello world！第一个GO Web页面",
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
```

### 路由

```go
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

// 创建user路由组
userGroup := router.Group("/user")
{
    userGroup.POST("/post_userinfo", func(c *gin.Context) {
        name := c.PostForm("name")
        password := c.PostForm("password")
        c.JSON(200, gin.H{
            "user": name,
            "pass": password,
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
```

