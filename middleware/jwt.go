package middleware

import (
	"Ai/errmsg"
	"Ai/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

var JwtKey = []byte(utils.JwtKey)
var code int

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// 生成token
func SetToken(username string) (string, int) {
	expireTime := time.Now().Add(10 * time.Hour)
	SetClaims := MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "Ai",
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	token, err := reqClaim.SignedString(JwtKey)
	if err != nil {
		return "", errmsg.ERR
	}
	return token, errmsg.SUCCESS
}

// 验证token
func CheckToken(token string) (*MyClaims, int) {
	setToken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if key, _ := setToken.Claims.(*MyClaims); setToken.Valid {
		return key, errmsg.SUCCESS
	}
	return nil, errmsg.ERR
}

// jwt中间件，控制token
func JwtToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenHeader := context.Request.Header.Get("Authorization")

		if tokenHeader == "" {
			code = errmsg.ERROR_TOKEN_EXIST
			context.JSON(200, gin.H{
				"code":    code,
				"message": errmsg.GetErrorMsg(code),
			})
			context.Abort()
			return
		}

		checkToken := strings.SplitN(tokenHeader, " ", 2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_TYPE
			context.JSON(200, gin.H{
				"code":    code,
				"message": errmsg.GetErrorMsg(code),
			})
			context.Abort()
			return
		}

		key, tCode := CheckToken(checkToken[1])
		if tCode == errmsg.ERR {
			code = errmsg.ERROR_TOKEN_WRONG
			context.JSON(200, gin.H{
				"code":    code,
				"message": errmsg.GetErrorMsg(code),
			})
			context.Abort()
			return
		}

		if time.Now().Unix() > key.ExpiresAt {
			code = errmsg.ERROR_TOKEN_RUNTIME
			context.JSON(200, gin.H{
				"code":    code,
				"message": errmsg.GetErrorMsg(code),
			})
			context.Abort()
			return
		}

		context.Set("username", key.Username)
		context.Next()
	}
}
