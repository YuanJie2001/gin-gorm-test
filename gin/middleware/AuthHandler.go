package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthHandler() gin.HandlerFunc {
	// 连接数据库 准备

	return func(c *gin.Context) {
		// 具体的逻辑判断
		if c.Param("id") == "520" {
			c.Abort() // 终止后续的处理器处理 然后由本处理器出栈返回前端结果
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "不准使用520这个id",
			})
		} else if c.Param("id") != "" {
			// 授权逻辑
			fmt.Println("验证通过,分配权限完毕")
			c.Set("auth", "admin") // 设置一个键值对 用来传递数据
			c.Next()               // 通过验证 继续后续的处理
		} else {
			// 直接放行
			c.Next() // 验证不通过 终止后续的处理
		}

	}
}
