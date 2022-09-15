package router

import (
	"github.com/e421083458/gin_scaffold/docs"
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/agent/agentcontroller"
	"github.com/noovertime7/gin-mysqlbak/conf"
	"github.com/noovertime7/gin-mysqlbak/controller"
	"github.com/noovertime7/gin-mysqlbak/middleware"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
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

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	// programatically set swagger info
	docs.SwaggerInfo.Title = conf.GetStringConf("swagger", "title")
	docs.SwaggerInfo.Description = conf.GetStringConf("swagger", "desc")
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = conf.GetStringConf("swagger", "host")
	docs.SwaggerInfo.BasePath = conf.GetStringConf("swagger", "base_path")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router := gin.Default()
	router.Use(middlewares...)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	adminLoginRouter := router.Group("/admin_login")
	adminLoginRouter.Use(
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.TranslationMiddleware(),
	)
	{
		controller.AdminLoginRegister(adminLoginRouter)
	}

	adminRouter := router.Group("/admin")
	adminRouter.Use(
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.JWTAuth(),
		middleware.TranslationMiddleware(),
	)
	{
		controller.AdminRegister(adminRouter)
	}

	taskRouter := router.Group("/task")
	taskRouter.Use(
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.JWTAuth(),
		middleware.TranslationMiddleware(),
	)
	{
		controller.TaskRegister(taskRouter)
	}

	BakRouter := router.Group("/bak")
	BakRouter.Use(
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.JWTAuth(),
		middleware.TranslationMiddleware(),
	)
	{
		controller.BakRegister(BakRouter)
	}

	PublicRouter := router.Group("/public")
	PublicRouter.Use(
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.TranslationMiddleware())
	{
		controller.PublicRegister(PublicRouter)
		agentcontroller.AgentRegister(PublicRouter)
	}

	dashRouter := router.Group("/dashboard")
	dashRouter.Use(
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.JWTAuth(),
		middleware.TranslationMiddleware())
	{
		controller.DashboardRegister(dashRouter)
	}

	hostRouter := router.Group("/host")
	hostRouter.Use(
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.JWTAuth(),
		middleware.TranslationMiddleware())
	{
		controller.HostRegister(hostRouter)
	}

	//Agent相关路由
	AgentRouter := router.Group("/agent")
	AgentRouter.Use(
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.JWTAuth(),
		middleware.TranslationMiddleware())
	{
		agentcontroller.AgentTaskRegister(AgentRouter)
		agentcontroller.AgentHostRegister(AgentRouter)
		agentcontroller.BakHistoryRegister(AgentRouter)
		agentcontroller.BakRegister(AgentRouter)
		agentcontroller.DashBoardRegister(AgentRouter)
	}
	//es备份相关路由
	EsRouter := router.Group("/agent/es")
	EsRouter.Use(
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.JWTAuth(),
		middleware.TranslationMiddleware(),
	)
	{
		agentcontroller.EsTaskRegister(EsRouter)
		agentcontroller.EsBakRegister(EsRouter)
		agentcontroller.EsHistoryRegister(EsRouter)
	}

	return router
}
