package collector

type Collector interface {
	Collect() (string, error)
}
