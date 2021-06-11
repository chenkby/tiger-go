package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"reflect"
	_ "tiger-go/applications/api/boot"
	"tiger-go/applications/api/router"
	"tiger-go/tiger"
	"time"
)

// @title 测试API生成
// @version 1.0
// @description 这是副标题
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

// @x-extension-openapi {"example": "value on a json format"}

var trans ut.Translator
var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}


func main() {

	//tiger.Run()
	//defer func() {
	//	if exception := recover(); exception != nil {
	//		fmt.Println("发生了异常：", exception)
	//	}
	//}()

	//1、初始化配置对象
	// 2、声明tiger，注入config
	// 3、运行


	serverConfig := tiger.Config().GetStringMapString("server")
	s := gin.Default()

	uni := ut.New(zh.New())
	trans, _ = uni.GetTranslator("zh")
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//注册翻译器
		_ = zh_translations.RegisterDefaultTranslations(v, trans)
		//注册自定义函数
		_ = v.RegisterValidation("bookabledate", bookableDate)

		//注册一个函数，获取struct tag里自定义的label作为字段名
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := fld.Tag.Get("label")
			return name
		})
		//根据提供的标记注册翻译
		v.RegisterTranslation("bookabledate", trans, func(ut ut.Translator) error {
			return ut.Add("bookabledate", "{0}不能早于当前时间或{1}格式错误!", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("bookabledate", fe.Field(), fe.Field())
			return t
		})

	}
	router.LoadRoute(s)
	s.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	s.NoRoute(func (c *gin.Context) {
		fmt.Println("404")
	})
	err := s.Run(fmt.Sprintf("%s:%s", serverConfig["address"], serverConfig["port"]))

	if err != nil {
		panic(err)
	}
}
