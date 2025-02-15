package collector

import (
	"channel-contents-collector/pkg/collector/domain"
)

type Collector interface {
	Collect(query string) (*domain.ContentResponse, error)
}
