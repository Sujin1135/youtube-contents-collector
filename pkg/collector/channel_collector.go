package collector

import (
	"channel-contents-collector/pkg/collector/domain"
	"channel-contents-collector/pkg/external"
)

type channelCollector struct {
	dataAPI external.DataAPI
}

func NewChannelCollector(dataAPI external.DataAPI) Collector {
	return &channelCollector{
		dataAPI: dataAPI,
	}
}

func (c *channelCollector) Collect(query string) (*domain.ContentResponse, error) {
	data, _ := c.dataAPI.Search(external.Channel, query)
	return data, nil
}
