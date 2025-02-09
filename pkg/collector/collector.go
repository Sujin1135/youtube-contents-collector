package collector

type Collector interface {
	collect() (string, error)
}
