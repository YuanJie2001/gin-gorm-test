package service

import (
	"yqc-portal/gin/database"
	"yqc-portal/gin/model"
)

func GetUserList() []model.User {
	// 获取数据库连接
	db := database.GetDb()
	users := []model.User{}
	result := db.Find(&users)
	if result.Error != nil {
		panic(result.Error) // 抛出异常
	}
	result.Scan(&users) // 将查询结果转换为结构体切片
	return users
}

func UpdateUserNameById(id uint, name string) int64 {
	db := database.GetDb()
	result := db.
		Model(&model.User{}).
		Where("id = ?", id).
		Update("name", name)
	if result.Error != nil {
		panic(result.Error)
	}
	return result.RowsAffected
}

func AddUser(user model.User) int64 {
	db := database.GetDb()
	result := db.Create(&user)
	if result.Error != nil {
		panic(result.Error)
	}
	return result.RowsAffected
}

func DeleteUserById(id uint) int64 {
	db := database.GetDb()
	result := db.
		Where("id = ?", id).
		Delete(&model.User{})
	if result.Error != nil {
		panic(result.Error)
	}
	return result.RowsAffected
}
