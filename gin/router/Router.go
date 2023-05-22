package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yqc-portal/gin/controller"
	"yqc-portal/gin/resource"
)
import "yqc-portal/gin/middleware"

func InitRouter() {
	router := gin.Default()
	// 配置「跨域中间件」
	router.Use(middleware.Cors())
	// 使用 session(cookie-based)
	//router.Use(sessions.Sessions("myyyyysession", Sessions.Store))

	// 路由组 v1 对前后端数据交互和后端处理器的基本练习
	v1 := router.Group("/v1", middleware.StatCost())
	//v1.Use(middleware.AuthHandler()) // 给该路由组添加授权中间件
	{
		v1_book := v1.Group("/book")
		// 注册一个路由和处理函数
		// GET is a shortcut for router.Handle("GET", path, handlers).
		v1_book.GET("/", controller.GetBook)
		v1_book.PUT("/", controller.UpdateBook)
		v1_book.POST("/", controller.PostBook)
		v1_book.DELETE("/", controller.DeleteBook)

		// query string
		v1_book.GET("/query", controller.GetBookByIdAndName)
		// form-data
		v1_book.POST("/form", controller.PostBookForm)
		// pathVariable
		v1_book.GET("/path/:id/:name", controller.GetBookByPath)
		// json
		v1_book.POST("/json", controller.PostBookJson)
		// upload file multipart/form-data
		v1_book.POST("/upload", controller.UploadFile)
		// upload files multipart/form-data
		v1_book.POST("/uploads", controller.UploadFiles)
		// http重定向
		v1_book.GET("/redirect", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently,
				"https://www.baidu.com")
		})
		// 路由重定向
		v1_book.GET("/redirect2", func(c *gin.Context) {
			c.Request.URL.Path = "/"
			router.HandleContext(c)
		})

		// 全路由匹配 any<=>@RequestMapping
		v1_book.Any("/any", func(c *gin.Context) {
			switch c.Request.Method {
			case http.MethodGet:
				c.JSON(http.StatusOK, gin.H{
					"message": "GET",
				})
			case http.MethodPost:
				c.JSON(http.StatusOK, gin.H{
					"message": "POST",
				})
			default:
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "请求类型错误",
				})
			}
			c.JSON(http.StatusOK, gin.H{
				"message": "any",
			})
		})

		// 多处理器
		v1_book.GET("/userInfoByIdAndName",
			middleware.AuthHandler(),      // 先认证处理
			controller.GetBookByIdAndName) // 业务处理

	}

	// 路由2 v2 对后端与数据库的交互练习
	v2 := router.Group("/v2")
	v2.Use(middleware.StatCost()) // 给该路由组添加统计耗时中间件
	{
		v2_user := v2.Group("/user")
		// localhost:8080/v2/ GET
		v2_user.GET("/", controller.GetUserList)
		// localhost:8080/v2/updateUserNameById POST {"id":1,"name":"test"}
		v2_user.POST("/updateUserNameById", controller.UpdateUserNameById)
		// localhost:8080/v2/user/add POST {"name":"test01"}
		v2_user.POST("/add", controller.AddUser)
		// localhost:8080/v2/user?id=21 DELETE
		v2_user.DELETE("/", controller.DeleteUserById)

	}

	// 访问不存在的路由
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "这个接口不存在",
		})
	})
	router.Run(resource.Server.Port)
}
