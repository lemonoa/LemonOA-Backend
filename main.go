package main

import (
	"fmt"

	"github.com/lemonoa/LemonOA-Go/database"

	"github.com/lemonoa/LemonOA-Go/controller"
	"github.com/lemonoa/LemonOA-Go/middleware"
	"github.com/lemonoa/LemonOA-Go/service"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	// 加载配置文件
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// 初始化数据库连接
	if err := database.InitMySQL(); err != nil {
		panic(fmt.Errorf("failed to initialize database: %w", err))
	}

	// 设置gin模式
	gin.SetMode(viper.GetString("server.mode"))
}

func main() {
	r := gin.Default()

	// 配置跨域中间件
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// 初始化认证服务和控制器
	authService := service.NewAuthService(database.DB)
	authController := controller.NewAuthController(authService)

	// 初始化服务和控制器
	addressBookService := service.NewAddressBookService(database.DB)
	addressBookController := controller.NewAddressBookController(addressBookService)

	notificationService := service.NewNotificationService(database.DB)
	notificationController := controller.NewNotificationController(notificationService)

	approvalService := service.NewApprovalService(database.DB)
	approvalController := controller.NewApprovalController(approvalService)

	todoService := service.NewTodoService(database.DB)
	todoController := controller.NewTodoController(todoService)

	systemService := service.NewSystemService(database.DB)
	systemController := controller.NewSystemController(systemService)

	// 基础数据服务和控制器
	basicCommonService := service.NewBasicCommonService(database.DB)
	basicCommonController := controller.NewBasicCommonController(basicCommonService)

	basicHRService := service.NewBasicHRService(database.DB)
	basicHRController := controller.NewBasicHRController(basicHRService)

	basicAdminService := service.NewBasicAdminService(database.DB)
	basicAdminController := controller.NewBasicAdminController(basicAdminService)

	basicFinanceService := service.NewBasicFinanceService(database.DB)
	basicFinanceController := controller.NewBasicFinanceController(basicFinanceService)

	basicCustomerService := service.NewBasicCustomerService(database.DB)
	basicCustomerController := controller.NewBasicCustomerController(basicCustomerService)

	basicContractService := service.NewBasicContractService(database.DB)
	basicContractController := controller.NewBasicContractController(basicContractService)

	basicProjectService := service.NewBasicProjectService(database.DB)
	basicProjectController := controller.NewBasicProjectController(basicProjectService)

	// 人事管理服务和控制器
	hrService := service.NewHRService(database.DB)
	hrController := controller.NewHRController(hrService)

	// 固定资产管理服务和控制器
	assetService := service.NewAssetService(database.DB)
	assetController := controller.NewAssetController(assetService)

	// 车辆管理服务和控制器
	vehicleService := service.NewVehicleService(database.DB)
	vehicleController := controller.NewVehicleController(vehicleService)

	// 会议室管理服务和控制器
	meetingService := service.NewMeetingService(database.DB)
	meetingController := controller.NewMeetingController(meetingService)

	// 印章管理服务和控制器
	sealService := service.NewSealService(database.DB)
	sealController := controller.NewSealController(sealService)

	// 文档管理服务和控制器
	documentService := service.NewDocumentService(database.DB)
	documentController := controller.NewDocumentController(documentService)

	// 初始化服务
	attendanceService := service.NewAttendanceService(database.DB)

	// 初始化控制器
	attendanceController := controller.NewAttendanceController(attendanceService)

	// 注册认证路由
	authController.RegisterRoutes(r)

	// 需要JWT认证的路由组
	api := r.Group("/api")
	api.Use(middleware.JWT())
	{
		// 注册路由
		addressBookController.RegisterRoutes(r)
		notificationController.RegisterRoutes(r)
		approvalController.RegisterRoutes(r)
		todoController.RegisterRoutes(r)
		systemController.RegisterRoutes(r)

		// 注册基础数据路由
		basicCommonController.RegisterRoutes(r)
		basicHRController.RegisterRoutes(r)
		basicAdminController.RegisterRoutes(r)
		basicFinanceController.RegisterRoutes(r)
		basicCustomerController.RegisterRoutes(r)
		basicContractController.RegisterRoutes(r)
		basicProjectController.RegisterRoutes(r)

		// 注册人事管理路由
		hrController.RegisterRoutes(r)

		// 注册固定资产管理路由
		assetController.RegisterRoutes(r)

		// 注册车辆管理路由
		vehicleController.RegisterRoutes(r)

		// 注册会议室管理路由
		meetingController.RegisterRoutes(r)

		// 注册印章管理路由
		sealController.RegisterRoutes(r)

		// 注册文档管理路由
		documentController.RegisterRoutes(r)

		// 注册考勤管理路由
		attendanceController.RegisterRoutes(r)
	}

	// 启动服务器
	port := viper.GetString("server.port")
	if err := r.Run(":" + port); err != nil {
		panic(err)
	}
}
