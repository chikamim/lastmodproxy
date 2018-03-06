package main

import (
	"time"

	"github.com/rapidloop/skv"
)

var bucket = []byte("lastmodified")

type BoltStore struct {
	Filepath string
}

type TimeStorer interface {
	Set(key string, value time.Time) error
	Get(key string) (value time.Time, err error)
	Delete(key string) error
}

func NewBoldTimeStore(filepath string) *BoltStore {
	s := BoltStore{filepath}

	return &s
}

func (s *BoltStore) Get(key string) (value time.Time, err error) {
	store, err := skv.Open(s.Filepath)
	if err != nil {
		return time.Time{}, err
	}
	defer store.Close()

	err = store.Get(key, &value)
	return value, err
}

func (s *BoltStore) Set(key string, value time.Time) error {
	store, err := skv.Open(s.Filepath)
	if err != nil {
		return err
	}
	defer store.Close()

	return store.Put(key, value)
}

func (s *BoltStore) Delete(key string) error {
	store, err := skv.Open(s.Filepath)
	if err != nil {
		return err
	}
	defer store.Close()

	return store.Delete(key)
}
