package collector

import (
	"channel-contents-collector/pkg/external"
	"fmt"
)

type channelCollector struct {
	dataAPI external.DataAPI
}

func NewChannelCollector(dataAPI external.DataAPI) Collector {
	return &channelCollector{
		dataAPI: dataAPI,
	}
}

func (c *channelCollector) Collect(query string) (string, error) {
	data, _ := c.dataAPI.Search(external.Channel, query)
	fmt.Println("found channels are " + data)
	return data, nil
}
