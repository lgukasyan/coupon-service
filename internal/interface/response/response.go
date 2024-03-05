package response

import "github.com/gin-gonic/gin"

func Error(ctx *gin.Context, statusCode int, err error) {
	ctx.AbortWithStatusJSON(statusCode, gin.H{
		"err_description": err.Error(),
		"url":             ctx.Request.URL.Path,
		"method":          ctx.Request.Method,
		"status_code":     statusCode,
	})
}

func JSON(ctx *gin.Context, statusCode int, data interface{}) {
	ctx.JSON(statusCode, gin.H{
		"data":        data,
		"url":         ctx.Request.URL.Path,
		"method":      ctx.Request.Method,
		"status_code": statusCode,
	})
}
