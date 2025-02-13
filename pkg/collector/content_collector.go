package collector

import (
	"channel-contents-collector/pkg/external"
	"fmt"
)

type contentCollector struct {
	dataAPI external.DataAPI
}

func NewContentCollector(dataAPI external.DataAPI) Collector {
	return &contentCollector{
		dataAPI: dataAPI,
	}
}

func (c *contentCollector) Collect(query string) (string, error) {
	data, _ := c.dataAPI.Search(external.Content, query)
	fmt.Println("found contents are " + data)
	return data, nil
}
