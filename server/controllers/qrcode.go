package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

// 生成二维码
func QrCodeController(ctx *gin.Context)  {
	if content := ctx.Query("content"); content != "" {
		png, err := qrcode.Encode(content, qrcode.Medium, 256)
		if err != nil {
			log.Fatal(err)
		}
		ctx.Data(http.StatusOK, "image/png", png)
	} else {
		ctx.Status(http.StatusPreconditionRequired)
	}
}