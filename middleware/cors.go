package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Cors() gin.HandlerFunc {
	return cors.New(
		cors.Config{
			//AllowAllOrigins:  true,
			AllowOrigins:  []string{"*"},
			AllowMethods:  []string{"*"},
			AllowHeaders:  []string{"*", "Authorization"},
			ExposeHeaders: []string{"Content-Length", "Authorization"},
			MaxAge:        12 * time.Hour,
		},
	)
}
