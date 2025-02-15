package collector

import "channel-contents-collector/api/v1/content/domain"

type Collector interface {
	Collect(query string) (*domain.ContentResponse, error)
}
