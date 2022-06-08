package gin

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
	"time"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				trace := make([]byte, 4096)
				runtime.Stack(trace, true)
				var msg string
				msg += fmt.Sprintf("%s\n", r)
				for i := 0; ; i++ {
					_, file, line, ok := runtime.Caller(i)
					if !ok {
						break
					}
					msg += fmt.Sprintf("%s:%d\n", file, line)
				}
				logger.Logger.Errorf("Unknown error: %s", msg)
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    -1,
					"message": "unknown error",
					"data":    nil,
				})
				c.Abort()
				return
			}
		}()
		c.Next()
	}
}

func LogRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := c.Request.Header.Get(string(HeaderKeyTraceID))
		if traceID == "" {
			traceID = uuid.New().String()
		}

		buf, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(buf))

		var requestPayload map[string]interface{}
		_ = json.Unmarshal(buf, &requestPayload)

		defer func(t time.Time) {
			if !strings.Contains(c.Request.URL.Path, "health") {
				latency := time.Since(t)
				logger.Logger.Infow("access log",
					"url", c.Request.URL.String(),
					"method", c.Request.Method,
					"latency", fmt.Sprintf("%d ms", latency.Milliseconds()),
					"request", requestPayload,
					"status", c.Writer.Status(),
					"trace-id", traceID,
				)
			}
		}(time.Now())

		ctx := context.WithValue(c.Request.Context(), ContextKeyTraceID, traceID)

		c.Request = c.Request.WithContext(ctx)
		c.Request.Header.Set(string(HeaderKeyTraceID), traceID)
		c.Writer.Header().Set(string(HeaderKeyTraceID), traceID)

		c.Next()
	}
}

func ClientIp() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: client ip 取法需要再更改
		ctx := context.WithValue(c.Request.Context(), ContextKeyClientIp, c.ClientIP())
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func RateLimiterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if limiter == nil {
			c.Next()
		}

		err := limiter.Wait(c.Request.Context())
		if err != nil {
			return
		}

		c.Next()
	}
}

func CORS() gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"*"}
	return cors.New(corsConfig)
}
