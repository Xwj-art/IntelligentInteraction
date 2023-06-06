package api

import (
	"Ai/errmsg"
	"Ai/model"
	"Ai/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UpLoad 上传图片接口

func Upload(context *gin.Context) {
	file, fileHeader, err := context.Request.FormFile("file")
	utils.HandleErr(err, "upload错误")

	fileSize := fileHeader.Size

	url, code := model.UpLoadFile(file, fileSize)

	context.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMsg(code),
		"url":     url,
	})
}
