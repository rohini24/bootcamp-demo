package core

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type Controller struct{}

func (controller Controller) Json(ctx *gin.Context, result *gorm.DB, data interface{}) {
	if result.Error == nil {
		controller.successResponse(ctx, http.StatusOK, gin.H{"data": data})
		return
	}
	controller.errorResponse(ctx, http.StatusInternalServerError, result.Error.Error())
}

func (controller Controller) successResponse(ctx *gin.Context, status int, h gin.H) {
	h["error"] = false
	ctx.JSON(status, h)
}

func (controller Controller) errorResponse(ctx *gin.Context, status int, message string) {
	ctx.JSON(status, gin.H{
		"error":         true,
		"error-message": message,
		"data":          nil,
	})
}
