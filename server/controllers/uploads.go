package controllers

import (
	"fmt"
	"log"
	"mz-sync-helper/server/config"
	"mz-sync-helper/server/serializer"
	"mz-sync-helper/server/util"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

type FileOption struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Type string `json:"type"`
}

// 查看上传的文件，返回文件内容
func UploadsController(ctx *gin.Context) {
	if path := ctx.Param("path"); path != "" {
		target := filepath.Join(config.UploadsDir, path)
		fmt.Println("target-path", target)
		ctx.Header("Content-Description", "File Transfer")
		ctx.Header("Content-Transfer-Encoding", "binary")
		ctx.Header("Content-Disposition", "attachment; filename="+path)
		ctx.Header("Content-Type", "application/octet-stream")
		ctx.File(target)
	} else {
		ctx.Status(http.StatusNotFound)
	}
}

// 获取uploads目录列表
func UploadListController(ctx *gin.Context) {
	var files []FileOption
	// uploads目录路径
	dir := config.UploadsDir
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		// 只遍历文件
		if paths := strings.Split(path, "/uploads/"); len(paths) == 2 {
			fileName := paths[1]
			if len(fileName) != 0 {
				// 获取文件后缀
				fileType := util.FileType(filepath.Ext(fileName))
				fileObj := FileOption{
					Name: fileName,
					Path: "/uploads/" + fileName,
					Type: fileType,
				}
				files = append(files, fileObj)
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	ctx.JSON(http.StatusOK, serializer.Response{
		Status: 200,
		Data:   files,
	})
}
