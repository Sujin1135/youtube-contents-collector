package collector

import (
	"channel-contents-collector/pkg/collector/domain"
	"channel-contents-collector/pkg/external"
)

type ContentCollector struct {
	dataAPI external.DataAPI
}

func NewContentCollector(dataAPI external.DataAPI) *ContentCollector {
	return &ContentCollector{
		dataAPI: dataAPI,
	}
}

func (c *ContentCollector) Collect(query string) (*domain.ContentResponse, error) {
	data, err := c.dataAPI.Search(external.Content, query)
	return data, err
}
