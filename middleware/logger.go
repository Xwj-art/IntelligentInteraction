package middleware

import (
	"Ai/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	reta "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"time"
)

func Logger() gin.HandlerFunc {
	filePath := "log/aiBlog"
	linkName := "lastest.log"
	src, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	utils.HandleErr(err, "日志错误")

	logger := logrus.New()
	logger.Out = src

	logger.SetLevel(logrus.DebugLevel)
	logWriter, err := reta.New(
		filePath+"%Y%m%d.log",
		reta.WithMaxAge(15*24*time.Hour),
		reta.WithRotationTime(24*time.Hour),
		reta.WithLinkName(linkName),
	)
	utils.HandleErr(err, "分割日志错误")
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
		logrus.TraceLevel: logWriter,
	}

	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.AddHook(Hook)
	return func(context *gin.Context) {
		startTime := time.Now()
		context.Next()
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%dms", int(math.Ceil(float64(stopTime.Nanoseconds())/1000000.0)))
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}
		statusCode := context.Writer.Status()
		clientIp := context.ClientIP()
		userAgent := context.Request.UserAgent()

		dataSize := context.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}

		method := context.Request.Method
		path := context.Request.RequestURI

		entry := logger.WithFields(logrus.Fields{
			"hostname":  hostName,
			"status":    statusCode,
			"spendtime": spendTime,
			"ip":        clientIp,
			"method":    method,
			"path":      path,
			"datasize":  dataSize,
			"agent":     userAgent,
		})
		if len(context.Errors) > 0 {
			entry.Error(context.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}
