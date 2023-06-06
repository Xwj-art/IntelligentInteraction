package routes

import (
	"Ai/api"
	"Ai/middleware"
	"Ai/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	// gin设置模式
	gin.SetMode(utils.AppMode)
	// 创建路由
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	// 路由分组
	router := r.Group("api")
	{
		router.GET("user", api.GetUser)
		router.GET("att/:id", api.GetAttribute)

		router.POST("login", api.Login)
		router.POST("user/add", api.AddUser)
		router.POST("att/add", api.AddAttribute)
	}

	auth := r.Group("api")
	auth.Use(middleware.JwtToken())
	{
		// User模块接口
		auth.PUT("user/:id", api.EditUser)
		auth.DELETE("user/:id", api.DelUser)
		// Attribute模块接口
		auth.DELETE("att/:id", api.DelAttribute)
		auth.PUT("att/:id", api.EditAttribute)

		auth.POST("upload", api.Upload)
	}

	r.Run(utils.HttpPort)
}
