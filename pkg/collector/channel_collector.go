package collector

import (
	"channel-contents-collector/pkg/collector/domain"
	"channel-contents-collector/pkg/external"
	"github.com/pkg/errors"
	"time"
)

type ChannelCollector struct {
	dataAPI external.DataAPI
}

func NewChannelCollector(dataAPI external.DataAPI) *ChannelCollector {
	return &ChannelCollector{
		dataAPI: dataAPI,
	}
}

func (c *ChannelCollector) Collect(query string) ([]*domain.Channel, error) {
	data, err := c.dataAPI.Search(external.Channel, query)
	if data == nil {
		return nil, errors.Wrap(err, "channel response data can not be null")
	}
	if err != nil {
		return nil, errors.Wrap(err, "occurred an error when try to search channel data")
	}
	channels := c.convertToChannels(data.Items)
	return channels, err
}

func (c *ChannelCollector) convertToChannels(items []domain.Item) []*domain.Channel {
	channels := make([]*domain.Channel, len(items))
	now := time.Now()
	for i, v := range items {
		channels[i] = &domain.Channel{
			Id:          nil,
			ExternalId:  v.Snippet.ChannelId,
			Name:        v.Snippet.Title,
			Description: v.Snippet.Description,
			Published:   v.Snippet.PublishedAt,
			Created:     now,
			Modified:    now,
			Deleted:     nil,
		}
	}
	return channels
}
