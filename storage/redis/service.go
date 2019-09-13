package redis

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/go-redis/redis"

	"github.com/farnasirim/drop"
)

var (
	maxChanSize = 10000
)

type StorageService struct {
	r *redis.Client
}

func NewStorageService(address string, db int) *StorageService {
	return &StorageService{
		r: redis.NewClient(&redis.Options{Addr: address, Password: "", DB: db}),
	}
}

func (s *StorageService) familyPrefix(family string) string {
	return fmt.Sprintf("%s__", family)
}

func (s *StorageService) recordsPrefix(family string) string {
	return fmt.Sprintf("%s%s__", s.familyPrefix(family), "records")
}

func (s *StorageService) getRecordKey(family, key string) string {
	return fmt.Sprintf("%s%s", s.recordsPrefix(family), key)
}

func (s *StorageService) preprocessKey(key int64) string {
	return strings.Replace(base64.StdEncoding.EncodeToString([]byte(fmt.Sprint(key))), "\n", "", -1)
}

func (s *StorageService) getCreateChannelKey(family string) string {
	return s.familyPrefix(family) + "create_channel"
}

func (s *StorageService) getDeleteChannelKey(family string) string {
	return s.familyPrefix(family) + "delete_channel"
}

func (s *StorageService) PutRecord(family string, rec drop.Record) (int64, error) {
	las := s.incrLastId(family)
	key := s.preprocessKey(las)

	bytes, err := json.Marshal(newRecord(las, rec.Text(), rec.Address()))
	if err != nil {
		log.Println("in put record: ", err.Error())
	}
	err = s.r.Set(s.getRecordKey(family, key), bytes, 0).Err()
	if err != nil {
		return 0, err
	}
	return las, s.r.Publish(s.getCreateChannelKey(family), bytes).Err()
}

func (s *StorageService) DeleteRecord(family string, ikey int64) error {
	key := s.preprocessKey(ikey)
	err := s.r.Del(s.getRecordKey(family, key)).Err()
	if err != nil {
		return err
	}
	rec := newRecord(ikey, "", "")
	bytes, err := json.Marshal(rec)
	if err != nil {
		log.Println("in DeleteRecord", err.Error())
	}

	return s.r.Publish(s.getDeleteChannelKey(family), bytes).Err()
}

func (s *StorageService) GetRecord(family string, ikey int64) (drop.Record, error) {
	key := s.preprocessKey(ikey)
	val, err := s.r.Get(s.getRecordKey(family, key)).Bytes()
	if err != nil {
		return nil, err
	}
	rec := &Record{}
	if err := json.Unmarshal(val, rec); err != nil {
		return nil, err
	}
	return rec, nil
}

// func (s *StorageService) noh {

// func (s *StorageService) GetObjectList(family string) ([]string, error) {
// 	keysList, err := s.r.Keys(s.familyPrefix(family) + "*").Result()
// 	if err != nil {
// 		return nil, err
// 	}
// 	ret := make([]string, 0, len(keysList))
//
// 	pref := s.familyPrefix(family)
// 	stripLen := len(pref)
//
// 	for _, text := range keysList {
// 		if strings.HasPrefix(text, pref) {
// 			ret = append(ret, text[stripLen:])
// 		}
// 	}
//
// 	return ret, nil
// }

func (s *StorageService) lastIdKey(family string) string {
	return s.familyPrefix(family) + "last_id"
}

func (s *StorageService) incrLastId(family string) int64 {
	val := s.r.Incr(s.lastIdKey(family)).Val()
	return val
}

func (s *StorageService) lastId(family string) int64 {
	val, err := s.r.Get(s.lastIdKey(family)).Int64()
	if err != nil {
		log.Println("error getting lastId: ", err.Error())
		val = 0
	}
	return val
}

func (s *StorageService) allRecords(ctx context.Context, family string) []*Record {
	recs := make([]*Record, 0)

	keysList, err := s.r.Keys(s.recordsPrefix(family) + "*").Result()
	if err != nil {
		return recs
	}

	for _, key := range keysList {
		rec := &Record{}
		value := s.r.Get(key).Val()
		err := json.Unmarshal([]byte(value), rec)
		if err != nil {
			log.Println("Error unmarshalling value", value, err.Error())
		}
		recs = append(recs, rec)
	}

	return recs
}

func (s *StorageService) AllRecords(ctx context.Context, family string) (<-chan drop.Record, int64) {
	ret := make(chan drop.Record, maxChanSize)
	las := s.lastId(family)

	go func() {
		for _, rec := range s.allRecords(ctx, family) {
			select {
			case <-ctx.Done():
				return
			default:
				if rec.ID() <= las {
					ret <- rec
				}
			}
		}
	}()

	return ret, las
}

func (s *StorageService) AllCreateEventsAfter(ctx context.Context, family string, lastId int64) <-chan drop.Record {
	return s.subscribeToEvents(ctx, lastId, s.getCreateChannelKey(family))
}

func (s *StorageService) AllDeleteEvents(ctx context.Context, family string) <-chan drop.Record {
	return s.subscribeToEvents(ctx, 0, s.getDeleteChannelKey(family))
}

func (s *StorageService) subscribeToEvents(ctx context.Context, lastId int64, channelName string) <-chan drop.Record {
	ret := make(chan drop.Record, maxChanSize)

	go func() {
		pubsub := s.r.Subscribe(channelName)
		ch := pubsub.Channel()
		done := false
		for !done {
			select {
			case <-ctx.Done():
				done = true
				break
			case msg := <-ch:
				rec := &Record{}
				err := json.Unmarshal([]byte(msg.Payload), rec)
				if err != nil {
					log.Println(err.Error())
				}
				if rec.ID() > lastId {
					ret <- rec
				}
			}
		}
		pubsub.Close()
	}()

	return ret
}
