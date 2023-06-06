package api

import (
	"Ai/errmsg"
	"Ai/middleware"
	"Ai/model"
	"github.com/gin-gonic/gin"
)

func Login(context *gin.Context) {
	var data model.User
	context.ShouldBindJSON(&data)
	var token string
	code, id := model.CheckLogin(data.Username, data.Password)

	if code == errmsg.SUCCESS {
		token, code = middleware.SetToken(data.Username)
	}
	context.JSON(200, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMsg(code),
		"token":   token,
		"id":      id,
	})
}
