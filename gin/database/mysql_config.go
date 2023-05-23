package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"yqc-portal/gin/model"
	"yqc-portal/resource"
)

// db 数据库连接 全局变量
var db *gorm.DB

// GetDb 获取数据库连接
func GetDb() *gorm.DB {
	return db
}

func MysqlInit() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	database := resource.Database
	if database == nil {
		panic("数据库配置读取失败")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		database.UserName,
		database.Password,
		database.Host,
		database.Port,
		database.DbName)
	config := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}
	conn, err := gorm.Open(mysql.New(config), &gorm.Config{})
	if err != nil {
		panic(err) // 抛出异常
	}

	sqlDB, err := conn.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(database.MaxIdleConn)                                     //最大空闲连接数
	sqlDB.SetMaxOpenConns(database.MaxOpenConn)                                     //最大连接数
	sqlDB.SetConnMaxLifetime(time.Duration(database.ConnMaxLifetime) * time.Second) //设置连接空闲超时

	// 初始化数据
	// ORM映射表字段和结构体字段
	// 1.默认情况下，gorm会将结构体名称转换为小写并使用复数形式，
	//例如，User结构体将默认映射到“users”表，
	//可以通过将表名设置为User结构体的TableName方法来覆盖此行为
	// 2.默认情况下，GORM将使用struct字段的名称作为列名，
	//可以通过将字段标记为column来覆盖此行为, 例如：gorm:"column:cloumn_name"
	conn.AutoMigrate(&model.User{}) // 自动迁移 User 结构体 生成 users表 如果表已经存在，则不执行任何操作

	// 创建数据行
	u := []model.User{
		{
			Name:   "渊洁",
			Age:    19,
			Gender: "男",
		},
		{
			Name:   "张三",
			Age:    19,
			Gender: "女",
		},
		{
			Name:   "李四",
			Age:    20,
			Gender: "男",
		},
		{
			Name:   "王五",
			Age:    21,
			Gender: "女",
		},
	}

	conn.CreateInBatches(&u, len(u)) // 批量插入数据 &u 是一个切片 len(u) 是切片的长度
	for idx, result := range u {
		fmt.Println(idx, result.ID) // 打印插入数据的ID
	}

	db = conn // 赋值给全局变量
}
