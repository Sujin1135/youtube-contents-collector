package external

type DataAPI interface {
	search() (string, error)
}

type dataAPI struct {
	unit int
}

func search() (string, error) {
	return "", nil
}
