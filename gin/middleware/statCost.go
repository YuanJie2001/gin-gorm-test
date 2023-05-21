package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next() // 继续后续的处理
		//c.Abort() // 终止后续的处理
		cost := time.Since(start)
		log.Printf("cost:%v", cost)
	}
}
