package tiger

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"mang/tiger/configure"
	"os"
	"time"
)

const (
	VERSION = "v1.0.0"
)

// 声明环境变量
const (
	EnvDebug = "debug"
	ENVProd  = "prod"
)

func New(config configure.Config) {

}

// 设置环境
func SetEnv(env string) {
	if err := os.Setenv("APP_ENV", env); err != nil {
		panic(err)
	}
}

func Run() {
	//config.InitConfig()
}

var config *configure.Config

func Config() *configure.Config {
	if config == nil {
		config = &configure.Config{
			Viper: viper.New(),
		}
	}
	return config
}

func GetDb(name ...string) *gorm.DB {
	defaultGroup := "default"
	if len(name) > 0 && name[0] != "" {
		defaultGroup = name[0]
	}
	dbConfig := Config().GetStringMap("db.default")
	fmt.Println("config",dbConfig, defaultGroup)
	dsn := "root:@tcp(127.0.0.1:3306)/mang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	fmt.Println("db addr:", &db)

	return db
}
