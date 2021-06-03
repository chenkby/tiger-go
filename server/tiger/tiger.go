package tiger

import (
	"os"
	"tiger-go/tiger/configure"
	"tiger-go/tiger/db"
	"tiger-go/tiger/i18n"
)

const (
	VERSION = "v1.0.0"
)

// 声明环境变量
const (
	EnvDev  = "dev"
	EnvTest = "test"
	ENVProd = "prod"
)

// 设置环境
func SetEnv(env string) {
	if err := os.Setenv("APP_ENV", env); err != nil {
		panic(err)
	}
}

func Config() *configure.Config {
	return configure.New()
}

// 返回数据库连接对象(gorm.DB)
func Db(name ...string) *db.DB {
	return db.New(name...)
}

func I18N() *i18n.I18N {
	return i18n.New()
}
