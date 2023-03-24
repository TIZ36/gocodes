package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"time"
)

func DefaultStructedLogger() gin.HandlerFunc {
	return StructLogger(&log.Logger)
}

func StructLogger(logger *zerolog.Logger) gin.HandlerFunc {
	return func(context *gin.Context) {
		start := time.Now()
		path := context.Request.URL.Path
		raw := context.Request.URL.RawQuery

		context.Next()

		param := gin.LogFormatterParams{}

		param.TimeStamp = time.Now()
		param.Latency = param.TimeStamp.Sub(start)
		if param.Latency > time.Minute {
			param.Latency = param.Latency.Truncate(time.Second)
		}

		param.ClientIP = context.ClientIP()
		param.Method = context.Request.Method
		param.StatusCode = context.Writer.Status()
		param.ErrorMessage = context.Errors.ByType(gin.ErrorTypePrivate).String()
		param.BodySize = context.Writer.Size()
		if raw != "" {
			path = path + "?" + raw
		}
		param.Path = path

		var logEvent *zerolog.Event

		if context.Writer.Status() >= 500 {
			logEvent = logger.Error()
		} else {
			logEvent = logger.Info()
		}

		logEvent.Str("client_id", param.ClientIP).
			Str("method", param.Method).
			Int("status_code", param.StatusCode).
			Int("body_size", param.BodySize).
			Str("path", param.Path).
			Str("latency", param.Latency.String()).
			Msg(param.ErrorMessage)
	}
}
