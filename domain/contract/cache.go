package contract

type CacheManager interface {
	GetStruct(key string, data interface{}) error
	SetStruct(key string, data interface{}) error
	Invalidate(key string) error
}
