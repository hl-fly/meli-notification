package constants

type errorString string

func (err errorString) Error() string {
	return string(err)
}

const (
	ErrCacheMiss = errorString("cache miss: key not found")
)
