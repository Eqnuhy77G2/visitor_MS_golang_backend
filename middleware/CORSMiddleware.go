package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin","http://127.0.0.1:8080")
		ctx.Writer.Header().Set("Access-Control-Max-Age","86400")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods","*")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers","*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials","*")

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(200)
		}else {
			ctx.Next()
		}
	}
}