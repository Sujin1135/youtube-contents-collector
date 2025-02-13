package collector

import "channel-contents-collector/pkg/external"

type contentCollector struct {
	dataAPI external.DataAPI
}

func NewContentCollector(dataAPI external.DataAPI) Collector {
	return &contentCollector{
		dataAPI: dataAPI,
	}
}

func (c *contentCollector) Collect() (string, error) {
	return "", nil
}
