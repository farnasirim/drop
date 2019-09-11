package redis

import (
	"fmt"
	"strings"

	"github.com/go-redis/redis"
)

type StorageService struct {
	r *redis.Client
}

func NewStorageService(address string) *StorageService {
	return &StorageService{
		r: redis.NewClient(&redis.Options{Addr: address, Password: "", DB: 0}),
	}
}

func (s *StorageService) familyPrefix(family string) string {
	return fmt.Sprintf("%s__", family)
}

func (s *StorageService) makeKey(family, key string) string {
	return fmt.Sprintf("%s%s", s.familyPrefix(family), key)
}

func (s *StorageService) PutObject(family, key string, value []byte) error {
	if err := s.r.Set(s.makeKey(family, key), value, 0).Err(); err != nil {
		return err
	}
	return nil
}

func (s *StorageService) DeleteObject(family, key string) error {
	return s.r.Del(s.makeKey(family, key)).Err()
}

func (s *StorageService) GetObjectValue(family, key string) ([]byte, error) {
	println("getting: ", s.makeKey(family, key))
	val, err := s.r.Get(s.makeKey(family, key)).Bytes()
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (s *StorageService) GetObjectList(family string) ([]string, error) {
	keysList, err := s.r.Keys(s.familyPrefix(family) + "*").Result()
	if err != nil {
		return nil, err
	}
	ret := make([]string, 0, len(keysList))

	pref := s.familyPrefix(family)
	stripLen := len(pref)

	for _, text := range keysList {
		if strings.HasPrefix(text, pref) {
			ret = append(ret, text[stripLen:])
		}
	}

	return ret, nil
}
