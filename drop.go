package drop

type StorageService interface {
	PutObject(family, key string, value []byte) error
	DeleteObject(family, key string) error
	GetObjectValue(family, key string) ([]byte, error)
	GetObjectList(family string) ([]string, error)

	// GetObjectsValues
	// GetObjectListPaginated
}
