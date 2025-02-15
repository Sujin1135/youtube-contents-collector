package main

import (
	"channel-contents-collector/api"
	"channel-contents-collector/pkg/collector"
	"channel-contents-collector/pkg/external"
	"fmt"
)

func main() {
	router := api.NewRouter()
	err := router.Run(":8080")
	if err != nil {
		return
	}
	dataApi := external.NewDataAPI()
	channelCollector := collector.NewChannelCollector(dataApi)
	contentCollector := collector.NewContentCollector(dataApi)

	fmt.Println(channelCollector.Collect("뷰티"))
	fmt.Println(contentCollector.Collect("뷰티"))
}
