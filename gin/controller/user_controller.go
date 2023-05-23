package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"yqc-portal/gin/model"
	"yqc-portal/gin/service"
)

func GetUserList(c *gin.Context) {
	list := service.GetUserList()
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    list,
	})
}

func UpdateUserNameById(c *gin.Context) {
	user := model.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
		return
	}
	if user.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请输入用户名",
		})
		return
	}
	rowsAffected := service.UpdateUserNameById(user.ID, user.Name)
	if rowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "更新失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
	})

}

func AddUser(c *gin.Context) {
	user := model.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
		return
	}
	if user.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请输入用户名",
		})
		return
	}
	rowsAffected := service.AddUser(user)
	if rowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "新增失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "新增成功",
	})
}

func DeleteUserById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Query("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	rowsAffected := service.DeleteUserById(uint(id))
	if rowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "删除失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
	})
}
