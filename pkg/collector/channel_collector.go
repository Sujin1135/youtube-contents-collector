package collector

import "channel-contents-collector/pkg/external"

type channelCollector struct {
	dataAPI *external.DataAPI
}

func NewChannelCollector(dataAPI *external.DataAPI) Collector {
	return &channelCollector{
		dataAPI: dataAPI,
	}
}

func (c *channelCollector) collect() (string, error) {
	return "", nil
}
