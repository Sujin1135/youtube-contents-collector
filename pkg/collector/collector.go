package collector

type Collector interface {
	Collect(query string) (string, error)
}
