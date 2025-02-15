package v1

import (
	"channel-contents-collector/api/v1/content"
	"github.com/gin-gonic/gin"
)

func InitV1Group(router *gin.Engine) {
	v1 := router.Group("/v1")
	content.InitContentsGroup(v1)
}
