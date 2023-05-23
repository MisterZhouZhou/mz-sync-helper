package routes

import (
	"mz-sync-helper/server/controllers"
	"mz-sync-helper/server/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

// 路由
func NewRouter() *gin.Engine {
	router := gin.Default()
	env := os.Getenv("GO_ENV")
	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
	}
	// router.SetTrustedProxies([]string{"192.168.125.15", "192.168.1.2"})

	// 配置跨域
	router.Use(middleware.Cors())

	// 下载文件
	router.GET("/uploads/:path", controllers.UploadsController)

	v1 := router.Group("/api/v1")
	{
		v1.GET("address", controllers.AddressesController)
		v1.GET("qrcode", controllers.QrCodeController)
		v1.POST("files", controllers.FilesController)
		v1.POST("texts", controllers.TextController)
		// 获取uploads文件上传列表
		v1.GET("uploads", controllers.UploadListController)
	}

	// router.POST("/api/v1/files", controllers.FilesController)

	return router
}
