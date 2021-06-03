package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "tiger-go/applications/admin-api/boot"
	"tiger-go/applications/admin-api/router"
	"tiger-go/tiger"
)

// @title 测试API生成
// @version 1.0
// @description 这是副标题
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

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

	if err := tiger.I18N().InitTrans("zh"); err != nil {
		fmt.Printf("init trans failed, err:%v\n", err)
		return
	}

	serverConfig := tiger.Config().GetStringMapString("server")
	s := gin.Default()

	router.Load(s)
	s.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := s.Run(fmt.Sprintf("%s:%s", serverConfig["address"], serverConfig["port"]))

	if err != nil {
		panic(err)
	}
}
