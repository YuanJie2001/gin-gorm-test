package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"yqc-portal/gin/model"
)

func GetBook(c *gin.Context) { // 定义处理函数 c *gin.Context 为gin的上下文 用来获取请求参数
	// gin.H  <=> map[string]interface{}
	data := map[string]interface{}{
		"name":   "渊洁",
		"age":    19,
		"gender": "男",
		"mode":   "GET",
	}
	c.JSON(http.StatusOK, data)
}

func PostBook(c *gin.Context) { // 定义处理函数 c *gin.Context 为gin的上下文 用来获取请求参数
	c.JSON(http.StatusOK, gin.H{
		"name":   "渊洁",
		"age":    19,
		"gender": "男",
		"mode":   "POST",
	})
}
func UpdateBook(c *gin.Context) { // 定义处理函数 c *gin.Context 为gin的上下文 用来获取请求参数
	c.JSON(http.StatusOK, gin.H{
		"name":   "渊洁",
		"age":    19,
		"gender": "男",
		"mode":   "PUT",
	})
}

func DeleteBook(c *gin.Context) { // 定义处理函数 c *gin.Context 为gin的上下文 用来获取请求参数
	c.JSON(http.StatusOK, gin.H{
		"name":   "渊洁",
		"age":    19,
		"gender": "男",
		"mode":   "DELETE",
	})
}

// query string
// /book?id=1&name=渊洁
func GetBookByIdAndName(c *gin.Context) {
	value, exists := c.Get("auth")
	if !exists {
		fmt.Println("匿名用户")
	}

	id := c.Query("id")
	name := c.Query("name")
	name = "我改你名字"
	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"name":  name,
		"你的身份是": value,
	})
}

// form-data
func PostBookForm(c *gin.Context) {
	id := c.PostForm("id")
	name := c.PostForm("name")
	c.JSON(http.StatusOK, gin.H{
		"id":   id,
		"name": name,
	})
}

// pathVariable
// /book/1/渊洁
func GetBookByPath(c *gin.Context) {
	id := c.Param("id")
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"id":   id,
		"name": name,
	})
}

// json
func PostBookJson(c *gin.Context) {
	book := model.Book{}     // 定义一个book结构体 <=> var book model.Book
	err := c.BindJSON(&book) // 将请求的body中的json数据绑定到book结构体中 传入的是指针
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"book":    book,
		})
	}
}

// upload file
func UploadFile(c *gin.Context) {
	// 单个文件 从请求中获取携带的文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = c.SaveUploadedFile(file, "./"+file.Filename)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

// upload files
func UploadFiles(c *gin.Context) {
	// 多个文件 从请求中获取携带的文件
	// MultipartForm() 会解析请求中的文件
	// 限制上传文件的大小
	// gin.Default() 默认限制上传文件大小为32MB
	// gin.New() 默认限制上传文件大小为8MB
	// gin.DisableBindValidation() 不限制上传文件大小
	// gin.SetMode(gin.ReleaseMode) 设置为发布模式 不限制上传文件大小
	// gin.SetMode(gin.DebugMode) 设置为调试模式 限制上传文件大小
	// gin.SetMode(gin.TestMode) 设置为测试模式 不限制上传文件大小
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	files := form.File["files"]
	for _, file := range files {
		err = c.SaveUploadedFile(file, "./"+file.Filename)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
