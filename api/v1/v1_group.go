package v1

import (
	"channel-contents-collector/api/v1/content"
	"channel-contents-collector/api/v1/middleware"
	"github.com/gin-gonic/gin"
)

func InitV1Group(router *gin.Engine) {
	v1 := router.Group("/v1")
	v1.Use(middleware.ErrorHandler())
	content.InitContentsGroup(v1)
}
