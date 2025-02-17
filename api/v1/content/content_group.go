package content

import (
	error2 "channel-contents-collector/api/v1/error"
	"channel-contents-collector/pkg/collector"
	"channel-contents-collector/pkg/collector/domain"
	"channel-contents-collector/pkg/external"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func InitContentsGroup(router *gin.RouterGroup) {
	dataApi := external.NewDataAPI()
	channelCollector := collector.NewChannelCollector(dataApi)
	content := router.Group("/contents")
	{
		content.POST("/collect", func(c *gin.Context) {
			var collectRequest domain.CollectRequest
			err := c.ShouldBindJSON(&collectRequest)
			if err != nil {
				log.Printf("invalid request body: %+v", err)
				c.Error(error2.NewHttpError("Invalid request body", err.Error(), http.StatusBadRequest))
				return
			}

			collect, err := channelCollector.Collect(collectRequest.Keyword)
			if err != nil {
				c.Error(error2.NewHttpError("Occurred an error when collecting youtube contents", err.Error(), http.StatusBadRequest))
				return
			}

			c.JSON(200, collect)
		})
	}
}
