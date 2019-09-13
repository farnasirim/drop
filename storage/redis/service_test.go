package redis

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/farnasirim/drop"
)

func TestStorageEndToEnd(t *testing.T) {
	redis := NewStorageService("localhost:6379", 1)
	redis.r.FlushDB()
	var s drop.StorageService = redis

	ctx, cancelfunc := context.WithCancel(context.Background())

	{
		text := "one text"
		addr := "one addr"
		expId := 1

		recOrig := newRecord(-1, text, addr)

		id, err := s.PutRecord("public", recOrig)
		assert.Nil(t, err)
		assert.Equal(t, int64(expId), id)

		rec, err := s.GetRecord("public", int64(expId))
		assert.Nil(t, err)

		assert.Equal(t, recOrig.Text(), rec.Text())
		assert.Equal(t, recOrig.Address(), rec.Address())
	}

	{
		text := "two text"
		addr := "two addr"
		expId := 2

		recOrig := newRecord(-1, text, addr)

		id, err := s.PutRecord("public", recOrig)
		assert.Nil(t, err)
		assert.Equal(t, int64(expId), id)

		rec, err := s.GetRecord("public", int64(expId))
		assert.Nil(t, err)

		assert.Equal(t, recOrig.Text(), rec.Text())
		assert.Equal(t, recOrig.Address(), rec.Address())
	}

	{
		expId := 3

		rec, err := s.GetRecord("public", int64(expId))
		assert.NotNil(t, err)

		assert.Nil(t, rec)
	}

	{
		assert.Nil(t, s.DeleteRecord("public", 2))
		expId := 2

		rec, err := s.GetRecord("public", int64(expId))
		assert.NotNil(t, err)

		assert.Nil(t, rec)
	}

	{
		text := "three text"
		addr := "three addr"
		expId := 3

		recOrig := newRecord(-1, text, addr)

		id, err := s.PutRecord("public", recOrig)
		assert.Nil(t, err)
		assert.Equal(t, int64(expId), id)

		rec, err := s.GetRecord("public", int64(expId))
		assert.Nil(t, err)

		assert.Equal(t, recOrig.Text(), rec.Text())
		assert.Equal(t, recOrig.Address(), rec.Address())
	}

	all, las := s.AllRecords(ctx, "public")
	cre := s.AllCreateEventsAfter(ctx, "public", las)
	del := s.AllDeleteEvents(ctx, "public")
	time.Sleep(10 * time.Millisecond)

	done := make(chan struct{}, 10)

	assert.Equal(t, int64(3), las)
	go func() {
		first := <-all
		second := <-all

		if first.ID() > second.ID() {
			temp := first
			first = second
			second = temp
		}
		{
			text := "one text"
			addr := "one addr"
			one := newRecord(1, text, addr)
			assert.Equal(t, one.Text(), first.Text())
			assert.Equal(t, one.Address(), first.Address())
			assert.Equal(t, one.ID(), first.ID())
		}
		{
			text := "three text"
			addr := "three addr"
			th := newRecord(3, text, addr)
			assert.Equal(t, th.Text(), second.Text())
			assert.Equal(t, th.Address(), second.Address())
			assert.Equal(t, th.ID(), second.ID())
		}
		done <- struct{}{}
		<-all
		assert.Fail(t, "Read more than expected")
	}()

	go func() {
		{
			assert.Nil(t, s.DeleteRecord("public", 1))
			expId := 1

			rec, err := s.GetRecord("public", int64(expId))
			assert.NotNil(t, err)

			assert.Nil(t, rec)
		}
	}()
	go func() {
		text := "four text"
		addr := "four addr"
		expId := 4

		recOrig := newRecord(int64(expId), text, addr)

		id, err := s.PutRecord("public", recOrig)
		assert.Nil(t, err)
		assert.Equal(t, int64(expId), id)

		rec, err := s.GetRecord("public", int64(expId))
		assert.Nil(t, err)

		assert.Equal(t, recOrig.Text(), rec.Text())
		assert.Equal(t, recOrig.Address(), rec.Address())
	}()

	go func() {
		c := <-cre
		text := "four text"
		addr := "four addr"
		expId := 4
		recOrig := newRecord(int64(expId), text, addr)

		assert.Equal(t, recOrig.Text(), c.Text())
		assert.Equal(t, recOrig.Address(), c.Address())
		assert.Equal(t, recOrig.ID(), c.ID())

		done <- struct{}{}
		<-cre
		assert.Fail(t, "Got more create messages than expected")
	}()

	go func() {
		d := <-del

		assert.Equal(t, int64(1), d.ID())

		done <- struct{}{}
		<-del
		assert.Fail(t, "Got more del messages than expected")
	}()

	time.Sleep(10 * time.Millisecond)
	cancelfunc()

	close(done)
	num := 0
	for _ = range done {
		num += 1
	}

	time.Sleep(10 * time.Millisecond)
	assert.Equal(t, 3, num)
}
