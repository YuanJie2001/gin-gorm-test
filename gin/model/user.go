package model

import "gorm.io/gorm"

type User struct {
	gorm.Model        // gorm.Model包含字段：ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string `gorm:"column:name;type:varchar(255);default:(-)" ` // 注意这里的字段首字母必须大写，否则无法被json包访问
	Age        int8
	Gender     string `gorm:"default:男"`
	Status     int8   `gorm:"default:1"`
}
