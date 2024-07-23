package corekit

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQlConfig struct {
	Username string
	Password string
	Address  string
	DBName   string
}

func ConnectMySQL(config *MySQlConfig) *gorm.DB {
	// 构建 MySQL 的 DSN (Data Source Name)
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Address,
		config.DBName,
	)

	var db *gorm.DB
	var err error

	// 使用 Gorm 开启 MySQL 连接，并配置外键约束迁移行为
	if db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	}); err != nil {
		panic(err)
	}

	return db
}
