package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"strconv"
	"tiger-go/tiger/configure"
	"time"
)

type DB struct {
	*gorm.DB
	PrimaryKey string
}

func New(name ...string) *DB {
	dsn := getDsn(name...)
	db := mysqlConnection(dsn)
	sqlDB, err := db.DB.DB()
	if err != nil {
		panic(err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}

func mysqlConnection(dsn string) *DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	return &DB{
		db,
		"",
	}
}

func getDsn(name ...string) string {
	defaultGroup := "default"
	if len(name) > 0 && name[0] != "" {
		defaultGroup = name[0]
	}

	dbConfig := configure.New().GetStringMapString(fmt.Sprintf("db.%s", defaultGroup))
	parseTime := true
	if dbConfig["parseTime"] != "" {
		parseTime, _ = strconv.ParseBool(dbConfig["parseTime"])
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%v&loc=Local", dbConfig["username"], dbConfig["password"], dbConfig["host"], dbConfig["port"], dbConfig["dbname"], dbConfig["charset"], parseTime)
}
