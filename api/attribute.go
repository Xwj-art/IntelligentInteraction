package api

import (
	"Ai/errmsg"
	"Ai/model"
	"Ai/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AddAttribute(context *gin.Context) {
	var attribute model.Attribute
	context.ShouldBindJSON(&attribute)

	code := model.CreateAttribute(&attribute)
	context.JSON(200, gin.H{
		"status":  code,
		"data":    attribute,
		"message": errmsg.GetErrorMsg(code),
	})
}

func GetAttribute(context *gin.Context) {
	id, err := strconv.Atoi(context.Query("id"))
	utils.HandleErr(err, "获取id错误")

	code, Attributes := model.GetAttributeInfo(id)
	context.JSON(200, gin.H{
		"status":  code,
		"data":    Attributes,
		"message": errmsg.GetErrorMsg(code),
	})
}

func DelAttribute(context *gin.Context) {
	id, err := strconv.Atoi(context.Query("id"))
	utils.HandleErr(err, "attribute删除失败")
	code := model.DeleteAttribute(id)
	context.JSON(200, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMsg(code),
	})
}

func EditAttribute(context *gin.Context) {
	id, err := strconv.Atoi(context.Query("id"))
	utils.HandleErr(err, "编辑用户信息失败")
	var attribute model.Attribute
	context.ShouldBindJSON(&attribute)
	code := model.UpdateAttribute(id, &attribute)
	context.JSON(200, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMsg(code),
	})
}
