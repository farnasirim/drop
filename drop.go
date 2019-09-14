package drop

import (
	"context"
)

type Record interface {
	ID() int64
	Text() string
	Address() string
}

type StorageService interface {
	PutRecord(family string, rec Record) (int64, error)
	DeleteRecord(family string, key int64) error
	GetRecord(family string, key int64) (Record, error)

	AllRecordsAfter(ctx context.Context, family string, lastId int64) (<-chan Record, int64)
	AllCreateEventsAfter(ctx context.Context, family string, lastId int64) <-chan Record
	AllDeleteEvents(ctx context.Context, family string) <-chan Record

	// GetObjectsValues
	// GetObjectListPaginated
}
