package api

import (
	v1 "channel-contents-collector/api/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	v1.InitV1Group(router)
	return router
}
