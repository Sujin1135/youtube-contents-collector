package collector

import (
	"channel-contents-collector/pkg/external"
	"fmt"
	"github.com/pkg/errors"
)

type ContentsCollector struct {
	channelCollector *ChannelCollector
}

func (c *ContentsCollector) Collect(query string) error {
	channels, err := c.channelCollector.Collect(query)
	if err != nil {
		return errors.Wrap(err, "error collecting channels")
	}

	fmt.Println("Channels as follow:")
	for _, v := range channels {
		fmt.Printf("%v %v %v\n", v.ExternalId, v.Name, v.Description)
	}
	return nil
}

func NewContentsCollector(api external.DataAPI) *ContentsCollector {
	return &ContentsCollector{
		channelCollector: NewChannelCollector(api),
	}
}
