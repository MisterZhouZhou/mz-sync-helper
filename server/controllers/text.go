package controllers

import (
	"io/ioutil"
	"log"
	"mz-sync-helper/server/config"
	"mz-sync-helper/server/serializer"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TextMsg struct {
	Raw string `json:"raw" form:"raw"`
}

// 发送消息
func TextController(ctx *gin.Context) {
	var textMsg TextMsg
	if err := ctx.ShouldBind(&textMsg); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	// uploads目录路径
	dir := config.UploadsDir
	fileName := uuid.New().String() + ".txt"
	// 创建目录
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	// 拼接全路径写入文件
	fullPath := path.Join("uploads", fileName)
	log.Println("text-path==", filepath.Join(dir, fileName))
	err = ioutil.WriteFile(filepath.Join(dir, fileName), []byte(textMsg.Raw), 0644)
	if err != nil {
		log.Fatal(err)
	}
	ctx.JSON(http.StatusOK, serializer.Response{
		Status: 200,
		Data:   gin.H{"url": "/" + fullPath},
	})
}
