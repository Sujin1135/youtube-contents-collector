package collector

import (
	"channel-contents-collector/pkg/collector/domain"
	"channel-contents-collector/pkg/external"
)

type contentCollector struct {
	dataAPI external.DataAPI
}

func NewContentCollector(dataAPI external.DataAPI) Collector {
	return &contentCollector{
		dataAPI: dataAPI,
	}
}

func (c *contentCollector) Collect(query string) (*domain.ContentResponse, error) {
	data, _ := c.dataAPI.Search(external.Content, query)
	return data, nil
}
