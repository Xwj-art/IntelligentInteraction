package api

import (
	"Ai/errmsg"
	"Ai/model"
	"Ai/utils"
	"Ai/validate"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 控制数据库的读写

var code int

// 新增用户
func AddUser(context *gin.Context) {
	var user model.User
	var msg string
	context.ShouldBindJSON(&user)
	msg, code = validate.Validate(&user)
	if code != errmsg.SUCCESS {
		context.JSON(200, gin.H{
			"status":  code,
			"message": msg,
		})
		return
	}

	code = model.CheckUser(user.Username)
	if code == errmsg.SUCCESS {
		model.CreateUser(&user)
	} else {
		code = errmsg.ERROR_NAME_USED
	}

	context.JSON(200, gin.H{
		"status":  code,
		"data":    user,
		"message": errmsg.GetErrorMsg(code),
	})
}

// 获取用户信息
func GetUser(context *gin.Context) {
	id, err := strconv.Atoi(context.Query("id"))
	utils.HandleErr(err, "id接收失败")

	code, user := model.GetUserInfo(id)
	context.JSON(200, gin.H{
		"status":  code,
		"data":    user,
		"message": errmsg.GetErrorMsg(code),
	})
}

// 删除用户
func DelUser(context *gin.Context) {
	id, err := strconv.Atoi(context.Query("id"))
	utils.HandleErr(err, "删除失败")
	code = model.DeleteUser(id)
	context.JSON(200, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMsg(code),
	})
}

// 更新用户信息
func EditUser(context *gin.Context) {
	id, err := strconv.Atoi(context.Query("id"))
	fmt.Println(id)
	utils.HandleErr(err, "编辑用户信息失败")
	var user model.User
	context.ShouldBindJSON(&user)
	code = model.CheckUser(user.Username)
	if code == errmsg.SUCCESS {
		model.UpdateUser(id, &user)
	} else if code == errmsg.ERROR_NAME_USED {
		context.Abort()
	}
	context.JSON(200, gin.H{
		"status":  code,
		"data":    user,
		"message": errmsg.GetErrorMsg(code),
	})
}
