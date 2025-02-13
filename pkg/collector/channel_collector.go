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

func (c *channelCollector) Collect() (string, error) {
	data, _ := c.dataAPI.Search()
	fmt.Println("data is " + data)
	return data, nil
}
