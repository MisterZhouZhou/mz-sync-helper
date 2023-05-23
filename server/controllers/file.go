package controllers

import (
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

// 上传文件
func FilesController(ctx *gin.Context) {
	file, err := ctx.FormFile("raw")
	if err != nil {
		log.Fatal(err)
	}
	// 项目入口文件所在目录
	// ep, err := os.Executable()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// uploads目录路径
	dir := config.UploadsDir
	// 处理文件名，拼接路径
	filename := uuid.New().String() + filepath.Ext(file.Filename)
	// 文件目录
	// fileDirPath := filepath.Join(dir, "uploads")
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	fullPath := path.Join("uploads", filename)
	log.Println("file-path==", filepath.Join(dir, filename))
	fileErr := ctx.SaveUploadedFile(file, filepath.Join(dir, filename))
	if fileErr != nil {
		log.Fatal(fileErr)
	}
	ctx.JSON(http.StatusOK, serializer.Response{
		Status: 200,
		Data:   gin.H{"url": "/" + fullPath},
	})
}
