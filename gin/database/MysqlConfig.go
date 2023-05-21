package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"yqc-portal/gin/model"
)

func MysqlConfig() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "user:pass@tcp(124.221.87.28:3306)/test_gorm_db?charset=utf8mb4&parseTime=True&loc=Local"
	config := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}
	db, err := gorm.Open(mysql.New(config), &gorm.Config{})
	s, err := db.DB()
	if err != nil {
		panic(err) // 抛出异常
	}
	s.SetMaxIdleConns(10)
	s.SetMaxOpenConns(100)

	// ORM映射表字段和结构体字段
	// 1.默认情况下，gorm会将结构体名称转换为小写并使用复数形式，
	//例如，User结构体将默认映射到“users”表，
	//可以通过将表名设置为User结构体的TableName方法来覆盖此行为
	// 2.默认情况下，GORM将使用struct字段的名称作为列名，
	//可以通过将字段标记为column来覆盖此行为, 例如：gorm:"column:cloumn_name"
	db.AutoMigrate(&model.User{}) // 自动迁移 User 结构体 生成 users表 如果表已经存在，则不执行任何操作

	// 创建数据行
	ul := model.User{
		Name:   "渊洁",
		Age:    19,
		Gender: "男",
	}

	db.Create(&ul) // 创建一条数据行
}
