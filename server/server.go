package server

import (
	"fmt"
	"io/fs"
	"log"
	"mz-sync-helper/server/routes"
	"mz-sync-helper/server/ws"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Run(port int, assets fs.FS) {
	// router
	router := routes.NewRouter()

	// websocket
	hub := ws.NewHub()
	go hub.Run()
	router.GET("/ws", func(c *gin.Context) {
		ws.HttpController(c, hub)
	})

	// 获取环境模式
	env := os.Getenv("GO_ENV")
	println("env==", env)
	// env := os.Getenv("ENV")
	// // release打包环境 去除打印
	// // 前端release加载dist
	// if env == "production" {
	// 	os.Setenv("GIN_MODE", "release")
	// 	gin.SetMode(gin.ReleaseMode)
	// 	// gin.SetMode(gin.ReleaseMode)
	// 	// gin.DisableConsoleColor()

	// 	// exe/../Resource/upload
	// 	// dir := config.FrontedDistDir
	// 	// router.StaticFS("/static", http.Dir(dir))

	// 	// 放在里面不行？？？
	// 	fp, _ := fs.Sub(assets, "frontend/dist")
	// 	router.StaticFS("/static", http.FS(fp))
	// }
	// 放在里面不行，先这么处理，以后看看是否有解决方案
	fp, _ := fs.Sub(assets, "frontend/dist")
	router.StaticFS("/static", http.FS(fp))

	// 运行服务
	runErr := router.Run(fmt.Sprintf(":%d", port))
	if runErr != nil {
		log.Fatal(runErr)
	}
}
