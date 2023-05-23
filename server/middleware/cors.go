package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// 跨域
func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
			// if origin == "http://127.0.0.1:8000" || origin == "http://localhost:8000" {
			// 	return true
			// } else {
			// 	log.Printf("%v is now allowed", origin)
			// 	return false
			// }
		},
		MaxAge: 12 * time.Hour,
	})
}
