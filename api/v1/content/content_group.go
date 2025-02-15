package content

import (
	"channel-contents-collector/api/v1/content/domain"
	"channel-contents-collector/pkg/collector"
	"channel-contents-collector/pkg/external"
	"github.com/gin-gonic/gin"
)

func InitContentsGroup(router *gin.RouterGroup) {
	dataApi := external.NewDataAPI()
	channelCollector := collector.NewChannelCollector(dataApi)
	content := router.Group("/contents")
	{
		content.POST("/collect", func(c *gin.Context) {
			var collectRequest domain.CollectRequest
			err := c.BindJSON(&collectRequest)
			if err != nil {
				return
			}

			collect, err := channelCollector.Collect(collectRequest.Keyword)
			if err != nil {
				return
			}

			c.JSON(200, collect)
		})
	}
}
